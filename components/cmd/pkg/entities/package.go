// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aid/components/cmd/pkg/storage"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

var logger = utilities.NewDefaultLogger()

// Package defines basic package information
type Package struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	LocalPath string    `db:"localpath"`
	Vendor    string    `db:"vendor"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	RemoteURL string    `db:"remote_url"`
}

// Pretrained defines the basic structure of pretrained file,
// it do not need to be stored in database
// and therefore has no `db` bindings.
type Pretrained struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

// Pretraineds is the collection/list of pretrained files
// this definition is used for toml parser
type Pretraineds struct {
	Model []Pretrained `toml:"model"`
}

// TableName defines the tablename in database
func (p *Package) TableName() string {
	return "package"
}

// PK defines the primary key of Package
func (p *Package) PK() string {
	return "id"
}

// PackageConfig is the toml interface as in aid.toml
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

// Save stores package into database
func (p *Package) Save() error {
	p.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(p)
}

// FetchPackages returns all packages
func FetchPackages() []Package {
	packagesPointers := make([]*Package, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&packagesPointers)
	packages := make([]Package, len(packagesPointers))
	for i := range packagesPointers {
		packages[i] = *packagesPointers[i]
	}
	return packages
}

// GetPackage returns a single package by given vendorName and packageName
func GetPackage(vendorName string, packageName string) Package {
	reqPackage := Package{Vendor: vendorName, Name: packageName}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&reqPackage)
	utilities.CheckError(err, "Cannot fetch package object with id"+id)
	return reqPackage
}

// LoadPretrainedsFromConfig reads the config string and returns the objects
func LoadPretrainedsFromConfig(tomlString string) Pretraineds {
	var pretraineds Pretraineds
	_, err := toml.Decode(tomlString, &pretraineds)
	utilities.CheckError(err, "Cannot Load Solvers")
	return pretraineds
}
