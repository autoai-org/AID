package daemon

import (
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	return r
}
