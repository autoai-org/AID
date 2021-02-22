// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

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

// AddBackend adds a new node into the serverpool
func (s *ServerPool) AddBackend(backend *Backend) {
	s.backends = append(s.backends, backend)
}

// NextIndex atomically increase the count and return an index
func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

// SetBackendStatus changes the status of a certain backend
func (s *ServerPool) SetBackendStatus(backendURL *url.URL, alive bool) {
	for _, b := range s.backends {
		if b.URL.String() == backendURL.String() {
			b.SetAlive(alive)
			break
		}
	}
}

// GetNextNode returns next active node to connect
func (s *ServerPool) GetNextNode() *Backend {
	next := s.NextIndex()
	l := len(s.backends) + next
	for i := next; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil
}

// IsBackendAlive checks if the given url is alive
func IsBackendAlive(u *url.URL) bool {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		utilities.Formatter.Error("cannot access " + u.String())
		return false
	}
	conn.Close()
	return true
}

// HealthCheck pings the remote node and update the status of the node
func (s *ServerPool) HealthCheck() {
	for _, b := range s.backends {
		status := "up"
		alive := IsBackendAlive(b.URL)
		b.SetAlive(alive)
		if !alive {
			status = "down"
		}
		log.Printf("%s [%s]\n", b.URL, status)
	}
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
