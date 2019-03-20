// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

func _getRunnerTpl() string {
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

func renderRunnerTpl(localFolder string, mysolvers solvers) {
	tpl, err := pongo2.FromString(_getRunnerTpl())
	if err != nil {
		log.Fatal(err)
	}
	for _, solver := range mysolvers.Solvers {
		writeRunner(tpl, localFolder, solver)
	}
}

func writeRunner(tpl *pongo2.Template, localFolder string, Solver solver) {
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
