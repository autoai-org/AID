package daemon

import (
	"os"

	"github.com/autoai-org/aid/internal/daemon/handlers"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	if os.Getenv("AID_PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	r.GET("/ping", handlers.PingHandler)
	r.POST("/preflight", handlers.PreflightHandler)
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	return r
}
