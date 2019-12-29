package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getPackages : GET /packages -> return packages
func getPackages(c *gin.Context) {
	packages := entities.FetchPackages()
	c.JSON(http.StatusOK, packages)
}

// version : Test if it works
func getVersion(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}