// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// Package runtime defines the running environment for each solver
// The runner defines the entrance in pre-1.0x version,
// Now the entrance can also be defined by Dockerfile and thus is not necessary.
package runtime

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/flosch/pongo2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

// getRunnerTpl fetches the template for running solver
// This is for ensuring backward compatibility
func getTpl(tplName string) string {
	var runnerTpl = "https://raw.githubusercontent.com/autoai-org/CVPM/master/templates/" + tplName + ".tpl"
	resp, err := http.Get(runnerTpl)
	utilities.CheckError(err, "Cannot Fetch Template..., Please check your internet connection!")
	body, err := ioutil.ReadAll(resp.Body)
	utilities.CheckError(err, "Cannot Read from Response..., Please check your internet connection!")
	return string(body)
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
	setupFileContent := utilities.ReadFileContent(filepath.Join(targetFilePath, "setup.sh"))
	var setupCommands string
	if setupFileContent != "Read "+filepath.Join(targetFilePath, "setup.sh")+" Failed!" {
		commands := strings.Split(strings.Replace(setupFileContent, "\r\n", "\n", -1), "\n")
		setupCommands = strings.Join(commands, "&& \\ \n")
	} else {
		setupCommands = "echo There is no command for extra installation"
	}
	out, err := tpl.Execute(pongo2.Context{"Solvername": solvername, "Setup": setupCommands})
	utilities.WriteContentToFile(filename, out)
}
