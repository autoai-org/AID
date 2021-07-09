package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	entSolver "github.com/autoai-org/aid/ent/generated/solver"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/runtime/vcs"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type SolverInfoResponse struct {
	Solvername  string         `json:"solvername"`
	Reponame    string         `json:"reponame"`
	Vendorname  string         `json:"vendorname"`
	RemoteURL   string         `json:"remoteURL"`
	Description string         `json:"description"`
	GitCommits  vcs.GitCommits `json:"commits"`
}

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
	c.JSON(http.StatusOK, SolverInfoResponse{
		Solvername:  solver.Name,
		Reponame:    repo.Name,
		Vendorname:  repo.Vendor,
		RemoteURL:   repo.RemoteURL,
		Description: description,
		GitCommits:  gitcommits,
	})
}
