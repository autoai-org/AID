// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package configuration

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aid/internal/utilities"
)

// LoadSolversFromConfig reads the config string and returns the objects
func LoadSolversFromConfig(tomlString string) PackageConfig {
	var packageConfig PackageConfig
	_, err := toml.Decode(tomlString, &packageConfig)
	if err != nil {
		utilities.Formatter.Error("Cannot load solvers from toml string: " + err.Error())
	}
	return packageConfig
}

// LoadPackageFromConfig reads the config string and returns a PackageConfig
func LoadPackageFromConfig(tomlString string) PackageConfig {
	var packageConfig PackageConfig
	_, err := toml.Decode(tomlString, &packageConfig)
	if err != nil {
		utilities.Formatter.Error("Cannot load solvers from toml string: " + err.Error())
	}
	return packageConfig
}

// LoadPretrainedsFromConfig reads the pretrained config string and returns a Pretraineds object
func LoadPretrainedsFromConfig(tomlString string) Pretraineds {
	var pretraineds Pretraineds
	_, err := toml.Decode(tomlString, &pretraineds)
	if err != nil {
		utilities.Formatter.Error("Cannot load pretraineds from toml string: " + err.Error())
	}
	return pretraineds
}
