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

func TestLog_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Filepath  string
		CreatedAt time.Time
		UpdatedAt time.Time
		Source    string
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
			l := &Log{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Filepath:  tt.fields.Filepath,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Source:    tt.fields.Source,
			}
			if got := l.TableName(); got != tt.want {
				t.Errorf("Log.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_PK(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Filepath  string
		CreatedAt time.Time
		UpdatedAt time.Time
		Source    string
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
			l := &Log{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Filepath:  tt.fields.Filepath,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Source:    tt.fields.Source,
			}
			if got := l.PK(); got != tt.want {
				t.Errorf("Log.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_Save(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Filepath  string
		CreatedAt time.Time
		UpdatedAt time.Time
		Source    string
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
			l := &Log{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Filepath:  tt.fields.Filepath,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Source:    tt.fields.Source,
			}
			if err := l.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Log.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLogObject(t *testing.T) {
	type args struct {
		title    string
		filepath string
		source   string
	}
	tests := []struct {
		name string
		args args
		want Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogObject(tt.args.title, tt.args.filepath, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchLogs(t *testing.T) {
	tests := []struct {
		name string
		want []Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchLogs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteLog(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteLog(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLog(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLog(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveInitLogObject(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveInitLogObject()
		})
	}
}
