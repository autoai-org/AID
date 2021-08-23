package handlers

import (
	"io"
	"path/filepath"
	"time"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/gin-gonic/gin"
)

func StreamLogHandler(c *gin.Context) {
	logUID := c.Param("logUID")
	logPath := filepath.Join(utilities.GetBasePath(), "logs", "builds", logUID)
	filePeriod := 10 * time.Second
	clientGone := c.Writer.CloseNotify()
	fileTicker := time.NewTicker(filePeriod)
	var lastMod time.Time
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case <-fileTicker.C:
			var p []byte
			p, lastMod, err := utilities.ReadFileIfModified(logPath, lastMod)
			if err != nil {
				utilities.Formatter.Info("New errors at " + lastMod.String())
				c.SSEvent("failed", err.Error())
			} else {
				utilities.Formatter.Info("New logs at " + lastMod.String())
				c.SSEvent("success", string(p))
			}
			return true
		}
	})
}
