// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/requests"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"path/filepath"
	"strings"
	"time"
)

var logger = utilities.NewDefaultLogger("./logs/system.log")

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

// InstallPackage fetches remote content to target folder
func InstallPackage(remoteURL string, targetFolder string) {
	// Check type of RemoteURL
	var remoteType string
	var pack Package
	if strings.HasPrefix(remoteURL, "https://github.com") {
		remoteType = "Git"
	} else if strings.HasPrefix(remoteURL, "git://") {
		remoteType = "Git"
	} else {
		remoteType = "Registry"
	}
	// Fetch Remote Content
	switch remoteType {
	case "Git":
		git := requests.NewGitClient()
		localFolderName := strings.Split(remoteURL, "/")
		vendorName := localFolderName[len(localFolderName)-2]
		repoName := localFolderName[len(localFolderName)-1]
		targetSubFolder := filepath.Join(targetFolder, vendorName, repoName)
		pack = Package{Name: repoName,
			LocalPath: targetSubFolder,
			Vendor:    vendorName,
			Status:    "installed",
			RemoteURL: remoteURL}
		git.Clone(remoteURL, targetSubFolder)
	case "Registry":
		logger.Error("Unsupported Remote Type.")
	default:
		logger.Error("Unsupported Remote Type.")
	}
	// Save package into database for future use
	err := pack.Save()
	utilities.CheckError(err, "Cannot save package into database!")
	// All is well
	logger.Info("Installation Finished!")
}

// Save stores package into database
func (p *Package) Save() error {
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(p)
}
