// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
)

// Package defines basic package information
type Package struct {
	ID        string
	Name      string
	LocalPath string
	Vendor    string
}

// PackageConfig is the toml interface as in cvpm.toml
type PackageConfig struct {
	Solvers []Solver
	Package Package
}

// LoadPackageFromConfig reads the config string and returns a PackageConfig
func LoadPackageFromConfig(tomlString string) PackageConfig {
	var packageConfig PackageConfig
	if _, err := toml.Decode(tomlString, &packageConfig); err != nil {
		logger.Fatal("Cannot Load Solvers frmom ")
	}
	return packageConfig
}
