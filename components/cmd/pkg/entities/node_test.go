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

func TestNode_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Address   string
		Token     string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
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
			n := &Node{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Address:   tt.fields.Address,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
			}
			if got := n.TableName(); got != tt.want {
				t.Errorf("Node.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_PK(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Address   string
		Token     string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
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
			n := &Node{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Address:   tt.fields.Address,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
			}
			if got := n.PK(); got != tt.want {
				t.Errorf("Node.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Save(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Address   string
		Token     string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
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
			n := &Node{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Address:   tt.fields.Address,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
			}
			if err := n.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Node.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchNodes(t *testing.T) {
	tests := []struct {
		name string
		want []Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchNodes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNodeByName(t *testing.T) {
	type args struct {
		nodeName string
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNodeByName(tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNodeByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
