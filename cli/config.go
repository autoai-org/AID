package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}
}
