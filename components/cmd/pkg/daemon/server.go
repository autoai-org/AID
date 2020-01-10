// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = utilities.NewDefaultLogger()

// beforeResponse set global header to enable cors and set response header
func beforeResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("aid-version", "1.0.0 @ dev")
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

// RunServer starts the https service
func RunServer(port string, sslcert string, sslkey string) {
	logger.Info("Initiating Service...")
	r := getRouter()
	var err error
	go func() {
		if utilities.IsExists(sslcert) && port == "443" {
			err = r.RunTLS("127.0.0.1:"+port, sslcert, sslkey)
		} else {
			err = r.Run("127.0.0.1:" + port)
		}
		utilities.CheckError(err, "Cannot Start Server")
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server...")
}
