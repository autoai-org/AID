// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*  This file handles config related tasks, include:
readConfig()
writeConfig()
validateConfig()
getDefaultConfig()

The config file is located at the home dir of current user, under ~/cvpm/config.toml
*/
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	raven "github.com/getsentry/raven-go"
	homedir "github.com/mitchellh/go-homedir"
)

type cvpmConfig struct {
	Local        local        `toml:"local"`
	Repositories []Repository `toml:"repository"`
}

type local struct {
	LocalFolder string
	Pip         string
	Python      string
}

var apiURL = "http://192.168.1.12:8080/"

func isPathExists(path string) (bool, error) {
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

func readConfig() cvpmConfig {
	var config cvpmConfig
	homepath, _ := homedir.Dir()
	configFile := filepath.Join(homepath, "cvpm", "config.toml")
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return config
	}
	return config
}

func readClientConfig(clientDir string) cvpmConfig {
	var config cvpmConfig
	configFile := filepath.Join(clientDir, "cvpm", "config.toml")
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func writeConfig(config cvpmConfig) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		log.Fatal(err)
	}
	homepath, _ := homedir.Dir()
	configFile := filepath.Join(homepath, "cvpm", "config.toml")
	err := ioutil.WriteFile(configFile, []byte(buf.String()), 0644)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
}

func getDefaultConfig() cvpmConfig {
	localPath, _ := homedir.Dir()
	cvpmPath := filepath.Join(localPath, "cvpm")
	var defaultLocal = local{LocalFolder: cvpmPath, Pip: "pip", Python: "python"}
	var defaultCVPMConfig = cvpmConfig{Local: defaultLocal, Repositories: []Repository{}}
	return defaultCVPMConfig
}

func createFolderIfNotExist(folderPath string) {
	exist, err := isPathExists(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createFileIfNotExist(filePath string) {
	exist, err := isPathExists(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}

func validateConfig() {
	if !isRoot() {
		homepath := getHomeDir()
		// Validate CVPM Path
		cvpmPath := filepath.Join(homepath, "cvpm")
		createFolderIfNotExist(cvpmPath)
		// check if cvpm.toml file exists
		cvpmConfigToml := filepath.Join(homepath, "cvpm", "config.toml")
		createFileIfNotExist(cvpmConfigToml)
		// create logs folder
		logsFolder := filepath.Join(cvpmPath, "logs")
		createFolderIfNotExist(logsFolder)
		// create webui folder
		webuiFolder := filepath.Join(cvpmPath, "webui")
		createFolderIfNotExist(webuiFolder)
		// check if system log file exists
		cvpmLogPath := filepath.Join(cvpmPath, "logs", "system.log")
		createFileIfNotExist(cvpmLogPath)
		// check if package log file exists
		cvpmPackageLogPath := filepath.Join(cvpmPath, "logs", "package.log")
		createFileIfNotExist(cvpmPackageLogPath)
	}
}

func getLogsLocation() string {
	homepath, _ := homedir.Dir()
	cvpmLogPath := filepath.Join(homepath, "cvpm", "logs")
	return cvpmLogPath
}
