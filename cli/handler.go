package main

import (
	"strings"
	"github.com/urfave/cli"
	"github.com/fatih/color"
)

func InstallHandler (c *cli.Context) {
	config := readConfig()
	localFolder := config.Local.LocalFolder
	remoteURL := c.Args().Get(0)
	if (remoteURL == "cvpm:test") {
		color.Cyan("Installing... Please wait patiently")
		pip([]string{"install", "--index-url", "https://test.pypi.org/simple/", "cvpm", "--user"})
		return
	} else {
		color.Cyan("Installing to " + localFolder)
	}
	var repoFolder string
	// Download Codebase
	if strings.HasPrefix(remoteURL, "https://github.com") {
		repoFolder = CloneFromGit(remoteURL, localFolder)	
	}
	// Install Dependencies
	color.Cyan("Installing Dependencies... please wait patiently")
	InstallDependencies(repoFolder)
	color.Blue("Generating Runners")
	GeneratingRunners(repoFolder)
}
