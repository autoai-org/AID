// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// Package runtime defines the running environment for each solver
// The runner defines the entrance in pre-1.0x version,
// Now the entrance can also be defined by Dockerfile and thus is not necessary.
package runtime

import (
	"bufio"
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/flosch/pongo2"
	"os"
	"path/filepath"
	"strings"
)

// getRunnerTpl fetches the template for running solver
// This is for ensuring backward compatibility
func getTpl(tplName string) string {
	var runnerTpl = "https://raw.githubusercontent.com/autoai-org/CVPM/master/templates/" + tplName + ".tpl"
	return utilities.GetRemoteFile(runnerTpl)
}

// RenderRunnerTpl returns the final runner file
func RenderRunnerTpl(tempFilePath string, mySolvers entities.Solvers) {
	tpl, err := pongo2.FromString(getTpl("runner"))
	utilities.CheckError(err, "Cannot read template file")
	for _, solver := range mySolvers.Solvers {
		filename := "runner_" + solver.Name + ".py"
		fileFullPath := filepath.Join(tempFilePath, filename)
		tplContext := strings.Split(solver.Class, "/")
		out, err := tpl.Execute(pongo2.Context{"Package": tplContext[0], "Filename": tplContext[1], "Classname": tplContext[2]})
		utilities.CheckError(err, "Failed to generate running file.")
		utilities.WriteContentToFile(fileFullPath, out)
	}
}

// RenderDockerfile returns the final dockerfile
func RenderDockerfile(solvername string, targetFilePath string) {
	tpl, err := pongo2.FromString(getTpl("dockerfile"))
	utilities.CheckError(err, "Cannot render dockerfile")
	filename := filepath.Join(targetFilePath, "docker_"+solvername)
	setupFilePath := filepath.Join(targetFilePath, "setup.sh")
	var setupCommands string = ""
	if utilities.IsExists(setupFilePath) {
		f, err := os.Open(setupFilePath)
		defer f.Close()
		if err != nil {
			utilities.CheckError(err, "Cannot open file "+setupFilePath)
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
	out, err := tpl.Execute(pongo2.Context{"Solvername": solvername, "Setup": setupCommands})
	utilities.WriteContentToFile(filename, out)
}
