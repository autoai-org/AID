// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utility

import (
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
)

// IsExists return true if path exists, otherwise returns false
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// GetHomeDir returns current user's home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

// IsRoot returns true if current user iis root, otherwise returns false
func IsRoot() bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	if usr.Username == "root" {
		return true
	}
	return false
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

// IsPortOpen returns true if current port is open, otherwise returns false
func IsPortOpen(port string, timeout time.Duration) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), timeout)
	if conn != nil {
		conn.Close()
		return false
	}
	return true
}

// FindNextOpenPort returns the next open port starting from port
func FindNextOpenPort(port int) string {
	var hasFound = false
	var strPort string
	for ; !hasFound; port++ {
		strPort = strconv.Itoa(port)
		if IsPortOpen(strPort, defaultTimeout) {
			hasFound = true
		}
	}
	return strPort
}

// ReadFileContent returns the content of filename
func ReadFileContent(filename string) string {
	var content string
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		content = "Read " + filename + "Failed!"
	} else {
		content = string(byteContent)
	}
	return content
}

// IsStringInSlice returns true if the string is in the slice(list).
func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsPathExists return true if the path exists, otherwise returns false
func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CreateFolderIfNotExist first checks if the @param folderPath exists, if not, it will create one
func CreateFolderIfNotExist(folderPath string) {
	exist, err := IsPathExists(folderPath)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Print("error when creating " + folderPath + ": " + err.Error())
	}
	if !exist {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Print("error when creating " + folderPath + ": " + err.Error())
		}
	}
}
