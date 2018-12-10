package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/getsentry/raven-go"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Repository struct {
	Name        string
	LocalFolder string
	Vendor      string
	Port        string
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
					RunningRepos = append(RunningRepos, Repository{Vendor, Name, Solver, Port})
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

func CloneFromGit(remoteURL string, targetFolder string) Repository {
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	localFolder := filepath.Join(targetFolder, vendorName, repoName)
	color.Cyan("Cloning " + remoteURL + " into " + localFolder)
	_, err := git.PlainClone(localFolder, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
	}
	repo := Repository{Name: repoName, Vendor: vendorName, LocalFolder: localFolder}
	return repo
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
			byte_config, err := ioutil.ReadFile(filepath.Join(existed_repo.LocalFolder, "cvpm.toml"))
			if err != nil {
				repositoryMeta.Config = "Read cvpm.toml failed"
			} else {
				repositoryMeta.Config = string(byte_config)
			}
			byte_dependency, err := ioutil.ReadFile(filepath.Join(existed_repo.LocalFolder, "requirements.txt"))
			if err != nil {
				repositoryMeta.Dependency = "Read requirements.txt failed"
			} else {
				repositoryMeta.Dependency = string(byte_dependency)
			}
			byte_readme, err := ioutil.ReadFile(filepath.Join(existed_repo.LocalFolder, "README.md"))
			if err != nil {
				repositoryMeta.Readme = "Read Readme.md failed"
			} else {
				repositoryMeta.Readme = string(byte_readme)
			}
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
	repo = CloneFromGit(remoteURL, config.Local.LocalFolder)
	repoFolder := repo.LocalFolder
	InstallDependencies(repoFolder)
	GeneratingRunners(repoFolder)
	config.Repositories = addRepo(config.Repositories, repo)
	writeConfig(config)
	PostInstallation(repoFolder)
}
