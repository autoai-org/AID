// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
)

// getPackages returns all packages
// GET /packages -> returns packages
func getPackages(c *gin.Context) {
	packages := entities.FetchPackages()
	c.JSON(http.StatusOK, packages)
}

// getImages returns all images that has been built
// GET /images -> returns all containers
func getImages(c *gin.Context) {
	images := entities.FetchImages()
	c.JSON(http.StatusOK, images)
}

// getSolvers returns all solvers
// GET /solvers -> returns all solvers
func getSolvers(c *gin.Context) {
	solvers := entities.FetchSolvers()
	c.JSON(http.StatusOK, solvers)
}

// getContainers returns all containers
// GET /containers -> returns all containers
func getContainers(c *gin.Context) {
	containers := entities.FetchContainers()
	c.JSON(http.StatusOK, containers)
}

// getSolverDockerfile returns current dockerfile of the solver
// GET /packages/:vendorName/:packageName/:solverName/dockerfile
func getDockerfileContent(c *gin.Context) {
	packageFolder := filepath.Join(utilities.GetBasePath(), "models", c.Param("vendorName"), c.Param("packageName"))
	dockerFilename := "docker_" + c.Param("solverName")
	dockerFilePath := filepath.Join(packageFolder, dockerFilename)
	fileContent, err := utilities.ReadFileContent(dockerFilePath)
	if err != nil {
		c.JSON(http.StatusOK, messageResponse{
			Code: 404,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, messageResponse{
			Code: 200,
			Msg:  fileContent,
		})
	}
}

// modifySolverDockerfile modifies the content of the dockerfile
// POST /packages/:vendorName/:packageName/:solverName/dockerfile
func modifySolverDockerfile(c *gin.Context) {
	var request modifySolverDockerfileRequest
	c.BindJSON(&request)
	packageFolder := filepath.Join(utilities.GetBasePath(), "models", c.Param("vendorName"), c.Param("packageName"))
	dockerFilename := "docker_" + c.Param("solverName")
	dockerFilePath := filepath.Join(packageFolder, dockerFilename)
	err := utilities.WriteContentToFile(dockerFilePath, request.Content)
	if err != nil {
		c.JSON(http.StatusOK, messageResponse{
			Code: 404,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, messageResponse{
			Code: 200,
			Msg:  "Successfully Changed",
		})
	}
}

// getMetaInfo returns all meta information about the package
// GET /:vendorName/:packageName/meta -> {readme, aid.toml, pretrained.toml}
func getMetaInfo(c *gin.Context) {
	var solvers entities.Solvers
	var pretraineds entities.Pretraineds
	var readmeContent string
	packageFolder := filepath.Join(utilities.GetBasePath(), "models", c.Param("vendorName"), c.Param("packageName"))
	aidtomlString, err := utilities.ReadFileContent(filepath.Join(packageFolder, "aid.toml"))
	if err == nil {
		solvers = entities.LoadSolversFromConfig(aidtomlString)
	}
	pretrainedtomlString, err := utilities.ReadFileContent(filepath.Join(packageFolder, "pretrained.toml"))
	if err == nil {
		pretraineds = entities.LoadPretrainedsFromConfig(pretrainedtomlString)
	}
	readmeContent, err = utilities.ReadFileContent(filepath.Join(packageFolder, "README.md"))
	c.JSON(http.StatusOK, metaResponse{
		Solvers: solvers, Pretraineds: pretraineds, Readme: readmeContent,
	})
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

// createContainers will create a container by using the given image
// PUT /containers/:imageId/containers
func createSolverContainer() {

}

// buildPackages build a new image
// PUT /:vendorName/:packageName/:solverName/images
func buildSolverImage(c *gin.Context) {
	dockerClient := runtime.NewDockerRuntime()
	packageFolder := filepath.Join(utilities.GetBasePath(), "models", c.Param("vendorName"), c.Param("packageName"))
	tomlString, err := utilities.ReadFileContent(filepath.Join(packageFolder, "aid.toml"))
	if err != nil {
		utilities.CheckError(err, "Cannot open file "+filepath.Join(packageFolder, "aid.toml"))
	}
	solvers := entities.LoadSolversFromConfig(tomlString)
	runtime.RenderRunnerTpl(packageFolder, solvers)
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	solverName := c.Param("solverName")
	var imageName string
	var log entities.Log
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
