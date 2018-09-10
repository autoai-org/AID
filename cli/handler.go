package main

import (
	"strings"
	"log"
	"github.com/urfave/cli"
)

func InstallHandler (c *cli.Context) {
	config := readConfig()
	localFolder := config.Local.LocalFolder
	remoteURL := c.Args().Get(0)
	if (remoteURL == "cvpm:test") {
		log.Printf("Installing... Please wait patiently")
		pip([]string{"install", "--index-url", "https://test.pypi.org/simple/", "cvpm", "--user"})
		return
	} else {
		log.Printf("Installing to " + localFolder)
	}
	var repoFolder string
	// Download Codebase
	if strings.HasPrefix(remoteURL, "https://github.com") {
		repoFolder = CloneFromGit(remoteURL, localFolder)	
	}
	// Install Dependencies
	log.Printf("Installing Dependencies... Please wait patiently")
	InstallDependencies(repoFolder)
}
