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
}

type solver struct {
	Name  string
	Class string
}

type solvers struct {
	Solvers []solver
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
					RunningRepos = append(RunningRepos, existed_repo)
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

func GeneratingRunners(localFolder string) {
	var mySolvers solvers
	cvpmFile := filepath.Join(localFolder, "cvpm.toml")
	if _, err := toml.DecodeFile(cvpmFile, &mySolvers); err != nil {
		log.Fatal(err)
	}
	renderRunnerTpl(localFolder, mySolvers)
}
