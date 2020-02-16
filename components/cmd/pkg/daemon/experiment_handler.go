// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"net/http"
	"path/filepath"

	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

func getDatasets(c *gin.Context) {
	datasets := entities.FetchAllDatasets()
	for i := range datasets {
		datasets[i].Size = utilities.GetDirSizeMB(datasets[i].LocalPath)
	}
	c.JSON(http.StatusOK, datasets)
}

func uploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	datasetPath := filepath.Join(utilities.GetBasePath(), "datasets", file.Filename)

	c.SaveUploadedFile(file, datasetPath)
	c.JSON(http.StatusOK, datasetPath)
}

func addDataset(c *gin.Context) {
	var request newDatasetRequest
	c.BindJSON(&request)
	status := "Compressed"
	var err error
	targetFolder := filepath.Join(utilities.GetBasePath(), "datasets", request.Name)
	if request.Uncompress {
		status = "Uncompressed"
		err = utilities.Uncompress(request.DatasetPath, targetFolder)
	}
	newDataset := entities.Dataset{Name: request.Name, LocalPath: targetFolder, Status: status}
	err = newDataset.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, messageResponse{
			Code: 400,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, newDataset)
	}
}

func addExperiment(c *gin.Context) {

}
