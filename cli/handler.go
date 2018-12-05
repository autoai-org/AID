/*
	This file defines the handlers for different command.
		Login
		Install
		List
		Repo
	and etc.
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/getsentry/raven-go"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

// Handle User Login
func LoginHandler(c *cli.Context) User {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Printf("Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := strings.TrimSpace(string(bytePassword))
	u := User{username, password, ""}
	currentUser := u.login()
	return currentUser
}

// Handle installation
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

// Handle List
func listRepos(c *cli.Context) {
	config := readConfig()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No.", "Name", "Vendor", "Path"})
	for i, repo := range config.Repositories {
		table.Append([]string{strconv.Itoa(i), repo.Name, repo.Vendor, repo.LocalFolder})
	}
	table.Render()
}

// Handle Daemon Related
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
		runningPort := c.Args().Get(2)
		if (runningPort == "") {
			runningPort = findNextOpenPort(8080)
			color.Red("No Running Port specified! Server will listen on: " + runningPort)
		}
		runParams := strings.Split(solverstring, "/")
		color.Cyan("Running " + runParams[0] + "/" + runParams[1] + "/" + runParams[2])
		requestParams := map[string]string{
			"vendor": runParams[0],
			"name":   runParams[1],
			"solver": runParams[2],
			"port":   runningPort,
		}
		ClientPost("repo", requestParams)
	case "ps":
		requestParams := map[string]string{}
		ClientGet("repos", requestParams)
	default:
		color.Red("Command Not Supported!")
	}
}

func ConfigHandler(c *cli.Context) {
	homepath, _ := homedir.Dir()
	configFilePath := filepath.Join(homepath, "cvpm", "config.toml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// config file not exists, create it
		file, err := os.Create(configFilePath)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			color.Red("An error occurred!")
		}
		defer file.Close()
		// config file not exists, write default to it
		writeConfig(getDefaultConfig())
	}
	prevConfig := readConfig()
	var nextConfig cvpmConfig
	nextConfig.Local.LocalFolder = prevConfig.Local.LocalFolder
	// Handle Python Location
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Python Location[" + prevConfig.Local.Python + "]")
	newPyLocation, _ := reader.ReadString('\n')
	newPyLocation = strings.TrimSpace(newPyLocation)
	if newPyLocation == "y" || newPyLocation == "Y" || newPyLocation == "Yes" || newPyLocation == "" {
		newPyLocation = prevConfig.Local.Python
	} else {
		if _, err := os.Stat(newPyLocation); os.IsNotExist(err) {
			log.Fatal("Python executable file not found: No such file")
		}
	}
	nextConfig.Local.Python = newPyLocation
	// Handle Pypi Location
	fmt.Printf("Pip Location[" + prevConfig.Local.Pip + "]")
	newPipLocation, _ := reader.ReadString('\n')
	newPipLocation = strings.TrimSpace(newPipLocation)
	if newPipLocation == "y" || newPipLocation == "Y" || newPipLocation == "Yes" || newPipLocation == "" {
		newPipLocation = prevConfig.Local.Pip
	} else {
		if _, err := os.Stat(newPipLocation); os.IsNotExist(err) {
			log.Fatal("Pip executable file not found: No such file")
		}
	}
	nextConfig.Local.Pip = newPipLocation
	writeConfig(nextConfig)
}

func InitHandler(c *cli.Context) {

}
