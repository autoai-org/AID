// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*  This file defines the handlers for different command.
	Login
	Install
	List
	Repo
and etc. */
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Run Repo Response Struct
type RunRepoResponse struct {
	Code string `json:"code"`
	Port string `json:"port"`
}

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
	if remoteURL == "cvpm:py" {
		color.Cyan("Installing... Please wait patiently")
		pip([]string{"install", "cvpm", "--user"})
		return
	} else if remoteURL == "webui" {
		InstallWebUi()
	} else {
		color.Cyan("Installing to " + localFolder)
	}
	// Download Codebase
	if strings.HasPrefix(remoteURL, "https://github.com") {
		InstallFromGit(remoteURL)
	}
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

// Handle Repo Related Command
func RepoHandler(c *cli.Context) {
	taskParams := c.Args().Get(0)
	switch taskParams {
	case "run":
		solverstring := c.Args().Get(1)
		runningPort := c.Args().Get(2)
		runParams := strings.Split(solverstring, "/")
		color.Cyan("Running " + runParams[0] + "/" + runParams[1] + "/" + runParams[2])
		requestParams := map[string]string{
			"vendor": runParams[0],
			"name":   runParams[1],
			"solver": runParams[2],
			"port":   runningPort,
		}
		resp := ClientPost("repo/running", requestParams)
		var respJson RunRepoResponse
		log.Println(resp.JSON(respJson))
		color.Red("No port is specified, solver will running on" + respJson.Port)
	case "ps":
		requestParams := map[string]string{}
		ClientGet("repos", requestParams)
	case "init":
		InitHandler(c)
	default:
		color.Red("Command Not Supported!")
	}
}

// Handle Config Related Command

// validate if python/pip/others exists or does not change
func validateIfProgramAllowed(rawInput string) error {
	input := strings.TrimSpace(rawInput)
	if input == "y" || input == "Y" || input == "Yes" || input == "" {
		return nil
	} else {
		if _, err := os.Stat(input); os.IsNotExist(err) {
			return errors.New(input + " not exists")
		} else {
			return nil
		}
	}
}

// trigger and parse input filepath
func InputAndParseConfigContent(label string, validate promptui.ValidateFunc) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	return result
}

func ConfigHandler(c *cli.Context) {
	prevConfig := readConfig()
	var nextConfig CvpmConfig
	nextConfig.Local.LocalFolder = prevConfig.Local.LocalFolder
	// Handle Python Location
	fmt.Println("Original Python Location: " + prevConfig.Local.Python)
	newPyLocation := InputAndParseConfigContent("Python(3)", validateIfProgramAllowed)
	newPyLocation = strings.TrimSpace(newPyLocation)
	if newPyLocation == "y" || newPyLocation == "Y" || newPyLocation == "Yes" || newPyLocation == "" {
		newPyLocation = prevConfig.Local.Python
	}
	nextConfig.Local.Python = newPyLocation
	// Handle Pypi Location
	fmt.Println("Original Pip Location: " + prevConfig.Local.Pip)
	newPipLocation := InputAndParseConfigContent("Pip(3)", validateIfProgramAllowed)
	newPipLocation = strings.TrimSpace(newPipLocation)
	if newPipLocation == "y" || newPipLocation == "Y" || newPipLocation == "Yes" || newPipLocation == "" {
		newPipLocation = prevConfig.Local.Pip
	}
	nextConfig.Local.Pip = newPipLocation
	writeConfig(nextConfig)
}

func InitHandler(c *cli.Context) {
	prompt := promptui.Prompt{
		Label: "Your Package Name",
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}
	InitNewRepo(result)
	// rename {package_name} to real package name
	pyFolderName := filepath.Join(result, "{package_name}")
	err = os.Rename(pyFolderName, filepath.Join(result, result))
	if err != nil {
		fmt.Println(err)
	}
}
