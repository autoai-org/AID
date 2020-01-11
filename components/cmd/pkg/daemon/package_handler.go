// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

// getPackages returns all packages
// GET /packages -> returns packages
func getPackages(c *gin.Context) {
	packages := entities.FetchPackages()
	c.JSON(http.StatusOK, packages)
}

// getSolvers returns all solvers
// GET /solvers -> returns all solvers
func getSolvers(c *gin.Context) {
	solvers := entities.FetchSolvers()
	c.JSON(http.StatusOK, solvers)
}

// installPackage performs installation
// PUT /packages
func installPackage(c *gin.Context) {
	var request installPackageRequest
	c.BindJSON(&request)
	targetPath := filepath.Join(utilities.GetBasePath(), "models")
	err := runtime.InstallPackage(request.RemoteURL, targetPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, messageResponse{
			Msg: err.Error(),
		})
	}
	c.JSON(http.StatusOK, messageResponse{
		Msg: "submitted success",
	})
}


// buildPackages : PUT /packages/:packageName/solvers/:solverName/images -> build a new image
func buildPackageImage(c *gin.Context) {
	dockerClient := runtime.NewDockerRuntime()
	packageFolder := filepath.Join(utilities.DefaultConfig.Read("install_path"), c.Param("packageName"))
	tomlString := utilities.ReadFileContent(filepath.Join(packageFolder, "./cvpm.toml"))
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	solverName := c.Param("solverName")
	buildSolvers := 1
	var imageName string
	if solverName == "*" {
		// TODO handle multi builds return value
		buildSolvers = len(packageInfo.Solvers)
		for _, solver := range packageInfo.Solvers {
			imageName := packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solver.Name
			// Check if docker file exists
			if utilities.ReadFileContent("./docker_"+solver.Name) == "Read "+"./docker_"+solver.Name+" Failed!" {
				runtime.RenderDockerfile(solver.Name, "./docker_"+solver.Name)
			}
			dockerClient.Build(strings.ToLower(imageName), "./docker_"+solver.Name)
		}
	} else {
		imageName = packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solverName
		// Check if docker file exists
		if utilities.ReadFileContent("./docker_"+solverName) == "Read "+"./docker_"+solverName+" Failed!" {
			runtime.RenderDockerfile(solverName, "./docker_"+solverName)
		}
		dockerClient.Build(strings.ToLower(imageName), "./docker_"+solverName)
	}
	c.JSON(http.StatusOK, gin.H{
		"build_total_num": buildSolvers,
		"log_file_path":   filepath.Join(utilities.DefaultConfig.Read("logs_path"), "builds", imageName),
	})
}
