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
	requiredFolders := [3]string{"logs", "models", "plugins"}
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
