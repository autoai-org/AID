// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package storage

import (
	"reflect"
	"testing"

	"github.com/ilibs/gosql/v2"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetDefaultDB(t *testing.T) {

}

func TestNewDB(t *testing.T) {
	type args struct {
		driver string
		uri    string
	}
	tests := []struct {
		name string
		args args
		want *Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(tt.args.driver, tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_CreateTables(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			db.CreateTables()
		})
	}
}

func TestDatabase_Connect(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
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
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Database.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Insert(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.Insert(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Database.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Fetch(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.Fetch(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Database.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_FetchOne(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.FetchOne(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Database.FetchOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Delete(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.Delete(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Database.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Update(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			if err := db.Update(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Database.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Count(t *testing.T) {
	type fields struct {
		configs     map[string]*gosql.Config
		isConnected bool
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int64
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				configs:     tt.fields.configs,
				isConnected: tt.fields.isConnected,
			}
			gotCount, err := db.Count(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("Database.Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
