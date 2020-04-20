// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/getsentry/sentry-go"
)

// SystemConfig stores system level configuration, will be stored under $aid/config.toml
type SystemConfig struct {
	RemoteReport bool
}

// DefaultConfig is the instance shared by all modules
var DefaultConfig *SystemConfig

// NewDefaultConfig returns the config object
func NewDefaultConfig() *SystemConfig {
	if DefaultConfig != nil {
		return DefaultConfig
	}
	DefaultConfig := ReadConfig()
	return DefaultConfig
}

// SaveConfig writes config to its required file path
func SaveConfig(config SystemConfig) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		CheckError(err, "Cannot encode config object")
	}
	configPath := filepath.Join(GetBasePath(), "config.toml")
	if err := WriteContentToFile(configPath, buf.String()); err != nil {
		CheckError(err, "Cannot write config object to "+configPath)
	}
}

// ReadConfig always read the config file
// If the file does not exist, it will return a default config
func ReadConfig() *SystemConfig {
	configPath := filepath.Join(GetBasePath(), "config.toml")
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return &SystemConfig{RemoteReport: true}
	}
	tomlString, err := ReadFileContent(configPath)
	if err != nil {
		CheckError(err, "Cannot open file "+configPath)
	}
	if _, err := toml.Decode(tomlString, &DefaultConfig); err != nil {
		CheckError(err, "Cannot load config file!")
	}
	if DefaultConfig.RemoteReport {
		sentry.Init(sentry.ClientOptions{
			Dsn: "https://e6770124b98e44cfafa9d0e67e2d3650@sentry.io/1919735",
		})
	}
	return DefaultConfig
}
