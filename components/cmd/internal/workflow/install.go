// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package workflow

import (
	"context"
	"path/filepath"
	"strings"

	ent "github.com/autoai-org/aid/ent/generated"
	"github.com/autoai-org/aid/internal/configuration"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/runtime/requests"
	"github.com/autoai-org/aid/internal/utilities"
)

// PullPackageSource tried to download the source code of the file from remote
// address, it now supports github and other git-based server.
func PullPackageSource(remoteURL string) {
	targetPath := filepath.Join(utilities.GetBasePath(), "models")
	var remoteType string
	var installedRepository *ent.Repository
	imageID := utilities.GenerateUUIDv4()
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
		targetSubFolder := filepath.Join(targetPath, vendorName, repoName)
		absTargetSubFolder, _ := filepath.Abs(targetSubFolder)
		err := git.Clone(remoteURL, absTargetSubFolder)
		if err == nil {
			installedRepository, err = database.NewDefaultDB().Repository.Create().SetName(repoName).SetLocalpath(absTargetSubFolder).SetVendor(vendorName).SetUID(imageID).SetRemoteURL(remoteURL).SetStatus("Source Code Installed").Save(context.Background())
			utilities.ReportError(err, "cannot save new package to database")
		}
	case "Registry":
		utilities.Formatter.Error("Registry will be supported in the near future.")
	default:
		utilities.Formatter.Error("Unsupported Remote Type.")
	}
	tomlString, err := utilities.ReadFileContent(filepath.Join(installedRepository.Localpath, "aid.toml"))
	utilities.ReportError(err, "cannot parse aid.toml file")
	packageConfig := configuration.LoadPackageFromConfig(tomlString)
	for _, solver := range packageConfig.Solvers {
		_, err = database.NewDefaultDB().Solver.Create().SetUID(utilities.GenerateUUIDv4()).SetName(solver.Name).SetRepository(installedRepository).SetClass(solver.Class).SetStatus("Code Installed").Save(context.Background())
		utilities.ReportError(err, "cannot save new solver to database")
	}
	pretrainedTomlString, err := utilities.ReadFileContent(filepath.Join(installedRepository.Localpath, "pretrained.toml"))
	pretraineds := configuration.LoadPretrainedsFromConfig(pretrainedTomlString)
	for _, pretrained := range pretraineds.Models {
		err = utilities.Download(pretrained.URL, filepath.Join(installedRepository.Localpath, "pretrained"))
		if err != nil {
			utilities.Formatter.Error("Cannot Download " + pretrained.URL + ": " + err.Error())
		}
	}
	if err == nil {
		utilities.Formatter.Info(installedRepository.Name + " installed successfully")
	}
}
