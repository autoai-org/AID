// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package handlers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	entContainer "github.com/autoai-org/aid/ent/generated/container"
	"github.com/autoai-org/aid/internal/database"
	"github.com/gin-gonic/gin"
)

func ForwardHandler(c *gin.Context) {
	runningContainerID := c.Param("runningId")
	requestPath := c.Param("path")
	runningContainer, err := database.NewDefaultDB().Container.Query().Where(
		entContainer.UID(runningContainerID)).First(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Cannot fetch container " + runningContainerID + " from database."})
	}
	if !runningContainer.Running {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "The container " + runningContainerID + " is not running."})

	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
	target := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1:" + runningContainer.Port,
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
