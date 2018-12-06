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
var RunningSolvers []RepoSolver
var socketServer *socketio.Server

// Struct of a Request to Run Repo
type RunRepoRequest struct {
	Name   string `json:name`
	Vendor string `json:vendor`
	Solver string `json:solver`
	Port   string `json:port`
}

// Handle Post Request -> Run a Repo
func PostRunningRepoHandler(c *gin.Context) {
	var runRepoRequest RunRepoRequest
	c.BindJSON(&runRepoRequest)
	log.Println(runRepoRequest.Port)
	if runRepoRequest.Port == "" {
		runRepoRequest.Port = findNextOpenPort(8080)
	}
	go runRepo(runRepoRequest.Vendor, runRepoRequest.Name, runRepoRequest.Solver, runRepoRequest.Port)
	c.JSON(http.StatusOK, gin.H{
		"code": "success",
		"port": runRepoRequest.Port,
	})
}

// Handle Get Request -> Get Running Repo
func GetRunningReposHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RunningRepos)
}

//
func GetRunningSolversHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RunningSolvers)
}

func GetRunningSolversByPackageHandler(c *gin.Context) {
	vendor := c.Param("vendor")
	packageName := c.Param("package")
	var runningSolversInPackage []RepoSolver
	for _, runningSolver := range RunningSolvers {
		if runningSolver.Vendor == vendor && runningSolver.Package == packageName {
			runningSolversInPackage = append(runningSolversInPackage, runningSolver)
		}
	}
	c.JSON(http.StatusOK, runningSolversInPackage)
}

// Handle Get Request -> Get All Repos
func GetReposHandler(c *gin.Context) {
	config := readConfig()
	c.JSON(http.StatusOK, config.Repositories)
}

// Handle Get Repository Meta Info
func GetRepoMetaHandler(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	c.JSON(http.StatusOK, GetMetaInfo(vendor, name))
}

// Handle Socket Request
func socketHandler(c *gin.Context) {
	socketServer.On("connection", func(so socketio.Socket) {
		so.Join("cvpm-webtail")
	})
	socketServer.ServeHTTP(c.Writer, c.Request)
}

// write log to socket stream
func writeLog(filepath string, server *socketio.Server, eventName string) {
	log.Println("Writing Logs")
	t, err := tail.TailFile(filepath, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		server.BroadcastTo("cvpm-webtail", eventName, line.Text)
	}
}

// watched log source
func watchLogs(server *socketio.Server) {
	// System Log
	cvpmLogsLocation := getLogsLocation()
	go writeLog(filepath.Join(cvpmLogsLocation, "system.log"), server, "system")
	go writeLog(filepath.Join(cvpmLogsLocation, "package.log"), server, "package")
}

// global header
func BeforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("cvpm-version", "0.0.3@alpha")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
		}
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
	// Repo Related Routes
	r.GET("/repo/meta/:vendor/:name", GetRepoMetaHandler)
	r.POST("/repo/running", PostRunningRepoHandler)
	r.GET("/repos", GetReposHandler)
	r.GET("/repos/running", GetRunningReposHandler)

	// Solver Related Routers
	r.GET("/solvers/running", GetRunningSolversHandler)
	r.GET("/solvers/running/:vendor/:package", GetRunningSolversByPackageHandler)
	// Socket Related Routes
	r.GET("/socket.io/", socketHandler)
	r.POST("/socket.io/", socketHandler)
	r.Handle("WS", "/socket.io/", socketHandler)
	r.Handle("WSS", "/socket.io/", socketHandler)
	r.Run("0.0.0.0:" + port)
}
