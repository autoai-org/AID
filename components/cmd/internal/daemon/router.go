package daemon

import (
	"os"

	"github.com/autoai-org/aid/internal/runtime/git"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	if os.Getenv("AID_PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	gitService := git.GetService()
	r := gin.Default()
	p := NewPrometheus("gin")
	p.Use(r)
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	r.Any("/git", gin.WrapH(gitService))
	return r
}
