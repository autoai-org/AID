// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"reflect"
	"testing"
)

func TestPrivateEnvironment_TableName(t *testing.T) {
	type fields struct {
		ID   string
		Name string
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
			p := &PrivateEnvironment{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := p.TableName(); got != tt.want {
				t.Errorf("PrivateEnvironment.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariable_TableName(t *testing.T) {
	type fields struct {
		ID          string
		Key         string
		Value       string
		Environment string
		PackageID   string
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
			ev := &EnvironmentVariable{
				ID:          tt.fields.ID,
				Key:         tt.fields.Key,
				Value:       tt.fields.Value,
				Environment: tt.fields.Environment,
				PackageID:   tt.fields.PackageID,
			}
			if got := ev.TableName(); got != tt.want {
				t.Errorf("EnvironmentVariable.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariable_PK(t *testing.T) {
	type fields struct {
		ID          string
		Key         string
		Value       string
		Environment string
		PackageID   string
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
			ev := &EnvironmentVariable{
				ID:          tt.fields.ID,
				Key:         tt.fields.Key,
				Value:       tt.fields.Value,
				Environment: tt.fields.Environment,
				PackageID:   tt.fields.PackageID,
			}
			if got := ev.PK(); got != tt.want {
				t.Errorf("EnvironmentVariable.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateEnvironment_PK(t *testing.T) {
	type fields struct {
		ID   string
		Name string
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
			p := &PrivateEnvironment{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := p.PK(); got != tt.want {
				t.Errorf("PrivateEnvironment.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironmentVariable_Save(t *testing.T) {
	type fields struct {
		ID          string
		Key         string
		Value       string
		Environment string
		PackageID   string
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
			ev := &EnvironmentVariable{
				ID:          tt.fields.ID,
				Key:         tt.fields.Key,
				Value:       tt.fields.Value,
				Environment: tt.fields.Environment,
				PackageID:   tt.fields.PackageID,
			}
			if err := ev.Save(); (err != nil) != tt.wantErr {
				t.Errorf("EnvironmentVariable.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrivateEnvironment_Save(t *testing.T) {
	type fields struct {
		ID   string
		Name string
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
			p := &PrivateEnvironment{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if err := p.Save(); (err != nil) != tt.wantErr {
				t.Errorf("PrivateEnvironment.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetEnvironmentVariablesbyPackageID(t *testing.T) {
	type args struct {
		packageID   string
		environment string
	}
	tests := []struct {
		name string
		args args
		want []EnvironmentVariable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvironmentVariablesbyPackageID(tt.args.packageID, tt.args.environment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvironmentVariablesbyPackageID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeEnvironmentVariables(t *testing.T) {
	type args struct {
		envs []EnvironmentVariable
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeEnvironmentVariables(tt.args.envs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeEnvironmentVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}
