// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"net/http"
	"path/filepath"

	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

func getExtensionResult(c *gin.Context) {
	logsDir := utilities.GetBasePath()
	logsDir = filepath.Join(logsDir, "logs", "plugins", "dive")
	fileContent, _ := utilities.ReadFileContent(filepath.Join(logsDir, "1.txt"))
	c.JSON(http.StatusOK, logContent{
		Content: fileContent,
	})
}
