// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
)

// GetHomeDir returns current user's home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

// IsExists returns true if the folder exists, false otherwise.
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// CreateFolderIfNotExist first checks if the @param folderPath exists, if not, it will create one
func CreateFolderIfNotExist(folderPath string) {
	exist := IsExists(folderPath)
	if !exist {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Print("error when creating " + folderPath + ": " + err.Error())
		}
	}
}

// InitFolders create all required folders under ~/.autoai/.aid
func InitFolders() {
	vendorDir := filepath.Join(GetHomeDir(), ".autoai")
	CreateFolderIfNotExist(vendorDir)
	targetDir := filepath.Join(vendorDir, ".aid")
	CreateFolderIfNotExist(targetDir)
	requiredFolders := [4]string{"logs", "models", "plugins", "datasets"}
	for _, each := range requiredFolders {
		CreateFolderIfNotExist(filepath.Join(targetDir, each))
	}
}

// GetBasePath returns the base filepath for aid
func GetBasePath() string {
	vendorDir := filepath.Join(GetHomeDir(), ".autoai")
	CreateFolderIfNotExist(vendorDir)
	targetDir := filepath.Join(vendorDir, ".aid")
	return targetDir
}

// GetDirSizeMB returns the size of path
func GetDirSizeMB(path string) float64 {
	var dirSize int64
	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}
		return nil
	}
	filepath.Walk(path, readSize)
	sizeMB := float64(dirSize) / 1024.0 / 1024.0
	return sizeMB
}
