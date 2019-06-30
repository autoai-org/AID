package daemon

import (
	"log"
	"bytes"
	"net/url"
	"net/http"
	"io/ioutil"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
	"github.com/getsentry/raven-go"
	"github.com/unarxiv/cvpm/pkg/query"
	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/entity"
	"github.com/unarxiv/cvpm/pkg/runtime"
	"github.com/unarxiv/cvpm/pkg/utility"
	"github.com/unarxiv/cvpm/pkg/repository"
)

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

// GetReposHandler handles GET /repos -> return all repositories
func GetReposHandler(c *gin.Context) {
	config := config.Read()
	c.JSON(http.StatusOK, config.Repositories)
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

// ReverseProxy for Calling Solvers and return real results
func ReverseProxy(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	solver := c.Param("solver")
	requestPath := c.Param("path")
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
		Path:   "/" + requestPath,
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
			"help":  "https://cvflow.autoai.org",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "200",
			"info":  "Successfully stopped",
			"help":  "https://cvflow.autoai.org",
		})
	}
}