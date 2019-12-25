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
func getRunnerTpl() string {
	var runnerTpl = "https://raw.githubusercontent.com/autoai-org/CVPM/master/templates/runner.tpl"
	resp, err := http.Get(runnerTpl)
	if err != nil {
		logger.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err)
	}
	return string(body)
}

// RenderRunnerTpl returns the final runner file
func RenderRunnerTpl(tempFilePath string, mySolvers entities.Solvers) {
	tpl, err := pongo2.FromString(getRunnerTpl())
	if err != nil {
		logger.Fatal(err)
	}
	for _, solver := range mySolvers.Solvers {
		filename := "runner_" + solver.Name + ".py"
		fileFullPath := filepath.Join(tempFilePath, filename)
		tplContext := strings.Split(solver.Class, "/")
		out, err := tpl.Execute(pongo2.Context{"Package": tplContext[0], "Filename": tplContext[1], "Classname": tplContext[2]})
		if err != nil {
			logger.Error("Failed to generate running file.")
			logger.Error(err.Error())
		}
		utilities.WriteContentToFile(fileFullPath, out)
	}
}
