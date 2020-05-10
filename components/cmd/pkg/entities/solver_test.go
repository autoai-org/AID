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

func TestSolver_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Class     string
		Vendor    string
		Package   string
		Status    string
		ImageName string
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
			s := &Solver{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Class:     tt.fields.Class,
				Vendor:    tt.fields.Vendor,
				Package:   tt.fields.Package,
				Status:    tt.fields.Status,
				ImageName: tt.fields.ImageName,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := s.TableName(); got != tt.want {
				t.Errorf("Solver.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
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
			i := &Image{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := i.TableName(); got != tt.want {
				t.Errorf("Image.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_TableName(t *testing.T) {
	type fields struct {
		ID        string
		ImageID   string
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
			c := &Container{
				ID:        tt.fields.ID,
				ImageID:   tt.fields.ImageID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := c.TableName(); got != tt.want {
				t.Errorf("Container.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunningSolver_TableName(t *testing.T) {
	type fields struct {
		ID          string
		ContainerID string
		Status      string
		ImageName   string
		EntryPoint  string
		CreatedAt   time.Time
		UpdatedAt   time.Time
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
			rs := &RunningSolver{
				ID:          tt.fields.ID,
				ContainerID: tt.fields.ContainerID,
				Status:      tt.fields.Status,
				ImageName:   tt.fields.ImageName,
				EntryPoint:  tt.fields.EntryPoint,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if got := rs.TableName(); got != tt.want {
				t.Errorf("RunningSolver.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_PK(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Class     string
		Vendor    string
		Package   string
		Status    string
		ImageName string
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
			s := &Solver{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Class:     tt.fields.Class,
				Vendor:    tt.fields.Vendor,
				Package:   tt.fields.Package,
				Status:    tt.fields.Status,
				ImageName: tt.fields.ImageName,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := s.PK(); got != tt.want {
				t.Errorf("Solver.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage_PK(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
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
			i := &Image{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := i.PK(); got != tt.want {
				t.Errorf("Image.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_PK(t *testing.T) {
	type fields struct {
		ID        string
		ImageID   string
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
			c := &Container{
				ID:        tt.fields.ID,
				ImageID:   tt.fields.ImageID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := c.PK(); got != tt.want {
				t.Errorf("Container.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunningSolver_PK(t *testing.T) {
	type fields struct {
		ID          string
		ContainerID string
		Status      string
		ImageName   string
		EntryPoint  string
		CreatedAt   time.Time
		UpdatedAt   time.Time
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
			rs := &RunningSolver{
				ID:          tt.fields.ID,
				ContainerID: tt.fields.ContainerID,
				Status:      tt.fields.Status,
				ImageName:   tt.fields.ImageName,
				EntryPoint:  tt.fields.EntryPoint,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if got := rs.PK(); got != tt.want {
				t.Errorf("RunningSolver.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_Save(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Class     string
		Vendor    string
		Package   string
		Status    string
		ImageName string
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
			s := &Solver{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Class:     tt.fields.Class,
				Vendor:    tt.fields.Vendor,
				Package:   tt.fields.Package,
				Status:    tt.fields.Status,
				ImageName: tt.fields.ImageName,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Solver.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestImage_Save(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
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
			i := &Image{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := i.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Image.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunningSolver_Save(t *testing.T) {
	type fields struct {
		ID          string
		ContainerID string
		Status      string
		ImageName   string
		EntryPoint  string
		CreatedAt   time.Time
		UpdatedAt   time.Time
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
			rs := &RunningSolver{
				ID:          tt.fields.ID,
				ContainerID: tt.fields.ContainerID,
				Status:      tt.fields.Status,
				ImageName:   tt.fields.ImageName,
				EntryPoint:  tt.fields.EntryPoint,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := rs.Save(); (err != nil) != tt.wantErr {
				t.Errorf("RunningSolver.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContainer_Save(t *testing.T) {
	type fields struct {
		ID        string
		ImageID   string
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
			c := &Container{
				ID:        tt.fields.ID,
				ImageID:   tt.fields.ImageID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := c.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Container.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetImage(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetImage(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImagebyName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetImagebyName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImagebyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContainer(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContainer(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextImageNumber(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNextImageNumber(); got != tt.want {
				t.Errorf("GetNextImageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchSolvers(t *testing.T) {
	tests := []struct {
		name string
		want []Solver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchSolvers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchSolvers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchImages(t *testing.T) {
	tests := []struct {
		name string
		want []Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchImages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchContainers(t *testing.T) {
	tests := []struct {
		name string
		want []Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchContainers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchContainers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadSolversFromConfig(t *testing.T) {
	type args struct {
		tomlString string
	}
	tests := []struct {
		name string
		args args
		want Solvers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadSolversFromConfig(tt.args.tomlString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadSolversFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRunningSolver(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name string
		args args
		want RunningSolver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRunningSolver(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRunningSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllRunningSolvers(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name string
		args args
		want RunningSolver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllRunningSolvers(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRunningSolvers() = %v, want %v", got, tt.want)
			}
		})
	}
}
