package daemon

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/autoai-org/aid/internal/daemon/handlers"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

//go:embed build
var console embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func getRouter() *gin.Engine {
	if os.Getenv("AID_PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	tp := initTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	r := gin.Default()
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware("aid-server"))
	r.Use(static.Serve("/", EmbedFolder(console, "build")))
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.POST("/preflight", handlers.PreflightHandler)
		api.POST("/query", graphqlHandler())
		api.GET("/playground", playgroundHandler())
		api.GET("/solver/:solverID", handlers.SolverInformationHandler)
		api.GET("/package/:vendor/:package", handlers.GetPackageConfiguration)
		api.PUT("/containers/", handlers.CreateContainerHandler)
		api.POST("/running/:runningId/:path", handlers.ForwardHandler)
		api.POST("/mutations", handlers.MutateSolverStatusHandler)
		api.POST("/install", handlers.InstallPackageHandler)
		api.GET("/logs/:logid", func(c *gin.Context) {
			serveWS(c)
		})
	}
	r.NoRoute(func(c *gin.Context) {
		filecontent, err := console.ReadFile(filepath.Join("build", "index.html"))
		if err != nil {
			utilities.Formatter.Error(err.Error())
			c.JSON(http.StatusNotFound, handlers.ErrorResponse{})
		}
		c.Data(200, "text/html;charset=utf-8", filecontent)
	})
	return r
}
