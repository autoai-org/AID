// Package daemon handles the daemon server
// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*  This file handles daemon and services related tasks.
By using cvpm daemon install, it will install a system service under current user.
You can uninstall that service by using cvpm daemon uninstall */

package daemon

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DaemonPort Default Running Port
const DaemonPort = "10590"

// RunRepoRequest is the struct of a Request to Run Repo
type RunRepoRequest struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Solver string `json:"solver"`
	Port   string `json:"port"`
}

// AddRepoRequest handles POST /repos -> install a repo
type AddRepoRequest struct {
	RepoType string `json:"type"`
	URL      string `json:"url"`
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

// RunServer starts the daemon server
func RunServer(port string) {
	color.Red("Initiating")
	r := getRouter()
	r.Run("0.0.0.0:" + port)
}
