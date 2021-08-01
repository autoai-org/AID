// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"bufio"
	_ "embed"
	"os"
	"path/filepath"
	"strings"

	ent "github.com/autoai-org/aid/ent/generated"
	"github.com/autoai-org/aid/internal/configuration"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/flosch/pongo2"
)

//go:embed assets/dockerfile.tpl
var dockerTpl string

//go:embed assets/runner.tpl
var runnerTpl string

//go:embed assets/dockerfile_gpu.tpl
var dockerGPUTpl string

//getTpl returns the template string
func getTpl(filename string) string {
	if filename == "dockerfile" {
		return dockerTpl
	}
	if filename == "runner" {
		return runnerTpl
	}
	if filename == "dockerfile_gpu" {
		return dockerGPUTpl
	}
	utilities.Formatter.Error("Unspecified template filename " + filename)
	return ""
}

// GenerateDockerFiles returns a DockerFile string that could be used to build image.
func GenerateDockerFiles(baseFilePath string) {
	configFile := filepath.Join(baseFilePath, "aid.toml")
	tomlString, err := utilities.ReadFileContent(configFile)
	if err != nil {
		utilities.ReportError(err, "cannot open file "+configFile)
	}
	packageInfo := configuration.LoadPackageFromConfig(tomlString)
	for _, solver := range packageInfo.Solvers {
		RenderDockerfile(solver.Name, baseFilePath, true)
		RenderDockerfile(solver.Name, baseFilePath, false)
	}
}

// RenderDockerfile returns the final dockerfile
func RenderDockerfile(solvername string, targetFilePath string, gpu bool) {
	var templateName string
	if gpu {
		templateName = "dockerfile_gpu"
	} else {
		templateName = "dockerfile"
	}
	tpl, err := pongo2.FromString(getTpl(templateName))
	utilities.ReportError(err, "Cannot render dockerfile")
	var filename string
	if gpu {
		filename = filepath.Join(targetFilePath, "docker_gpu_"+solvername)
	} else {
		filename = filepath.Join(targetFilePath, "docker_"+solvername)
	}

	setupFilePath := filepath.Join(targetFilePath, "setup.sh")
	var setupCommands string = ""
	if utilities.IsExists(setupFilePath) {
		f, err := os.Open(setupFilePath)
		defer f.Close()
		if err != nil {
			utilities.ReportError(err, "Cannot open file "+setupFilePath)
			setupCommands = "echo An error occured in parsing setup file"
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if setupCommands == "" {
				setupCommands = scanner.Text()
			} else {
				setupCommands = setupCommands + " && " + scanner.Text()
			}
		}
	} else {
		setupCommands = "echo There is no command for extra installation"
	}
	prepipFilePath := filepath.Join(targetFilePath, "prepip.sh")
	var prepipCommands string = ""
	if utilities.IsExists(prepipFilePath) {
		f, err := os.Open(prepipFilePath)
		defer f.Close()
		if err != nil {
			utilities.ReportError(err, "Cannot open file "+prepipFilePath)
			prepipCommands = "echo An error occured in parsing setup file"
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if prepipCommands == "" {
				prepipCommands = scanner.Text()
			} else {
				prepipCommands = prepipCommands + " && " + scanner.Text()
			}
		}
	} else {
		prepipCommands = "echo There is no command for extra installation"
	}
	out, err := tpl.Execute(pongo2.Context{"Solvername": solvername, "Setup": setupCommands, "PrePIP": prepipCommands})
	utilities.WriteContentToFile(filename, out)
}

// RenderRunnerTpl returns the final runner file
func RenderRunnerTpl(tempFilePath string, mySolvers []ent.Solver) {
	tpl, err := pongo2.FromString(getTpl("runner"))
	utilities.ReportError(err, "Cannot read template file")
	for _, solver := range mySolvers {
		filename := "runner_" + solver.Name + ".py"
		fileFullPath := filepath.Join(tempFilePath, filename)
		tplContext := strings.Split(solver.Class, "/")
		out, err := tpl.Execute(pongo2.Context{"Package": tplContext[0], "Filename": tplContext[1], "Classname": tplContext[2]})
		utilities.ReportError(err, "Failed to generate running file.")
		utilities.WriteContentToFile(fileFullPath, out)
	}
}
