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
		Code: 200,
		Msg:  "submitted success",
	})
}

// buildPackages build a new image
// PUT /vendor/:vendorName/packages/:packageName/solvers/:solverName/images
func buildPackageImage(c *gin.Context) {
	dockerClient := runtime.NewDockerRuntime()
	packageFolder := filepath.Join(utilities.GetBasePath(), "models", c.Param("vendorName"), c.Param("packageName"))
	tomlString := utilities.ReadFileContent(filepath.Join(packageFolder, "cvpm.toml"))
	solvers := entities.LoadSolversFromConfig(tomlString)
	runtime.RenderRunnerTpl(packageFolder, solvers)
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	solverName := c.Param("solverName")
	var imageName string
	var log entities.Log
	var err error
	imageName = packageInfo.Package.Vendor + "-" + packageInfo.Package.Name + "-" + solverName
	// Check if docker file exists
	if !utilities.IsExists(filepath.Join(packageFolder, "docker_"+solverName)) {
		runtime.RenderDockerfile(solverName, packageFolder)
	}
	log, err = dockerClient.Build(strings.ToLower(imageName), filepath.Join(packageFolder, "docker_"+solverName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, messageResponse{
			Code: 500,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"logid": log.ID,
		})
	}
}
