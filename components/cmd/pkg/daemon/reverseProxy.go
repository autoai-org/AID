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
	"sync"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

const (
	// MaxAttempts specifies the max allowed requests
	MaxAttempts int = 3
	// Attempts is the retried times
	Attempts int = iota
	// Retry saves the retry time
	Retry
)

// Backend defines the struct for remote nodes
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// SetAlive tells the server that a certain node is alive or not
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns if the node is alive or not
func (b *Backend) IsAlive() (alive bool) {
	b.mux.Lock()
	alive = b.Alive
	b.mux.RUnlock()
	return alive
}

// ServerPool saves all backends and current backend
type ServerPool struct {
	backends []*Backend
	current  uint64
}

func serveReverseProxy(c *gin.Context) {
	runningSolverID := c.Param("runningId")
	requestPath := c.Param("path")
	runningSolver := entities.GetRunningSolver(runningSolverID)
	// if the solver is not found

	if runningSolver.EntryPoint == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404",
			"msg":   "Target Not Found",
			"help":  "https://aid.autoai.org/guide/misc/troubleshooting#404-target-not-found",
		})
	}
	// the solver is running
	body, err := ioutil.ReadAll(c.Request.Body)
	utilities.CheckError(err, "Cannot read the request")
	target := url.URL{
		Scheme: "http",
		Host:   runningSolver.EntryPoint,
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
