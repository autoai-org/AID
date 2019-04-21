// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"github.com/unarxiv/cvpm/pkg/entity"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

// GetRunnerTpl returns the template string
func GetRunnerTpl() string {
	var runnerTpl = "https://raw.githubusercontent.com/unarxiv/CVPM/master/templates/runner.tpl"
	resp, err := http.Get(runnerTpl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// RenderRunnerTpl returns the final runner file
func RenderRunnerTpl(localFolder string, mysolvers entity.Solvers) {
	tpl, err := pongo2.FromString(GetRunnerTpl())
	if err != nil {
		log.Fatal(err)
	}
	for _, solver := range mysolvers.Solvers {
		WriteRunner(tpl, localFolder, solver)
	}
}

// WriteRunner edits the runner file
func WriteRunner(tpl *pongo2.Template, localFolder string, Solver entity.Solver) {
	filename := "runner_" + Solver.Name + ".py"
	fileFullPath := filepath.Join(localFolder, filename)
	tplContext := strings.Split(Solver.Class, "/")
	out, err := tpl.Execute(pongo2.Context{"Package": tplContext[0], "Filename": tplContext[1], "Classname": tplContext[2]})
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fileFullPath, []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
