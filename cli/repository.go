// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	raven "github.com/getsentry/raven-go"
	git "gopkg.in/src-d/go-git.v4"
)

type Repository struct {
	Name        string
	LocalFolder string
	Vendor      string
	Port        string
	Origin      string
}

type RepositoryMetaInfo struct {
	Config     string
	Dependency string
	DiskSize   float64
	Readme     string
}

type solver struct {
	Name  string
	Class string
}

type solvers struct {
	Solvers []solver
}

type RepoSolver struct {
	Vendor     string
	Package    string
	SolverName string
	Port       string
}

func readRepos() []Repository {
	configs := readConfig()
	repos := configs.Repositories
	return repos
}

func readClientRepos(currentHomedir string) []Repository {
	configs := readClientConfig(currentHomedir)
	repos := configs.Repositories
	return repos
}

func addRepo(repos []Repository, repo Repository) []Repository {
	alreadyInstalled := false
	for _, existed_repo := range repos {
		if repo.Name == existed_repo.Name && repo.Vendor == existed_repo.Vendor {
			alreadyInstalled = true
		}
	}
	if alreadyInstalled {
		log.Fatal("Repo Already Exists! Remove it first.")
	}
	repos = append(repos, repo)
	return repos
}

func delRepo(repos []Repository, Vendor string, Name string) []Repository {
	for i, repo := range repos {
		if repo.Name == Name && repo.Vendor == Vendor {
			repos = append(repos[:i], repos[i+1:]...)
		}
	}
	return repos
}

func runRepo(Vendor string, Name string, Solver string, Port string) {
	repos := readRepos()
	existed := false
	for _, existed_repo := range repos {
		if existed_repo.Name == Name && existed_repo.Vendor == Vendor {
			files, _ := ioutil.ReadDir(existed_repo.LocalFolder)
			for _, file := range files {
				if file.Name() == "runner_"+Solver+".py" {
					existed = true
					RunningRepos = append(RunningRepos, Repository{Vendor, Name, Solver, Port, ""})
					RunningSolvers = append(RunningSolvers, RepoSolver{Vendor: Vendor, Package: Name, SolverName: Solver, Port: Port})
					runfileFullPath := filepath.Join(existed_repo.LocalFolder, file.Name())
					python([]string{runfileFullPath, Port})
				}
			}
		}
	}
	if !existed {
		log.Fatal("Solver Not Found! Expecting " + Solver)
	}
}

// Clone a repo from @params remoteURL to @params targetFolder by Git Protocol.
// Used for installing and initialize a repo
func CloneFromGit(remoteURL string, targetFolder string) {
	color.Cyan("Cloning " + remoteURL + " into " + targetFolder)
	_, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
	}
}

func InstallDependencies(localFolder string) {
	pip([]string{"install", "-r", filepath.Join(localFolder, "requirements.txt"), "--user"})
}

// Generating Runners for future use
func GeneratingRunners(localFolder string) {
	var mySolvers solvers
	cvpmFile := filepath.Join(localFolder, "cvpm.toml")
	if _, err := toml.DecodeFile(cvpmFile, &mySolvers); err != nil {
		log.Fatal(err)
	}
	renderRunnerTpl(localFolder, mySolvers)
}

// After Installation
func PostInstallation(repoFolder string) {
	// Create pretrained folder
	preTrainedFolder := filepath.Join(repoFolder, "pretrained")
	exist, err := isPathExists(preTrainedFolder)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		err = os.Mkdir(preTrainedFolder, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Return Repository Meta Info: Dependency, Config, Disk Size and Readme
func GetMetaInfo(Vendor string, Name string) RepositoryMetaInfo {
	repos := readRepos()
	repositoryMeta := RepositoryMetaInfo{}
	for _, existed_repo := range repos {
		if existed_repo.Name == Name && existed_repo.Vendor == Vendor {
			// Read config file etc
			readmeFilePath := filepath.Join(existed_repo.LocalFolder, "README.md")
			cvpmConfigFilePath := filepath.Join(existed_repo.LocalFolder, "cvpm.toml")
			requirementsFilePath := filepath.Join(existed_repo.LocalFolder, "requirements.txt")
			repositoryMeta.Config = readFileContent(cvpmConfigFilePath)
			repositoryMeta.Dependency = readFileContent(requirementsFilePath)
			repositoryMeta.Readme = readFileContent(readmeFilePath)
			packageSize := getDirSizeMB(existed_repo.LocalFolder)
			repositoryMeta.DiskSize = packageSize
		}
	}
	return repositoryMeta
}

// Install Repository from Git
func InstallFromGit(remoteURL string) {
	config := readConfig()
	var repo Repository
	// prepare local folder
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	localFolder := filepath.Join(config.Local.LocalFolder, vendorName, repoName)
	CloneFromGit(remoteURL, localFolder)
	repo = Repository{Name: repoName, Vendor: vendorName, LocalFolder: localFolder, Origin: remoteURL}

	repoFolder := repo.LocalFolder
	InstallDependencies(repoFolder)
	GeneratingRunners(repoFolder)
	config.Repositories = addRepo(config.Repositories, repo)
	writeConfig(config)
	PostInstallation(repoFolder)
}

// Init a new repoo by using bolierplate
func InitNewRepo(repoName string) {
	bolierplateURL := "https://github.com/cvmodel/bolierplate.git"
	CloneFromGit(bolierplateURL, repoName)
}
