// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/autoai-org/aid/components/cmd/pkg/storage"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

// Dataset defines basic structure for Dataset
type Dataset struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	LocalPath string `db:"localpath"`
	Status    string `db:"status"`
	Size      float64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Experiment defines the basic structure for Experiment
type Experiment struct {
	ID       string `db:"id"`
	Datapath string `db:"datapath"`
	Vendor   string `db:"vendor"`
	Package  string `db:"package"`
	Solver   string `db:"solver"`
	LogID    string `db:"logid"`
}

// TableName defines the tablename for dataset
func (d *Dataset) TableName() string {
	return "dataset"
}

// TableName defines the tablename for Experiment
func (e *Experiment) TableName() string {
	return "experiment"
}

// PK returns the primary key of Dataset
func (d *Dataset) PK() string {
	return "id"
}

// PK returns the primary key of experiment
func (e *Experiment) PK() string {
	return "id"
}

// Save stores dataset into database
func (d *Dataset) Save() error {
	d.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(d)
}

// Save stores experiment into database
func (e *Experiment) Save() error {
	e.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(e)
}

// FetchAllDatasets returns all datasets
func FetchAllDatasets() []Dataset {
	datasetPointers := make([]*Dataset, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&datasetPointers)
	datasets := make([]Dataset, len(datasetPointers))
	for i := range datasetPointers {
		datasets[i] = *datasetPointers[i]
	}
	return datasets
}

// FetchAllExperiments returns all experiments
func FetchAllExperiments() []Experiment {
	experimentPointers := make([]*Experiment, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&experimentPointers)
	experiments := make([]Experiment, len(experimentPointers))
	for i := range experimentPointers {
		experiments[i] = *experimentPointers[i]
	}
	return experiments
}
