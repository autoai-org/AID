/*  This file handles daemon and services related tasks.
By using cvpm daemon install, it will install a system service under current user.
You can uninstall that service by using cvpm daemon uninstall */
package main

import (
	"github.com/fatih/color"
	"github.com/gin-contrib/static"
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
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Solver string `json:"solver"`
	Port   string `json:"port"`
}

// POST /repos/running -> run a solver in this repo
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

// GET /repos/running -> return running repositories
func GetRunningReposHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RunningRepos)
}

// GET /solvers/running -> return running solvers
func GetRunningSolversHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RunningSolvers)
}

// GET /solvers/running/:vendor/:name -> return running solvers in this package
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

// POST /repos -> install a repo
type AddRepoRequest struct {
	RepoType string `json:"type"`
	URL      string `json:"url"`
}

func PostReposHandler(c *gin.Context) {
	config := readConfig()
	var addRepoRequest AddRepoRequest
	c.BindJSON(&addRepoRequest)
	log.Println(addRepoRequest.RepoType)
	if addRepoRequest.RepoType == "git" {
		InstallFromGit(addRepoRequest.URL)
		c.JSON(http.StatusOK, config.Repositories)
	} else {
		c.JSON(http.StatusBadRequest, config.Repositories)
	}
}

// GET /repos -> return all repositories
func GetReposHandler(c *gin.Context) {
	config := readConfig()
	c.JSON(http.StatusOK, config.Repositories)
}

// GET /repo/meta/:vendor/:name -> return repository meta info
func GetRepoMetaHandler(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	c.JSON(http.StatusOK, GetMetaInfo(vendor, name))
}

// GET /system -> return system info
func GetSystemHandler(c *gin.Context) {
	c.JSON(http.StatusOK, getSystemInfo())
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

// set global header to enable cors
// and set response header
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
	config := readConfig()
	webuiFolder := filepath.Join(config.Local.LocalFolder, "webui")
	color.Red("Initiating")
	var err error
	socketServer, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(BeforeResponse())
	watchLogs(socketServer)
	r.Use(static.Serve("/", static.LocalFile(webuiFolder, false)))
	// Status Related Handlers
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"daemon": "running",
		})
	})
	// System Related Handlers
	r.GET("/system", GetSystemHandler)
	// Repo Related Routes
	r.GET("/repo/meta/:vendor/:name", GetRepoMetaHandler)
	r.POST("/repo/running", PostRunningRepoHandler)
	r.GET("/repos", GetReposHandler)
	r.GET("/repos/running", GetRunningReposHandler)
	r.POST("/repos", PostReposHandler)
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
