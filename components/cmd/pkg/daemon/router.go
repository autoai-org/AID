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
	r.GET("/packages", getPackages)
	r.GET("/logs", getLogs)
	r.GET("/logs/:logid", getLog)
	r.GET("/solvers", getSolvers)
	// all post/put requests
	r.PUT("/vendors/:vendorName/packages/:packageName/solvers/:solverName/images", buildPackageImage)
	r.POST("/packages", installPackage)
	// all delete requests
	r.DELETE("/logs/:logid", deleteLog)
	// websocket
	r.GET("/socket/:logid", func(c *gin.Context) {
		serveWs(c)
	})
	return r
}
