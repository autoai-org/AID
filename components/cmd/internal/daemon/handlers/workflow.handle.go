// This file describes the workflow handler.
// Licensed under MIT License (see LICENSE).
package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	entRepository "github.com/autoai-org/aid/ent/generated/repository"
	"github.com/autoai-org/aid/internal/configuration"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/autoai-org/aid/internal/workflow"
	"github.com/gin-gonic/gin"
)

// InstallPackageRequest is the request struct for installing a package.
type InstallPackageRequest struct {
	RemoteURL string `json:"remoteURL"`
}

// MutateSolverStatusRequest is the request struct for mutating the status of a solver.
type MutateSolverStatusRequest struct {
	Operation   string `json:"operation"`
	SolverID    string `json:"solverId"`
	ImageID     string `json:"imageID"`
	ContainerID string `json:"containerID"`
	Hostport    string `json:"hostport"`
	VendorName  string `json:"vendorName"`
	PackageName string `json:"packageName"`
	SolverName  string `json:"solverName"`
}

// InstallPackageHandler is the handler for installing a package.
func InstallPackageHandler(c *gin.Context) {
	var request InstallPackageRequest
	c.BindJSON(&request)
	err := workflow.PullPackageSource(request.RemoteURL)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

// GetPackageConfiguration returns parsed the aid.toml content, such that the front end could utilized it.
func GetPackageConfiguration(c *gin.Context) {
	vendor := c.Param("vendor")
	packageName := c.Param("package")
	repository, err := database.NewDefaultDB().Repository.Query().Where(
		entRepository.Vendor(vendor), entRepository.Name(packageName)).First(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	}
	configurationFilePath := filepath.Join(repository.Localpath, "aid.toml")
	configurationToml, err := utilities.ReadFileContent(configurationFilePath)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	}
	packageConfig := configuration.LoadPackageFromConfig(configurationToml)
	c.JSON(http.StatusOK, packageConfig)
}

// MutateSolverStatusHandler is the handler for mutating the status of a solver.
func MutateSolverStatusHandler(c *gin.Context) {
	var request MutateSolverStatusRequest
	c.BindJSON(&request)
	switch request.Operation {
	case "build":
		logID := workflow.BuildDockerImage(request.VendorName, request.PackageName, request.SolverName, false, false)
		if logID != "" {
			c.JSON(200, gin.H{
				"logID": logID,
			})
		} else {
			c.JSON(500, gin.H{
				"error": "build failed",
			})
		}
	case "create":
		// if the port is not specified, we let the system choose one
		if request.Hostport == "" {

		}
		workflow.CreateContainer(request.ImageID, request.Hostport, false)
	case "stop":
		workflow.StopContainer(request.ContainerID)
	}
}
