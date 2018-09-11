package main

import (
	"os"
	"log"
	"path/filepath"
	"fmt"
	"strings"
	"gopkg.in/src-d/go-git.v4"
	"github.com/fatih/color"
	"github.com/BurntSushi/toml"
)

type Repository struct {
	Name string
	LocalFolder string
}

type solver struct {
	Name string
	Class string
}

type solvers struct {
	Solvers []solver
}

func readRepos () {

}

func addRepo () {

}

func delRepo () {

}

func CloneFromGit (remoteURL string, targetFolder string) string {
	localFolderName := strings.Split(remoteURL, "/")
	vendorName := localFolderName[len(localFolderName)-2]
	repoName := localFolderName[len(localFolderName)-1]
	localFolder := filepath.Join(targetFolder, vendorName,repoName)
	color.Cyan("Cloning " + remoteURL + " into " +localFolder)
	_, err := git.PlainClone(localFolder, false, &git.CloneOptions{
		URL: remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}
	return localFolder
}

func InstallCVPMPyPI () {
	
}

func InstallDependencies (localFolder string) {
	pip([]string{"install", "-r", filepath.Join(localFolder, "requirements.txt"), "--user"})
}

func GeneratingRunners (localFolder string) {
  var mySolvers solvers
	cvpmFile := filepath.Join(localFolder, "cvpm.toml")
	if _, err := toml.DecodeFile(cvpmFile, &mySolvers); err!=nil {
		log.Fatal(err)
	}
	fmt.Println(len(mySolvers.Solvers))
	renderRunnerTpl(localFolder, mySolvers)
}
