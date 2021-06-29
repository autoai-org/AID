package daemon

import (
	"context"
	"log"
	"os"

	"github.com/autoai-org/aid/internal/daemon/handlers"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

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
	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware("aid-server"))
	r.Use(beforeResponse())
	r.GET("/ping", handlers.PingHandler)
	r.POST("/preflight", handlers.PreflightHandler)
	r.POST("/query", graphqlHandler())
	r.GET("/playground", playgroundHandler())
	r.GET("/", handlers.HelloHandler)
	return r
}
