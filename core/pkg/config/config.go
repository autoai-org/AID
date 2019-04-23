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
package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	raven "github.com/getsentry/raven-go"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/unarxiv/cvpm/pkg/entity"
	"github.com/unarxiv/cvpm/pkg/utility"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type CvpmConfig struct {
	Local        Local               `toml:"local"`
	Repositories []entity.Repository `toml:"repository"`
}

type Local struct {
	LocalFolder string
	Pip         string
	Python      string
}

// Read returns current config object stored in the config.toml
func Read() CvpmConfig {
	var config CvpmConfig
	homepath, _ := homedir.Dir()
	configFile := filepath.Join(homepath, "cvpm", "config.toml")
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return config
	}
	return config
}

// ReadClient returns config object stored in a client directory named config.toml
func ReadClient(clientDir string) CvpmConfig {
	var config CvpmConfig
	configFile := filepath.Join(clientDir, "cvpm", "config.toml")
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	return config
}

// Write writes config object into config.toml file
func Write(config CvpmConfig) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	homepath, _ := homedir.Dir()
	configFile := filepath.Join(homepath, "cvpm", "config.toml")
	err := ioutil.WriteFile(configFile, buf.Bytes(), 0644)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
}

func getDefaultConfig() CvpmConfig {
	homePath, _ := homedir.Dir()
	cvpmPath := filepath.Join(homePath, "cvpm")
	var defaultLocal = Local{LocalFolder: cvpmPath, Pip: "pip", Python: "python"}
	var defaultCVPMConfig = CvpmConfig{Local: defaultLocal, Repositories: []entity.Repository{}}
	return defaultCVPMConfig
}

func createFolderIfNotExist(folderPath string) {
	exist, err := utility.IsPathExists(folderPath)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	if !exist {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Fatal(err)
		}
	}
}

func createFileIfNotExist(filePath string) {
	exist, err := utility.IsPathExists(filePath)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	if !exist {
		f, err := os.Create(filePath)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Fatal(err)
		}
		defer f.Close()
	}
}

// Validate init all params in config file
func Validate() {
	if !utility.IsRoot() {
		homepath := utility.GetHomeDir()
		// Validate CVPM Path
		cvpmPath := filepath.Join(homepath, "cvpm")
		createFolderIfNotExist(cvpmPath)
		// check if cvpm.toml file exists
		cvpmConfigToml := filepath.Join(homepath, "cvpm", "config.toml")
		if _, err := os.Stat(cvpmConfigToml); os.IsNotExist(err) {
			createFileIfNotExist(cvpmConfigToml)
			// config file not exists, write default to it
			Write(getDefaultConfig())
		}
		// create logs folder
		logsFolder := filepath.Join(cvpmPath, "logs")
		createFolderIfNotExist(logsFolder)
		// create webui folder
		webuiFolder := filepath.Join(cvpmPath, "webui")
		createFolderIfNotExist(webuiFolder)
		// create database folder
		databaseFolder := filepath.Join(cvpmPath, "database")
		createFolderIfNotExist(databaseFolder)
		createFileIfNotExist(filepath.Join(databaseFolder, "cvpm-database.db"))
		// check if system log file exists
		cvpmLogPath := filepath.Join(cvpmPath, "logs", "system.log")
		createFileIfNotExist(cvpmLogPath)
		// check if package log file exists
		cvpmPackageLogPath := filepath.Join(cvpmPath, "logs", "package.log")
		createFileIfNotExist(cvpmPackageLogPath)
	}
	utility.InitRaven()
}

// GetLogLocation returns the logs file location (the directory, not the file)
// Typically, there are system.log & package.log files in the directory
func GetLogLocation() string {
	homepath, _ := homedir.Dir()
	cvpmLogPath := filepath.Join(homepath, "cvpm", "logs")
	return cvpmLogPath
}
