// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
* This file handles virtual environment for engine to use.
* It is designed to support different virtual environment frameworks, for now,
* it support:
* Venv, Supported
* Docker, Planned but not supported yet.
* Besides, the file also handles environment variables.
 */

package runtime

import (
	"github.com/unarxiv/cvpm/pkg/utility"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

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
	if utility.IsStringInSlice(virtualenv.Name, SUPPORTEDVIRTUALENV) {
		return true
	}
	return false
}

func (virtualenv VirtualEnv) initiate() {
	Python([]string{"-m", "venv", virtualenv.RepositoryFolder}, "")
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

func parseEnvs(envs []EnvironmentVariable) string {
	var envsString string
	for _, env := range envs {
		envString := env.Key + "=" + env.Value
		envsString += envString
		envsString += " "
	}
	return strings.TrimSpace(envsString)
}

// QueryEnvString returns the string for runnning
func QueryEnvString(Vendor string, PackageName string) string {
	envs := QueryVariables(Vendor, PackageName)
	return parseEnvs(envs)
}

// Gin Handlers for Envs

// QueryRepoEnvironments is a web handler returning all env variables
// for this repo and package
func QueryRepoEnvironments(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	c.JSON(http.StatusOK, QueryVariables(vendor, name))
}

// AddRepoEnvironments receives the env vars and add it to the database
func AddRepoEnvironments(c *gin.Context) {
	vendor := c.Param("vendor")
	name := c.Param("name")
	var addEnvvarRequest EnvironmentVariable
	c.BindJSON(&addEnvvarRequest)
	AddNewEnvToPackage(vendor, name, addEnvvarRequest.Key, addEnvvarRequest.Value)
	c.JSON(http.StatusOK, QueryVariables(vendor, name))
}
