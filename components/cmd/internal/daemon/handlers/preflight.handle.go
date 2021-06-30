// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package handlers

import (
	"context"
	"net/http"

	"github.com/autoai-org/aid/ent/generated/container"
	"github.com/autoai-org/aid/ent/generated/image"
	"github.com/autoai-org/aid/ent/generated/repository"
	"github.com/autoai-org/aid/ent/generated/solver"
	"github.com/autoai-org/aid/internal/database"
	"github.com/gin-gonic/gin"
)

type PreflightRequests struct {
	Vendor      string `json:"vendor"`
	PackageName string `json:"package"`
	Solver      string `json:"solver"`
}

type PreflightResponse struct {
	Port string `json:"port"`
}

func PreflightHandler(c *gin.Context) {
	var preflightRequests PreflightRequests
	c.BindJSON(&preflightRequests)
	foundedSolver, err := database.NewDefaultDB().Solver.Query().Where(
		solver.HasRepositoryWith(repository.And(
			repository.Vendor(preflightRequests.Vendor),
			repository.Name(preflightRequests.PackageName),
		)),
	).First(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
	foundedImage, err := foundedSolver.Image(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
	foundedContainer, err := database.NewDefaultDB().Container.Query().Where(
		container.HasImageWith(image.ID(foundedImage.ID)),
	).First(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}
	if !foundedContainer.Running {
		c.JSON(http.StatusOK, ErrorResponse{Error: "The requested solver is not running"})
	}
	c.JSON(http.StatusOK, PreflightResponse{Port: foundedContainer.Port})
}
