// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"log"
	"os"
	"os/user"
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
