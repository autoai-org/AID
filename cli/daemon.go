/*
	This file handles daemon and services related tasks.
	By using cvpm daemon install, it will install a system service under current user.
	You can uninstall that service by using cvpm daemon uninstall
*/
package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Default Running Port
const DaemonPort = "10590"

// Definition of Running Repos
var RunningRepos []Repository

// Struct of a Request to Run Repo
type RunRepoRequest struct {
	Name   string `json:name`
	Vendor string `json:vendor`
	Solver string `json:solver`
	Port   string `json:port`
}

// Handle Post Request -> Run a Repo
func PostRepoHandler(c *gin.Context) {
	var runRepoRequest RunRepoRequest
	c.BindJSON(&runRepoRequest)
	go runRepo(runRepoRequest.Vendor, runRepoRequest.Name, runRepoRequest.Solver, runRepoRequest.Port)
	c.JSON(http.StatusOK, gin.H{
		"code": "success",
	})
}

// Handle Get Request -> Get Running Repo
func GetReposHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RunningRepos)
}

/* Run the Server and Do Mount Endpoint
/status -> Get to fetch System Status
/repo -> Post to Run a Repo
/repos -> Get to fetch Running Repos
*/
func runServer(port string) {
	color.Red("Initiating")
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"daemon": "running",
		})
	})
	r.POST("/repo", PostRepoHandler)
	r.GET("/repos", GetReposHandler)
	r.Run("127.0.0.1:" + port)
}
