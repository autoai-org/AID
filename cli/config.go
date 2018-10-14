package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/getsentry/raven-go"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type cvpmConfig struct {
	Local        local        `toml:"local"`
	Repositories []Repository `toml:repository`
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

func validateConfig() {
	// localConfig := readConfig()
	homepath, _ := homedir.Dir()
	// Validate CVPM Path
	cvpmPath := filepath.Join(homepath, "cvpm")
	exist, err := isPathExists(cvpmPath)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	if !exist {
		err := os.Mkdir(cvpmPath, os.ModePerm)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			log.Fatal(err)
		}
	}
}
