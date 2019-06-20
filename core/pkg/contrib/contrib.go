// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/
package contrib

import (
	"io"
	"os"
	"path"
	"strings"
	"log"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"github.com/gin-gonic/gin"
	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/entity"
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
	var (
		src multipart.File
		file *multipart.FileHeader
		dist *os.File
		err error
	)
	if file, err = c.FormFile(FILE_FIELD); err != nil {
		parseFormFail(c)
		return
	}
	
	extname := strings.ToLower(path.Ext(file.Filename))
	if src, err = file.Open(); err != nil {
		log.Print(err)
	}
	defer src.Close()
	hash := md5.New()
	io.Copy(hash, src)
	md5string := hex.EncodeToString(hash.Sum([]byte("")))
	fileName := md5string + extname
	distPath := path.Join(config.Read().Local.TmpFolder, fileName)
	if dist, err = os.Create(distPath); err != nil {
		log.Print(err)
	}
	defer dist.Close()
	io.Copy(dist, src)
	c.JSON(http.StatusOK, gin.H{
		"hash":     md5string,
		"filename": fileName,
		"origin":   file.Filename,
		"size":     file.Size,
	})
}

// QueryFilesList GET /contrib/files/list
func QueryFilesList(c *gin.Context) {
	var filesList []entity.FileObject
	files := getFileLists(config.Read().Local.TmpFolder)
	for _, f := range files {
		stat, err := f.Stat()

		if err != nil {
			log.Print("")
		}

		filesList = append(filesList, entity.FileObject{
			Name: f.Name(),
			Size: stat.Size(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result":     filesList,
	})
}