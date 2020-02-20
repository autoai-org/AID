// Package contrib defines extra stuff for cvpm
//Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/
package contrib

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/utility"
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

// UploadFile POST /files/upload/:type
func UploadFile(c *gin.Context) {
	var (
		src  multipart.File
		file *multipart.FileHeader
		dist *os.File
		err  error
	)
	if file, err = c.FormFile(FileField); err != nil {
		parseFormFail(c)
		return
	}
	fileType := c.Param("type")
	extname := strings.ToLower(path.Ext(file.Filename))
	if src, err = file.Open(); err != nil {
		log.Print(err)
	}
	defer src.Close()
	md5string := uniuri.New()
	fileName := md5string + extname
	distPath := path.Join(config.Read().Local.TmpFolder, fileType, fileName)
	distFolder := path.Join(config.Read().Local.TmpFolder, fileType)
	utility.CreateFolderIfNotExist(distFolder)
	if dist, err = os.Create(distPath); err != nil {
		log.Print(err)
	}
	defer dist.Close()
	io.Copy(dist, src)
	// set filetype to dataset since all uploaded large file should be 'dataset'
	insertFileObjectToDB(fileName, distPath, fileType, file.Size, "compressed")
	c.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
	})
}

// QueryFilesList GET /contrib/files/list
func QueryFilesList(c *gin.Context) {
	filesList := queryFiles()
	c.JSON(http.StatusOK, gin.H{
		"result": filesList,
	})
}

// QueryTrainsList returns all train records
func QueryTrainsList(c *gin.Context) {
	trainLists := fetchAllTrainRecord()
	c.JSON(http.StatusOK, gin.H{
		"result": trainLists,
	})
}

// UncompressFile POST /contrib/files/uncompress/:id
func UncompressFile(c *gin.Context) {
	result := uncompressFile(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// QueryAnnotationFile returns the content of requested file
func QueryAnnotationFile(c *gin.Context) {
	result := getAnnotationFile(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"annotation": result,
	})
}

// trainRecordRequest is the request definition for AddTrainRecord
type trainRecordRequest struct {
	FilePath string `json:"filePath"`
	Vendor   string `json:"vendor"`
	Package  string `json:"package"`
	Solver   string `json:"solver"`
	RayId    string `json:"rayID"`
}

// AddTrainRecords receives feedback from training process and stores the result
func AddTrainRecord(c *gin.Context) {
	var addTrainRecordRequest trainRecordRequest
	c.BindJSON(&addTrainRecordRequest)
	// get DataId by Data Path
	dataId := getFileByDataPath(addTrainRecordRequest.FilePath).ObjectID
	success := insertTrainRecordToDB(addTrainRecordRequest.RayId, addTrainRecordRequest.Vendor, addTrainRecordRequest.Package, addTrainRecordRequest.Solver, dataId)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "Success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "Failed",
		})
	}
}
