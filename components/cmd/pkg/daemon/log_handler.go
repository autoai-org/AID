// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// getLogs returns all logs
// Get /logs
func getLogs(c *gin.Context) {
	logs := entities.FetchLogs()
	c.JSON(http.StatusOK, logs)
}

// getlog returns the specified log
// Get /logs/:logid
func getLog(c *gin.Context) {
	requestedFilename := entities.GetLog(c.Param("logid")).Filepath
	fileContent := utilities.ReadFileContent(requestedFilename)
	c.JSON(http.StatusOK, logContent{
		Content: fileContent,
	})
}

// deleteLog deletes the specified log and its database entry, a real deletion is performed
// DELETE /logs/:logid
func deleteLog(c *gin.Context) {
	requestedLog := entities.GetLog(c.Param("logid"))
	os.Remove(requestedLog.Filepath)
	err := entities.DeleteLog(requestedLog.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageResponse{
			Code: 400,
			Msg: err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, messageResponse{
			Code: 200,
			Msg: requestedLog.ID,
		})
	}
}