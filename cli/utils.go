package main

import (
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
	} else {
		return false
	}
}

func getDirSizeMB(path string) float64 {
	var dirSize int64 = 0
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
	} else {
		return true
	}
}

func findNextOpenPort(port int) string {
	var hasFound bool = false
	var str_port string
	for ; !hasFound; port++ {
		str_port = strconv.Itoa(port)
		if isPortOpen(str_port, defaultTimeout) {
			hasFound = true
		}
	}
	return str_port
}
