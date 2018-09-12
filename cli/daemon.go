package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fatih/color"
	"net/http"
)

const DaemonPort = "10590"
var RunningRepos []Repository
type RunRepoRequest struct {
	Name string `json:name`
	Vendor string `json:vendor`
	Solver string `json:solver`
}

func PostRepoHandler (c *gin.Context) {
	var runRepoRequest RunRepoRequest
	c.BindJSON(&runRepoRequest)
	go runRepo(runRepoRequest.Vendor, runRepoRequest.Name, runRepoRequest.Solver)
	c.JSON(http.StatusOK, gin.H{
		"code": "success",
	})
}

func GetReposHandler (c *gin.Context) {

}

func runServer (port string) {
	color.Red("Initiating")
  r := gin.Default()
  r.GET("/status", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"daemon":"running",
		})
  })
	r.POST("/repo", PostRepoHandler)
  r.Run("127.0.0.1:" + port)
}
