// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getConfigs(c *gin.Context) {
	configs := utilities.ReadConfig()
	c.JSON(http.StatusOK, configs)
}

func updateConfigs(c *gin.Context) {
	var config utilities.SystemConfig
	err := c.BindJSON(&config)
	utilities.SaveConfig(config)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageResponse{
			Code: 400,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, config)
	}
}
