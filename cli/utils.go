// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
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

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func isRoot() bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	if usr.Username == "root" {
		return true
	}
	return false
}

func getDirSizeMB(path string) float64 {
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

func isPortOpen(port string, timeout time.Duration) bool {
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), timeout)
	if conn != nil {
		conn.Close()
		return false
	}
	return true
}

func findNextOpenPort(port int) string {
	var hasFound = false
	var strPort string
	for ; !hasFound; port++ {
		strPort = strconv.Itoa(port)
		if isPortOpen(strPort, defaultTimeout) {
			hasFound = true
		}
	}
	return strPort
}

func readFileContent(filename string) string {
	var content string
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		content = "Read " + filename + "Failed!"
	} else {
		content = string(byteContent)
	}
	return content
}

func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
