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

import (
	"log"

	"github.com/cvpm-contrib/database"
)

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

// EnvironmentVariable declares the envs for runner to RUN
type EnvironmentVariable struct {
	Key         string `db:"env-Key"`
	Value       string `db:"env-Var"`
	Vendor      string `db:"env-Vendor`
	PackageName string `db:"env-PackageName"`
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

func QueryVariables(Vendor string, PackageName string) []EnvironmentVariable {
	var envs []EnvironmentVariable
	sess := database.GetDBInstance()
	envCollection := sess.Collection("environment")
	res := envCollection.Find()
	err := res.All(&envs)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
	database.CloseDB(sess)
	return envs
}

func AddNewEnvToPackage(Vendor string, PackageName string, Key string, Value string) {
	envObject := EnvironmentVariable{
		Vendor:      Vendor,
		PackageName: PackageName,
		Key:         Key,
		Value:       Value,
	}
	sess := database.GetDBInstance()
	envCollection := sess.Collection("environment")
	err := envCollection.InsertReturning(&envObject)
	if err != nil {
		panic(err)
	}
	database.CloseDB(sess)
}
