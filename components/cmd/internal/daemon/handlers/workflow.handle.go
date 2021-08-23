// This file describes the workflow handler.
// Licensed under MIT License (see LICENSE).
package handlers

import (
	"github.com/autoai-org/aid/internal/workflow"
	"github.com/gin-gonic/gin"
)

type InstallPackageRequest struct {
	RemoteURL string `json:"remoteURL"`
}

type MutateSolverStatus struct {
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
	c.BindJSON(request)
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

// MutateSolverStatusHandler is the handler for mutating the status of a solver.
func MutateSolverStatusHandler(c *gin.Context) {
	var request MutateSolverStatus
	c.BindJSON(request)
	switch request.Operation {
	case "build":
		filepath := workflow.BuildDockerImage(request.VendorName, request.PackageName, request.SolverName, false)
		if filepath != "" {
			c.JSON(200, gin.H{
				"filepath": filepath,
			})
		} else {
			c.JSON(500, gin.H{
				"error": "build failed",
			})
		}
	case "create":
		workflow.CreateContainer(request.ImageID, request.Hostport, false)
	case "stop":
		workflow.StopContainer(request.ContainerID)
	}

}
