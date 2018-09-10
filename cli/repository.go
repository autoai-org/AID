package main

import (
	"os"
	"log"
	"path/filepath"
	"strings"
	"gopkg.in/src-d/go-git.v4"
)

type Repository struct {
	Name string
}

func CloneFromGit (remoteURL string, targetFolder string) string {
	localFolderName := strings.Split(remoteURL, "/")
	repoName := localFolderName[len(localFolderName)-1]
	localFolder := filepath.Join(targetFolder, repoName)
	log.Printf("Cloning " + remoteURL + " into " + localFolder)
	_, err := git.PlainClone(localFolder, false, &git.CloneOptions{
		URL: remoteURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
	return localFolder
}

func InstallCVPMPyPI () {
	
}

func InstallDependencies (localFolder string) {
	log.Printf("Installing Dependencies from " + localFolder)
	pip([]string{"install", "-r", filepath.Join(localFolder, "requirements.txt"), "--user"})
}
