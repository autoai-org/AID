// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/requests"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

// InstallPackage fetches remote content to target folder
func InstallPackage(remoteURL string, targetFolder string) error {
	// Check type of RemoteURL
	var remoteType string
	var pack entities.Package
	var localFolder string
	if strings.HasPrefix(remoteURL, "https://github.com") {
		remoteType = "Git"
	} else if strings.HasPrefix(remoteURL, "git://") {
		remoteType = "Git"
	} else {
		remoteType = "Registry"
	}
	// Fetch Remote Content (Source Code)
	switch remoteType {
	case "Git":
		git := requests.NewGitClient()
		localFolderName := strings.Split(remoteURL, "/")
		vendorName := localFolderName[len(localFolderName)-2]
		repoName := localFolderName[len(localFolderName)-1]
		targetSubFolder := filepath.Join(targetFolder, vendorName, repoName)
		absTargetSubFolder, _ := filepath.Abs(targetSubFolder)
		localFolder = absTargetSubFolder
		pack = entities.Package{Name: repoName,
			LocalPath: absTargetSubFolder,
			Vendor:    vendorName,
			Status:    "installed",
			RemoteURL: remoteURL}
		git.Clone(remoteURL, absTargetSubFolder)
		os.Remove(filepath.Join(localFolder, ".git"))
	case "Registry":
		logger.Error("Unsupported Remote Type.")
	default:
		logger.Error("Unsupported Remote Type.")
	}
	// Fetch Remote Content (Model File)
	// If pretrained.toml exists, then load the pretrained file
	pretrainedFile := filepath.Join(localFolder, "pretrained.toml")
	if utilities.IsExists(pretrainedFile) {
		pretrainedTomlString, err := utilities.ReadFileContent(pretrainedFile)
		fmt.Println(pretrainedTomlString)
		utilities.CheckError(err, "Cannot open file"+pretrainedFile)
		pretrainedInfo := entities.LoadPretrainedsFromConfig(pretrainedTomlString)
		utilities.Formatter.Progress("Downloading pretrained models...")
		for _, pretrained := range pretrainedInfo.Models {
			utilities.Formatter.Progress("Downloading " + pretrained.Name + "...")
			utilities.Download(pretrained.URL, filepath.Join(localFolder, "pretrained"))
		}
		utilities.Formatter.Success("Finished downloading pretrained files")
	}
	// Generate Dockerfile
	configFile := filepath.Join(localFolder, "aid.toml")
	tomlString, err := utilities.ReadFileContent(configFile)
	if err != nil {
		utilities.CheckError(err, "Cannot open file "+configFile)
	}
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	for _, solver := range packageInfo.Solvers {
		solver.Vendor = pack.Vendor
		solver.Package = pack.Name
		solver.Status = "Ready"
		solver.Save()
		RenderDockerfile(solver.Name, localFolder)
	}
	// Save package into database for future use
	err = pack.Save()
	utilities.CheckError(err, "Cannot save package into database!")
	// All is well
	logger.Info("Installation Finished!")
	return err
}
