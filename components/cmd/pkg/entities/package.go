// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"time"
)

// Package defines basic package information
type Package struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	LocalPath string    `db:"localpath"`
	Vendor    string    `db:"vendor"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// TableName defines the tablename in database
func (p *Package) TableName() string {
	return "package"
}

// PK defines the primary key of Package
func (p *Package) PK() string {
	return "id"
}

// PackageConfig is the toml interface as in cvpm.toml
type PackageConfig struct {
	Solvers []Solver
	Package Package
}

// LoadPackageFromConfig reads the config string and returns a PackageConfig
func LoadPackageFromConfig(tomlString string) PackageConfig {
	var packageConfig PackageConfig
	_, err := toml.Decode(tomlString, &packageConfig)
	utilities.CheckError(err, "Cannot Load Solvers from toml string, please check its syntax!")
	return packageConfig
}
