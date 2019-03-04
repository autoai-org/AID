// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
* This file handles virtual environment for engine to use.
* It is designed to support different virtual environment frameworks, for now,
* it support:
* Venv, Supported
* Docker, Planned but not supported yet.
 */
package main

// Constants
// SUPPORTEDVIRTUALENV declares all supported virtual environment method
var SUPPORTEDVIRTUALENV = []string{"venv"}

// VirtualEnv defines how a Virtual Environment works.
type VirtualEnv struct {
	Name               string // Name can be 'venv'
	RelativePythonPath string
	RelativePipPath    string
	IsEnabled          bool
	RepositoryFolder   string
}

func (virtualenv VirtualEnv) validate() bool {
	if isStringInSlice(virtualenv.Name, SUPPORTEDVIRTUALENV) {
		return true
	}
	return false
}

func (virtualenv VirtualEnv) initiate() {
	python([]string{"-m", "venv", virtualenv.RepositoryFolder})
}

func (virtualenv VirtualEnv) triggerEnable() {

}
