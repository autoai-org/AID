// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"net/http"

	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
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

func addWebhook(c *gin.Context) {
	var request newWebhookRequest
	c.BindJSON(&request)
	webhook := entities.Webhook{
		RemoteURL: request.PayloadURL,
		Status:    request.Status,
	}
	err := webhook.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, messageResponse{
			Code: 500,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, webhook)
	}
}

func fetchAllWebhooks(c *gin.Context) {
	webhooks := entities.FetchWebhooks()
	c.JSON(http.StatusOK, webhooks)
}

func fetchRunningErrors(c *gin.Context) {
	c.JSON(http.StatusOK, entities.RunningErrors)
}
