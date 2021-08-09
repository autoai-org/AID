// This file describes the API endpoints for the daemon (solver part).
// Licensed under MIT License (see LICENSE).
package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/autoai-org/aid/ent/generated"
	entSolver "github.com/autoai-org/aid/ent/generated/solver"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/runtime/docker"
	"github.com/autoai-org/aid/internal/runtime/vcs"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// SolverHandler is the handler for the solver information.
type SolverInfoResponse struct {
	Solvername  string                 `json:"solvername"`
	Reponame    string                 `json:"reponame"`
	Vendorname  string                 `json:"vendorname"`
	RemoteURL   string                 `json:"remoteURL"`
	Description string                 `json:"description"`
	Pretrained  string                 `json:"pretrained"`
	Containers  []*generated.Container `json:"containers"`
	GitCommits  vcs.GitCommits         `json:"commits"`
}

// CreateContainerRequest is  the struct for the request to create a container.
type CreateContainerRequest struct {
	ImageUID string `json:"imageUID"`
	HostPort string `json:"hostPort"`
	GPU      bool   `json:"gpu"`
}

// SolverInformationHandler is the handler for the solver information fetching.
func SolverInformationHandler(c *gin.Context) {
	solverID := c.Param("solverID")
	solver, err := database.NewDefaultDB().Solver.Query().Where(
		entSolver.UID(solverID)).First(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	}
	repo, err := solver.QueryRepository().First(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
	}
	description, err := utilities.ReadFileContent(filepath.Join(repo.Localpath, "README.md"))
	if err != nil {
		description = err.Error()
	}
	pretrainedFileContent, err := utilities.ReadFileContent(filepath.Join(repo.Localpath, "pretrained.toml"))
	if err != nil {
		pretrainedFileContent = ""
	}
	r, err := git.PlainOpen(repo.Localpath)
	if err != nil {
		description = err.Error()
	}
	ref, err := r.Head()
	if err != nil {
		description = err.Error()
	}
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		description = err.Error()
	}
	var gitcommits vcs.GitCommits
	_ = cIter.ForEach(func(c *object.Commit) error {
		gitcommits = append(gitcommits, vcs.GitCommit{
			Author:  c.Author.Name,
			Message: c.Message,
			When:    c.Author.When,
		})
		return nil
	})
	image, err := solver.QueryImage().First(context.Background())
	if err != nil {
		description = err.Error()
	}
	containers, err := image.QueryContainer().All(context.Background())
	if err != nil {
		description = err.Error()
	}
	c.JSON(http.StatusOK, SolverInfoResponse{
		Solvername:  solver.Name,
		Containers:  containers,
		Reponame:    repo.Name,
		Vendorname:  repo.Vendor,
		Pretrained:  pretrainedFileContent,
		RemoteURL:   repo.RemoteURL,
		Description: description,
		GitCommits:  gitcommits,
	})
}

// CreateContainerHandler is the handler for the container creation.
func CreateContainerHandler(c *gin.Context) {
	var req CreateContainerRequest
	c.BindJSON(req)
	docker.Create(req.ImageUID, req.HostPort, docker.GPURequest{NeedGPU: req.GPU})
}
