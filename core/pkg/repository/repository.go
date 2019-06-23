// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package repository

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/unarxiv/cvpm/pkg/entity"
	"github.com/unarxiv/cvpm/pkg/runtime"
	"github.com/unarxiv/cvpm/pkg/utility"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	raven "github.com/getsentry/raven-go"
	git "gopkg.in/src-d/go-git.v4"
)

func readRepos() []entity.Repository {
	configs := config.Read()
	repos := configs.Repositories
	return repos
}

func readClientRepos(currentHomedir string) []entity.Repository {
	configs := config.ReadClient(currentHomedir)
	repos := configs.Repositories
	return repos
}

func addRepo(repos []entity.Repository, repo entity.Repository) []entity.Repository {
	alreadyInstalled := false
	for _, existedRepo := range repos {
		if repo.Name == existedRepo.Name && repo.Vendor == existedRepo.Vendor {
			alreadyInstalled = true
		}
	}
	if alreadyInstalled {
		log.Fatal("Repo Already Exists! Remove it first.")
	}
	repos = append(repos, repo)
	return repos
}

func delRepo(repos []entity.Repository, Vendor string, Name string) []entity.Repository {
	for i, repo := range repos {
		if repo.Name == Name && repo.Vendor == Vendor {
			repos = append(repos[:i], repos[i+1:]...)
		}
	}
	return repos
}

// Run starts a solver
func Run(Vendor string, Name string, Solver string, Port string) {
	repos := readRepos()
	existed := false
	for _, existedRepo := range repos {
		if existedRepo.Name == Name && existedRepo.Vendor == Vendor {
			files, _ := ioutil.ReadDir(existedRepo.LocalFolder)
			for _, file := range files {
				if file.Name() == "runner_"+Solver+".py" {
					existed = true
					runtime.RunningRepos = append(runtime.RunningRepos, entity.Repository{Vendor, Name, Solver, Port, ""})
					runtime.RunningSolvers = append(runtime.RunningSolvers, entity.RepoSolver{Vendor: Vendor, Package: Name, SolverName: Solver, Port: Port})
					runfileFullPath := filepath.Join(existedRepo.LocalFolder, file.Name())
					// TODO: add environment vars
					envString := runtime.QueryEnvString(Vendor, Name)
					runtime.Python([]string{runfileFullPath, Port}, envString)
				}
			}
		}
	}
	if !existed {
		log.Fatal("Solver Not Found! Expecting " + Solver)
	}
}

// CloneFromGit clones a repo from @params remoteURL to @params targetFolder by Git Protocol.
// Used for installing and initializing a repo
func CloneFromGit(remoteURL string, targetFolder string) {
	color.Cyan("Cloning " + remoteURL + " into " + targetFolder)
	_, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
}

// InstallDependencies installs all the dependencies declared in the requirements.txt file
func InstallDependencies(localFolder string) {
	runtime.Pip([]string{"install", "-r", filepath.Join(localFolder, "requirements.txt"), "--user"})
}

// GeneratingRunners generates runners for future use, it uses template
func GeneratingRunners(localFolder string) {
	var mySolvers entity.Solvers
	cvpmFile := filepath.Join(localFolder, "cvpm.toml")
	if _, err := toml.DecodeFile(cvpmFile, &mySolvers); err != nil {
		log.Fatal(err)
	}
	runtime.RenderRunnerTpl(localFolder, mySolvers)
}

// PostInstallation handles folder creation and other things after Installation
func PostInstallation(repoFolder string) {
	// Create pretrained folder
	preTrainedFolder := filepath.Join(repoFolder, "pretrained")
	exist, err := utility.IsPathExists(preTrainedFolder)
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

// GetMetaInfo returns Repository Meta Info: Dependency, Config, Disk Size and Readme
func GetMetaInfo(Vendor string, Name string) entity.RepositoryMetaInfo {
	repos := readRepos()
	repositoryMeta := entity.RepositoryMetaInfo{}
	for _, existedRepo := range repos {
		if existedRepo.Name == Name && existedRepo.Vendor == Vendor {
			// Read config file etc
			readmeFilePath := filepath.Join(existedRepo.LocalFolder, "README.md")
			cvpmConfigFilePath := filepath.Join(existedRepo.LocalFolder, "cvpm.toml")
			requirementsFilePath := filepath.Join(existedRepo.LocalFolder, "requirements.txt")
			repositoryMeta.Config = utility.ReadFileContent(cvpmConfigFilePath)
			repositoryMeta.Dependency = utility.ReadFileContent(requirementsFilePath)
			repositoryMeta.Readme = utility.ReadFileContent(readmeFilePath)
			packageSize := utility.GetDirSizeMB(existedRepo.LocalFolder)
			repositoryMeta.DiskSize = packageSize
		}
	}
	return repositoryMeta
}

// InstallFromGit Install Repository from Git
func InstallFromGit(remoteURL string) {
	localConfig := config.Read()
	var repo entity.Repository
	// prepare local folder
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	localFolder := filepath.Join(localConfig.Local.LocalFolder, vendorName, repoName)
	CloneFromGit(remoteURL, localFolder)
	repo = entity.Repository{Name: repoName, Vendor: vendorName, LocalFolder: localFolder, Origin: remoteURL}

	repoFolder := repo.LocalFolder
	InstallDependencies(repoFolder)
	GeneratingRunners(repoFolder)
	localConfig.Repositories = addRepo(localConfig.Repositories, repo)
	config.Write(localConfig)
	PostInstallation(repoFolder)
}

// InitNewRepo inits a new repoo by using bolierplate
func InitNewRepo(repoName string) {
	bolierplateURL := "https://github.com/cvmodel/bolierplate.git"
	CloneFromGit(bolierplateURL, repoName)
}

// GetPretrained returns the pretrained file list
func GetPretrained(vendorName string, packageName string) []os.FileInfo {
	localConfig := config.Read()
	localFolder := filepath.Join(localConfig.Local.LocalFolder, vendorName, packageName)
	localPretrainedFolder := filepath.Join(localFolder, "pretrained")
	files, _ := ioutil.ReadDir(localPretrainedFolder)
	return files
}

// VerifyLocalRepository verifies if the local repository is well set
// check if cvpm.toml exists [err]
// check if the classes it refers to exists [err]
// check if pretrained.toml exists [warn]
// check if the files that pretrained.toml refers exists[err]
func VerifyLocalRepository(path string) bool {
	cvpmConfigFilePath := filepath.Join(path, "cvpm.toml")
	pretrainedFilePath := filepath.Join(path, "pretrained.toml")
	cvpmConfigFileExist, err := utility.IsPathExists(cvpmConfigFilePath)
	pretrainedFileExist, err := utility.IsPathExists(pretrainedFilePath)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
	if !cvpmConfigFileExist {
		log.Print("[Error] Cannot find cvpm.toml under " + path)
	}
	if !pretrainedFileExist {
		log.Print(("[Warn] Cannot find pretrained.toml under " + path))
	}
	if cvpmConfigFileExist && pretrainedFileExist {
		return true
	}
	return false
}
