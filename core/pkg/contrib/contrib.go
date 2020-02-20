// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/
package contrib

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllDatasets GET /datasets
func GetAllDatasets(c *gin.Context) {
	c.JSON(http.StatusOK, fetchAllDatasets())
}

// AddRegistryRequest Definition for AddRegistryRequest
type AddRegistryRequest struct {
	URL string `json:"url"`
}

// AddNewRegistry POST /datasets/registry
func AddNewRegistry(c *gin.Context) {
	var addRegistryRequest AddRegistryRequest
	c.BindJSON(&addRegistryRequest)
	addNewRegistry(addRegistryRequest.URL)
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"status": "Finished",
	})
}

// StreamCamera GET /camera
func StreamCamera(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		return true
	})
}

// UploadFile POST /files/upload
func UploadFile(c *gin.Context) {
	
}