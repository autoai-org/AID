// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"reflect"
	"testing"
	"time"
)

func TestDataset_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Status    string
		Size      float64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Status:    tt.fields.Status,
				Size:      tt.fields.Size,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := d.TableName(); got != tt.want {
				t.Errorf("Dataset.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExperiment_TableName(t *testing.T) {
	type fields struct {
		ID       string
		Datapath string
		Vendor   string
		Package  string
		Solver   string
		LogID    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Experiment{
				ID:       tt.fields.ID,
				Datapath: tt.fields.Datapath,
				Vendor:   tt.fields.Vendor,
				Package:  tt.fields.Package,
				Solver:   tt.fields.Solver,
				LogID:    tt.fields.LogID,
			}
			if got := e.TableName(); got != tt.want {
				t.Errorf("Experiment.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataset_PK(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Status    string
		Size      float64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Status:    tt.fields.Status,
				Size:      tt.fields.Size,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := d.PK(); got != tt.want {
				t.Errorf("Dataset.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExperiment_PK(t *testing.T) {
	type fields struct {
		ID       string
		Datapath string
		Vendor   string
		Package  string
		Solver   string
		LogID    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Experiment{
				ID:       tt.fields.ID,
				Datapath: tt.fields.Datapath,
				Vendor:   tt.fields.Vendor,
				Package:  tt.fields.Package,
				Solver:   tt.fields.Solver,
				LogID:    tt.fields.LogID,
			}
			if got := e.PK(); got != tt.want {
				t.Errorf("Experiment.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataset_Save(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Status    string
		Size      float64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Status:    tt.fields.Status,
				Size:      tt.fields.Size,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := d.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Dataset.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExperiment_Save(t *testing.T) {
	type fields struct {
		ID       string
		Datapath string
		Vendor   string
		Package  string
		Solver   string
		LogID    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Experiment{
				ID:       tt.fields.ID,
				Datapath: tt.fields.Datapath,
				Vendor:   tt.fields.Vendor,
				Package:  tt.fields.Package,
				Solver:   tt.fields.Solver,
				LogID:    tt.fields.LogID,
			}
			if err := e.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Experiment.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchAllDatasets(t *testing.T) {
	tests := []struct {
		name string
		want []Dataset
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchAllDatasets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchAllDatasets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchAllExperiments(t *testing.T) {
	tests := []struct {
		name string
		want []Experiment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchAllExperiments(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchAllExperiments() = %v, want %v", got, tt.want)
			}
		})
	}
}
