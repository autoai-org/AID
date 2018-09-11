package main

import(
	"log"
	"path/filepath"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type cvpmConfig struct {
	Local local `toml:"local"`
}

type local struct {
	LocalFolder string
	Pip string
	Python string
}

var apiURL = "http://192.168.1.12:8080/"

func readConfig () cvpmConfig {
	var config cvpmConfig
	homepath, _ := homedir.Dir()
	configFile := filepath.Join(homepath, "cvpm", "config.toml")
	if _, err := toml.DecodeFile(configFile, &config); err!=nil {
		log.Fatal(err)
		return config
	}
	return config
}
