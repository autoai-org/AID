// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// AgentDaemon is a lightweight daemon server that only handles limited tasks
// It is separated from the
package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

func runAgentServer() {

}

// An HTTP Server that supports receiving large size files
// Its not included in RPC Server because http outperformed RPC in receiving files,
// as there is no encode/decode overhead.
func startImageReceiver() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		targetPath := filepath.Join(utilities.GetBasePath(), "temp", file.Filename+".aidimg")
		c.SaveUploadedFile(file, targetPath)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Successfully Uploaded",
		})
	})
	fmt.Println("Listening on :10591, waiting for upload requests...")
	router.Run(":10591")
}
