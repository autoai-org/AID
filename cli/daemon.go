/*
	This file handles daemon and services related tasks.
	By using cvpm daemon install, it will install a system service under current user.
	You can uninstall that service by using cvpm daemon uninstall
*/
package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"github.com/hpcloud/tail"
	"log"
	"net/http"
	"path/filepath"
)

// Default Running Port
const DaemonPort = "10590"

// Definition of Running Repos
var RunningRepos []Repository
var socketServer *socketio.Server

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

// Handle Socket Request
func socketHandler(c *gin.Context) {
	socketServer.On("connection", func(so socketio.Socket) {
		so.Join("cvpm-webtail")
	})
	socketServer.ServeHTTP(c.Writer, c.Request)
}

// write log to socket stream
func writeLog(filepath string, server *socketio.Server) {
	log.Println("Writing Logs")
	t, err := tail.TailFile(filepath, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		log.Println("Broadcasting")
		server.BroadcastTo("cvpm-webtail", "logevent", line.Text)
	}
}

// watched log source
func watchLogs(server *socketio.Server) {
	// System Log
	cvpmLogsLocation := getLogsLocation()
	log.Println(cvpmLogsLocation)
	go writeLog(filepath.Join(cvpmLogsLocation, "system.log"), server)
}

// global header
func BeforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("cvpm-version", "0.0.3@alpha")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

/* Run the Server and Do Mount Endpoint
/status -> Get to fetch System Status
/repo -> Post to Run a Repo
/repos -> Get to fetch Running Repos
*/
func runServer(port string) {
	color.Red("Initiating")
	var err error
	socketServer, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(BeforeResponse())
	watchLogs(socketServer)
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"daemon": "running",
		})
	})
	r.POST("/repo", PostRepoHandler)
	r.GET("/repos", GetReposHandler)
	r.GET("/socket.io/", socketHandler)
	r.POST("/socket.io/", socketHandler)
	r.Handle("WS", "/socket.io/", socketHandler)
	r.Handle("WSS", "/socket.io/", socketHandler)
	r.Run("0.0.0.0:" + port)
}
