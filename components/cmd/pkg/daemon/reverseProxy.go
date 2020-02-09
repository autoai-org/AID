// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

func serveReverseProxy(c *gin.Context) {
	// runningSolverID := c.Param("runningId")
	requestPath := c.Param("path")
	//runningSolver := entities.GetRunningSolver(runningSolverID)
	// if the solver is not found
	/*
		if runningSolver.EntryPoint == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "404",
				"msg":   "Target Not Found",
				"help":  "https://aid.autoai.org",
			})
		}
	*/
	// the solver is running
	body, err := ioutil.ReadAll(c.Request.Body)
	utilities.CheckError(err, "Cannot read the request")
	target := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1:8081",
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
