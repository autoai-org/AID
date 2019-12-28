// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/requests"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"strings"
	"time"
	"path/filepath"
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
	if strings.HasPrefix(remoteURL, "https://github.com") {
		remoteType = "Git"
	} else if strings.HasPrefix(remoteURL, "git://") {
		remoteType = "Git"
	} else {
		remoteType = "Registry"
	}
	switch remoteType {
	case "Git":
		git := requests.NewGitClient()
		localFolderName := strings.Split(remoteURL, "/")
		vendorName := localFolderName[len(localFolderName)-2]
		repoName := localFolderName[len(localFolderName)-1]
		targetSubFolder := filepath.Join(targetFolder, vendorName, repoName)
		git.Clone(remoteURL, targetSubFolder)
	case "Registry":
		logger.Error("Unsupported Remote Type.")
	default:
		logger.Error("Unsupported Remote Type.")
	}
}
