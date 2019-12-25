package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewLogger()

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
	if _, err := toml.Decode(tomlString, &solvers); err != nil {
		logger.Fatal("Cannot Load Solvers frmom ")
	}
	return solvers
}
