// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package entity

type Repository struct {
	Name        string
	LocalFolder string
	Vendor      string
	Port        string
	Origin      string
}

type RepositoryMetaInfo struct {
	Config     string
	Dependency string
	DiskSize   float64
	Readme     string
}

type Solver struct {
	Name  string
	Class string
}

type Solvers struct {
	Solvers []Solver
}

type RepoSolver struct {
	Vendor     string
	Package    string
	SolverName string
	Port       string
}