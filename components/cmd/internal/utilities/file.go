// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

// GetHomeDir returns current user's home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

// IsExists returns true if the folder or file exists, false otherwise.
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

// GetBasePath returns the base filepath for aid
func GetBasePath() string {
	vendorDir := filepath.Join(GetHomeDir(), ".autoai")
	CreateFolderIfNotExist(vendorDir)
	targetDir := filepath.Join(vendorDir, "aid")
	return targetDir
}

// WriteContentToFile writes contents to file, it will create new file if not exists
func WriteContentToFile(filepath string, content string) error {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		_, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	contentInByte := []byte(content)
	err = ioutil.WriteFile(filepath, contentInByte, 0644)
	if err != nil {
		logger.Error("Cannot write to file " + filepath)
		return err
	}
	return nil
}

// GetRemoteFile returns a string that is included in a remote file
func GetRemoteFile(url string) string {
	resp, err := http.Get(url)
	ReportError(err, "Cannot fetch remote file..., Please check your internet connection!")
	body, err := ioutil.ReadAll(resp.Body)
	ReportError(err, "Cannot read from response..., Please check your internet connection!")
	return string(body)
}

// ReadFileIfModified will return filecontent in byte mode if file has been modified
func ReadFileIfModified(filename string, lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	fileindicator := []byte(filename + "\n")
	msg := append(fileindicator, p...)
	return msg, fi.ModTime(), nil
}

// IsFileExists is a shortcut to check if file exists
func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFileContent returns the content of filename
func ReadFileContent(filename string) (string, error) {
	var content string
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		content = "Read " + filename + " Failed!"
	} else {
		content = string(byteContent)
	}
	return content, err
}

// GetPackageFolder returns the folder to the specific package
func GetPackageFolder(vendorName string, packageName string) string {
	return filepath.Join(GetBasePath(), MODELSFOLDER, vendorName, packageName)
}

// GetFolder returns the parent folder of models, plugins, temp and datasets
func GetFolder(folderName string) string {
	return filepath.Join(GetBasePath(), folderName)
}
