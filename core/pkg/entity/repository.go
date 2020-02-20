// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package entity

// Repository defines the basic structure of a repository/package
type Repository struct {
	Name        string
	LocalFolder string
	Vendor      string
	Port        string
	Origin      string
}

// RepositoryMetaInfo defines the meta-information struct of a repository
type RepositoryMetaInfo struct {
	Config     string
	Dependency string
	DiskSize   float64
	Readme     string
}

// Solver defines the struct of a solver, the minimal struct of a inference program
type Solver struct {
	Name  string
	Class string
}

// Solvers defines a list of solvers
type Solvers struct {
	Solvers []Solver
}

// RepoSolver defines a struct for using the port
type RepoSolver struct {
	Vendor     string
	Package    string
	SolverName string
	Port       string
}
