// Package contrib defines extra stuff for cvpm
//Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/

package contrib

import (
	"io/ioutil"
	"log"
	"time"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/unarxiv/cvpm/pkg/entity"

)

const FILE_FIELD = "file"

func parseFormFail(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Can not parse form",
	})
}

func getFileLists(filepath string) []os.FileInfo {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Print(err)
	}
	return files
}

func insertFileObjectToDB(filename string, filepath string, filetype string, filesize int64) {
	fileObject := entity.FileObject{Name: filename, Size: filesize, Filepath: filepath, Type: filetype, UploadedAt: time.Now().Format(time.RFC3339)}
	sess := GetDBInstance()
	fileCollection := sess.Collection("file")
	err := fileCollection.InsertReturning(&fileObject)
	if err != nil {
		panic(err)
	}
	CloseDB(sess)
}

func queryFiles() []entity.FileObject {
	sess := GetDBInstance()
	fileCollection := sess.Collection("file")
	res := fileCollection.Find()
	var fileObjects []entity.FileObject

	err := res.All(&fileObjects)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
	CloseDB(sess)
	return fileObjects
}