package runtime

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/requests"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"path/filepath"
	"strings"
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
	// Fetch Remote Content
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
	case "Registry":
		logger.Error("Unsupported Remote Type.")
	default:
		logger.Error("Unsupported Remote Type.")
	}
	// Generate Dockerfile
	configFile := filepath.Join(localFolder, "cvpm.toml")
	tomlString := utilities.ReadFileContent(configFile)
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	for _, solver := range packageInfo.Solvers {
		solver.Vendor = pack.Vendor
		solver.Package = pack.Name
		solver.Status = "Ready"
		solver.Save()
		RenderDockerfile(solver.Name, localFolder)
	}
	// Save package into database for future use
	err := pack.Save()
	utilities.CheckError(err, "Cannot save package into database!")
	// All is well
	logger.Info("Installation Finished!")
	return err
}
