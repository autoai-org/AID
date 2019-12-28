// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewDefaultLogger("./logs/system.log")

// Solver defines the struct of a solver, the minimal struct of a inference program
type Solver struct {
	Name  string
	Class string
}

// Solvers defines a list of solvers
type Solvers struct {
	Solvers []Solver
}

// LoadSolversFromConfig reads the config string and returns the objects
func LoadSolversFromConfig(tomlString string) Solvers {
	var solvers Solvers
	_, err := toml.Decode(tomlString, &solvers)
	utilities.CheckError(err, "Cannot Load Solvers")
	return solvers
}
