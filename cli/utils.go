package main

import (
	"os"
	"os/user"
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
