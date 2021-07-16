// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package handlers

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-sysinfo"
	"github.com/gin-gonic/gin"
)

type HostInfoResponse struct {
	Error        string `json:"error"`
	Memory       string `json:"memory"`
	CPUTime      string `json:"cpu"`
	HostName     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

func PingHandler(c *gin.Context) {
	hostInfo, err := sysinfo.Host()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HostInfoResponse{Error: err.Error()})
	}
	memoryInfo, err := hostInfo.Memory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HostInfoResponse{Error: err.Error()})
	}
	cpuInfo, err := hostInfo.CPUTime()
	architecture := hostInfo.Info().Architecture
	os := hostInfo.Info().OS
	hostName := hostInfo.Info().Hostname
	if err != nil {
		c.JSON(http.StatusInternalServerError, HostInfoResponse{Error: err.Error()})
	}
	c.JSON(http.StatusOK, HostInfoResponse{
		OS:           os.Name + "@" + os.Build,
		Memory:       fmt.Sprint(memoryInfo.Available/(1024*1024*1024)) + "GB/" + fmt.Sprint(memoryInfo.Total/(1024*1024*1024)) + "GB",
		CPUTime:      cpuInfo.System.String(),
		Architecture: architecture,
		HostName:     hostName,
	})
}

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}
