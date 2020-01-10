// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	r := gin.Default()
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	// All get requests
	r.GET("/logs", getLogs)
	r.GET("/packages", getPackages)
	r.GET("/logs/:logid", getLog)
	// all post/put requests
	r.PUT("/packages/:packageName/solvers/:solverName/images", buildPackageImage)

	// websocket
	r.GET("/socket/:logid", func(c *gin.Context) {
		serveWs(c)
	})
	return r
}
