// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

// Dataset defines basic strcuture for Dataset
type Dataset struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	LocalPath string `db:"localpath"`
	Status    string `db:"status"`
	Size      float64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// TableName defines the tablename for dataset
func (d *Dataset) TableName() string {
	return "dataset"
}

// PK returns the primary key of Dataset
func (d *Dataset) PK() string {
	return "id"
}

// Save stores dataset into database
func (d *Dataset) Save() error {
	d.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(d)
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
