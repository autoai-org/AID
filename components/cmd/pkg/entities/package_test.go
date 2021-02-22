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

func TestPackage_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Vendor    string
		Status    string
		CreatedAt time.Time
		UpdatedAt time.Time
		RemoteURL string
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
			p := &Package{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Vendor:    tt.fields.Vendor,
				Status:    tt.fields.Status,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				RemoteURL: tt.fields.RemoteURL,
			}
			if got := p.TableName(); got != tt.want {
				t.Errorf("Package.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_PK(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Vendor    string
		Status    string
		CreatedAt time.Time
		UpdatedAt time.Time
		RemoteURL string
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
			p := &Package{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Vendor:    tt.fields.Vendor,
				Status:    tt.fields.Status,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				RemoteURL: tt.fields.RemoteURL,
			}
			if got := p.PK(); got != tt.want {
				t.Errorf("Package.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadPackageFromConfig(t *testing.T) {
	type args struct {
		tomlString string
	}
	tests := []struct {
		name string
		args args
		want PackageConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadPackageFromConfig(tt.args.tomlString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadPackageFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackage_Save(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		LocalPath string
		Vendor    string
		Status    string
		CreatedAt time.Time
		UpdatedAt time.Time
		RemoteURL string
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
			p := &Package{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				LocalPath: tt.fields.LocalPath,
				Vendor:    tt.fields.Vendor,
				Status:    tt.fields.Status,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				RemoteURL: tt.fields.RemoteURL,
			}
			if err := p.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Package.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchPackages(t *testing.T) {
	tests := []struct {
		name string
		want []Package
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchPackages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPackage(t *testing.T) {
	type args struct {
		vendorName  string
		packageName string
	}
	tests := []struct {
		name string
		args args
		want Package
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPackage(tt.args.vendorName, tt.args.packageName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadPretrainedsFromConfig(t *testing.T) {
	type args struct {
		tomlString string
	}
	tests := []struct {
		name string
		args args
		want Pretraineds
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadPretrainedsFromConfig(tt.args.tomlString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadPretrainedsFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPackageByImageID(t *testing.T) {
	type args struct {
		ImageID string
	}
	tests := []struct {
		name string
		args args
		want Package
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPackageByImageID(tt.args.ImageID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPackageByImageID() = %v, want %v", got, tt.want)
			}
		})
	}
}
