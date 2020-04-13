// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/autoai-org/aid/components/cmd/pkg/daemon"
	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/runtime"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

// absoluteBuild builds the docker image without the current context information
// unlike "build", it relies on vendor/package/solver names
func absoluteBuild(vendorName string, packageName string, solverName string) {
	dockerClient := runtime.NewDockerRuntime()
	basePath := path.Join(utilities.GetBasePath(), "models", vendorName, packageName)
	tomlString, err := utilities.ReadFileContent(path.Join(basePath, "./aid.toml"))
	utilities.CheckError(err, path.Join(basePath, "./aid.toml"))
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	imageName := packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solverName
	_, err = utilities.ReadFileContent(path.Join(basePath, "./docker_"+solverName))
	if err != nil {
		runtime.RenderDockerfile(solverName, path.Join(basePath, "./docker_"+solverName))
	}
	dockerClient.Build(strings.ToLower(imageName), path.Join(basePath, "./docker_"+solverName))
}

func build(solverName string) {
	// Read Dockerfile
	dockerClient := runtime.NewDockerRuntime()
	// Read Base Image Name
	tomlString, err := utilities.ReadFileContent("./aid.toml")
	utilities.CheckError(err, "Cannot readfile "+"./aid.toml")
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	if solverName == "*" {
		for _, solver := range packageInfo.Solvers {
			imageName := packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solver.Name
			// Check if docker file exists
			_, err = utilities.ReadFileContent("./docker_" + solver.Name)
			if err != nil {
				runtime.RenderDockerfile(solver.Name, "./docker_"+solver.Name)
			}
			dockerClient.Build(strings.ToLower(imageName), "./docker_"+solver.Name)
		}
	} else {
		imageName := packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solverName
		// Check if docker file exists
		_, err = utilities.ReadFileContent("./docker_" + solverName)
		if err != nil {
			runtime.RenderDockerfile(solverName, "./docker_"+solverName)
		}
		dockerClient.Build(strings.ToLower(imageName), "./docker_"+solverName)
	}
}

func generate() {
	tomlFilePath := filepath.Join("./", "aid.toml")
	cvpmToml, err := utilities.ReadFileContent(tomlFilePath)
	if err != nil {
		utilities.CheckError(err, "Cannot open file "+tomlFilePath)
	}
	solvers := entities.LoadSolversFromConfig(cvpmToml)
	runtime.RenderRunnerTpl("./", solvers)
	runtime.GenerateDockerFiles("./")
}

func startServer(port string) {
	daemon.RunServer(port, filepath.Join("./", "keys", "server.crt"), filepath.Join("./", "keys", "server.key"))
}

func installPackage(remoteURL string) {
	targetPath := filepath.Join(utilities.GetBasePath(), "models")
	runtime.InstallPackage(remoteURL, targetPath)
}

func createSolverContainer(imageName string, hostPort string) {
	dockerClient := runtime.NewDockerRuntime()
	image := entities.GetImagebyName(imageName)
	resp, err := dockerClient.Create(image.ID, hostPort)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Created successfully with the id " + resp.ID)
	}
}
