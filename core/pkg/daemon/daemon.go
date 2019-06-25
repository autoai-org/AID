// Package daemon handles the daemon server
// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*  This file handles daemon and services related tasks.
By using cvpm daemon install, it will install a system service under current user.
You can uninstall that service by using cvpm daemon uninstall */

package daemon

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/entity"
	"github.com/unarxiv/cvpm/pkg/query"
	"github.com/unarxiv/cvpm/pkg/repository"
	"github.com/unarxiv/cvpm/pkg/runtime"
	"github.com/unarxiv/cvpm/pkg/utility"

	"github.com/fatih/color"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

// DaemonPort Default Running Port
const DaemonPort = "10590"

// socketServer defines of Running Repos
var socketServer *socketio.Server

// RunRepoRequest is the struct of a Request to Run Repo
type RunRepoRequest struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Solver string `json:"solver"`
	Port   string `json:"port"`
}

// PostRunningRepoHandler handles POST /repos/running -> run a solver in this repo
func PostRunningRepoHandler(c *gin.Context) {
	var runRepoRequest RunRepoRequest
	c.BindJSON(&runRepoRequest)
	if runRepoRequest.Port == "" {
		runRepoRequest.Port = utility.FindNextOpenPort(8080)
	}
	go repository.Run(runRepoRequest.Vendor, runRepoRequest.Name, runRepoRequest.Solver, runRepoRequest.Port)
	c.JSON(http.StatusOK, gin.H{
		"code": "success",
		"port": runRepoRequest.Port,
	})
}

// GetRunningReposHandler handles GET /repos/running -> return running repositories
func GetRunningReposHandler(c *gin.Context) {
	c.JSON(http.StatusOK, runtime.RunningRepos)
}

// GetRunningSolversHandler handles GET /solvers/running -> return running solvers
func GetRunningSolversHandler(c *gin.Context) {
	c.JSON(http.StatusOK, runtime.RunningSolvers)
}

// GetRunningSolversByPackageHandler handles GET /solvers/running/:vendor/:name -> return running solvers in this package
func GetRunningSolversByPackageHandler(c *gin.Context) {
	vendor := c.Param("vendor")
	packageName := c.Param("package")
	var runningSolversInPackage []entity.RepoSolver
	for _, runningSolver := range runtime.RunningSolvers {
		if runningSolver.Vendor == vendor && runningSolver.Package == packageName {
			runningSolversInPackage = append(runningSolversInPackage, runningSolver)
		}
	}
	c.JSON(http.StatusOK, runningSolversInPackage)
}

// AddRepoRequest handles POST /repos -> install a repo
type AddRepoRequest struct {
	RepoType string `json:"type"`
	URL      string `json:"url"`
}

// PostReposHandler handles POST requests to /repos
func PostReposHandler(c *gin.Context) {
	config := config.Read()
	var addRepoRequest AddRepoRequest
	c.BindJSON(&addRepoRequest)
	if addRepoRequest.RepoType == "git" {
		repository.InstallFromGit(addRepoRequest.URL)
		c.JSON(http.StatusOK, config.Repositories)
	} else {
		c.JSON(http.StatusBadRequest, config.Repositories)
	}
}

// GetReposHandler handles GET /repos -> return all repositories
func GetReposHandler(c *gin.Context) {
	config := config.Read()
	c.JSON(http.StatusOK, config.Repositories)
}

// GetRepoMetaHandler : GET /repo/meta/:vendor/:name -> return repository meta info
func GetRepoMetaHandler(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	c.JSON(http.StatusOK, repository.GetMetaInfo(vendor, name))
}

// GetSystemHandler : GET /system -> return system info
func GetSystemHandler(c *gin.Context) {
	c.JSON(http.StatusOK, utility.GetSystemInfo())
}

// socketHandler handles Socket Request, this is to be fixed
func socketHandler(c *gin.Context) {
	socketServer.OnConnect("/", func(so socketio.Conn) error {
		// so.J
		return nil
	})
	/* Todo: Upgrade to v1.4 of go-socket.io
	socketServer.On("connection", func(so socketio.Socket) {
		so.Join("cvpm-webtail")
	})
	*/
	socketServer.ServeHTTP(c.Writer, c.Request)
}

// BeforeResponse set global header to enable cors and set response header
func BeforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("cvpm-version", "0.0.3@alpha")
		if c.Writer.Header().Get("Access-Control-Allow-Origin") == "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(http.StatusOK)
		}
	}
}

// StopInferProcess deletes Running Solver Process
func StopInferProcess(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	solver := c.Param("solver")
	var runningPort string
	for _, runningSolver := range runtime.RunningSolvers {
		if runningSolver.Vendor == vendor && runningSolver.Package == name && runningSolver.SolverName == solver {
			runningPort = runningSolver.Port
		}
	}
	// if the solver is not found
	if runningPort == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404",
			"info":  "Solver Not Running, if you want to force it running, please attach a force=true in your request body",
			"help":  "https://cvpm.autoai.org",
		})
	}
	stopReturnValue := query.StopInferEngine(runningPort)
	if !stopReturnValue {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "403",
			"info":  "Bad request from the infer engine",
			"help":  "https://cvpm.autoai.org",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "200",
			"info":  "Successfully stopped",
			"help":  "https://cvpm.autoai.org",
		})
	}
}

// ReverseProxy for Calling Solvers and return real results
func ReverseProxy(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	solver := c.Param("solver")
	var runningPort string
	for _, runningSolver := range runtime.RunningSolvers {
		if runningSolver.Vendor == vendor && runningSolver.Package == name && runningSolver.SolverName == solver {
			runningPort = runningSolver.Port
		}
	}
	// if the solver is not found
	if runningPort == "" {
		c.JSON(http.StatusNotImplemented, gin.H{
			"error": "501",
			"info":  "Solver Not Running, if you want to force it running, please attach a force=true in your request body",
			"help":  "https://cvflow.autoai.org",
		})
	}
	// the solver is running
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		raven.CaptureErrorAndWait(err, nil)
	}
	target := url.URL{
		Scheme: "http",
		Host:   "localhost:" + runningPort,
		Path:   "/infer",
	}

	director := func(req *http.Request) {
		req.Host = target.Host
		req.URL.Host = req.Host
		req.URL.Scheme = target.Scheme
		req.URL.Path = target.Path
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	proxy := httputil.NewSingleHostReverseProxy(&target)
	proxy.Director = director
	proxy.ServeHTTP(c.Writer, c.Request)
}

// RunServer starts the daemon server
func RunServer(port string) {
	color.Red("Initiating")
	var err error
	socketServer, err = socketio.NewServer(nil)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	r := getRouter()
	r.Run("0.0.0.0:" + port)
}
