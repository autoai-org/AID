package main

import (
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func InstallHandler(c *cli.Context) {
	config := readConfig()
	localFolder := config.Local.LocalFolder
	remoteURL := c.Args().Get(0)
	if remoteURL == "cvpm:test" {
		color.Cyan("Installing... Please wait patiently")
		pip([]string{"install", "--index-url", "https://test.pypi.org/simple/", "cvpm", "--user"})
		return
	} else {
		color.Cyan("Installing to " + localFolder)
	}
	var repoFolder string
	var repo Repository
	// Download Codebase
	if strings.HasPrefix(remoteURL, "https://github.com") {
		repo = CloneFromGit(remoteURL, localFolder)
	}
	repoFolder = repo.LocalFolder
	// Install Dependencies
	color.Cyan("Installing Dependencies... please wait patiently")
	InstallDependencies(repoFolder)
	color.Blue("Generating Runners")
	GeneratingRunners(repoFolder)
	color.Cyan("Adding to Local Configuration")
	config.Repositories = addRepo(config.Repositories, repo)
	writeConfig(config)
}

func listRepos(c *cli.Context) {
	config := readConfig()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No.", "Name", "Vendor", "Path"})
	for i, repo := range config.Repositories {
		table.Append([]string{strconv.Itoa(i), repo.Name, repo.Vendor, repo.LocalFolder})
	}
	table.Render()
}

func DaemonHandler(c *cli.Context) {
	params := c.Args().Get(0)
	switch params {
	case "install":
		InstallService()
	case "uninstall":
		UninstallService()
	case "run":
		runServer(DaemonPort)
	default:
		color.Red("Unsupported command")
	}
}

func RepoHandler(c *cli.Context) {
	taskParams := c.Args().Get(0)
	switch taskParams {
	case "run":
		solverstring := c.Args().Get(1)
		runParams := strings.Split(solverstring, "/")
		color.Cyan("Running " + runParams[0] + "/" + runParams[1] + "/" + runParams[2])
		requestParams := map[string]string{
			"vendor": runParams[0],
			"name":   runParams[1],
			"solver": runParams[2],
		}
		ClientPost("repo", requestParams)
		// runRepo(runParams[0], runParams[1], runParams[2])
	default:
		color.Red("Command Not Supported!")
	}
}
