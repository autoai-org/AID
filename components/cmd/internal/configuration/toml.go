package configuration

import (
	ent "github.com/autoai-org/aid/ent/generated"
)

// Pretrained defines the basic structure of pretrained file,
// it do not need to be stored in database
// and therefore has no `db` bindings.
type Pretrained struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

// Pretraineds is the collection/list of pretrained files
// this definition is used for toml parser
type Pretraineds struct {
	Models []Pretrained `toml:"models"`
}

// PackageConfig is the toml interface as in aid.toml
type PackageConfig struct {
	// Solvers is the declaration of all solvers
	Solvers []ent.Solver   `toml:"solvers"`
	Package ent.Repository `toml:"package"`
}
