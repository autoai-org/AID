// Package contrib defines extra stuff for cvpm
//Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*This file handles third party contributions and libraries.*/

package contrib

import (
	"github.com/gen2brain/go-unarr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/entity"
	"github.com/unarxiv/cvpm/pkg/utility"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const FileField = "file"

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

func insertFileObjectToDB(filename string, filepath string, filetype string, filesize int64, status string) {
	objectID := uuid.New().String()
	fileObject := entity.FileObject{
		ObjectID:   objectID,
		Name:       filename,
		Size:       filesize,
		Filepath:   filepath,
		Type:       filetype,
		UploadedAt: time.Now().Format(time.RFC3339),
		Status:     status,
	}
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

func getFileObjectByID(objectId string) entity.FileObject {
	var fileObject entity.FileObject
	sess := GetDBInstance()
	fileCollection := sess.Collection("file")
	res := fileCollection.Find("ObjectId", objectId)
	err := res.One(&fileObject)
	if err != nil {
		log.Println(err)
	}
	return fileObject
}

func getFileByDataPath(filePath string) entity.FileObject {
	var fileObject entity.FileObject
	sess := GetDBInstance()
	fileCollection := sess.Collection("file")
	res := fileCollection.Find("Filepath", filePath)
	err := res.One(&fileObject)
	if err != nil {
		log.Println(err)
	}
	return fileObject
}

func setFileStatus(objectID string, status string) {
	sess := GetDBInstance()
	q := sess.Update("file").Set("Status", status).Where("ObjectId = ?", objectID)
	_, err := q.Exec()
	if err != nil {
		log.Println(err)
	}
}

func uncompressFile(objectID string) bool {
	setFileStatus(objectID, "uncompressing")
	fileObject := getFileObjectByID(objectID)
	archiveFile, err := unarr.NewArchive(fileObject.Filepath)
	if err != nil {
		setFileStatus(objectID, "uncompressed")
		log.Println(err)
	}
	fileHash := strings.TrimSuffix(fileObject.Name, filepath.Ext(fileObject.Name))
	distPath := path.Join(config.Read().Local.TmpFolder, fileObject.Type, fileHash)
	err = archiveFile.Extract(distPath)
	if err != nil {
		setFileStatus(objectID, "uncompressed")
		log.Println(err)
	}
	defer archiveFile.Close()
	setFileStatus(objectID, "uncompressed")
	return true
}

// getAnnotationFile returns the content of JSON annotation file
func getAnnotationFile(objectID string) string {
	fileObject := getFileObjectByID(objectID)
	fileHash := strings.TrimSuffix(fileObject.Name, filepath.Ext(fileObject.Name))
	annotationFilePath := path.Join(config.Read().Local.TmpFolder, fileObject.Type, fileHash, "annotations.json")
	return utility.ReadFileContent(annotationFilePath)
}
