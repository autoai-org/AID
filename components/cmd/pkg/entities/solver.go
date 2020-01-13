// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/BurntSushi/toml"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

// Solver defines the struct of a solver, the minimal struct of a inference program
type Solver struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	Class   string `db:"solverpath"`
	Vendor  string `db:"vendor"`
	Package string `db:"package"`
	Status  string `db:"status"`
}

// Solvers defines a list of solvers
type Solvers struct {
	Solvers []Solver
}

// TableName defines the tablename in database
func (s *Solver) TableName() string {
	return "solver"
}

// PK defines the primary key of Package
func (s *Solver) PK() string {
	return "id"
}

// Save stores package into database
func (s *Solver) Save() error {
	s.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(s)
}

// FetchSolvers returns all solvers in a single call
func FetchSolvers() []Solver {
	solversPointers := make([]*Solver, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&solversPointers)
	solvers := make([]Solver, len(solversPointers))
	for i := range solversPointers {
		solvers[i] = *solversPointers[i]
	}
	return solvers
}

// LoadSolversFromConfig reads the config string and returns the objects
func LoadSolversFromConfig(tomlString string) Solvers {
	var solvers Solvers
	_, err := toml.Decode(tomlString, &solvers)
	utilities.CheckError(err, "Cannot Load Solvers")
	return solvers
}
