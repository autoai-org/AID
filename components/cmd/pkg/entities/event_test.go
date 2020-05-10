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

func TestIBE(t *testing.T) {

}

func TestEvent_TableName(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Data      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
		From      string
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
			e := &Event{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Data:      tt.fields.Data,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				From:      tt.fields.From,
			}
			if got := e.TableName(); got != tt.want {
				t.Errorf("Event.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_PK(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Data      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
		From      string
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
			e := &Event{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Data:      tt.fields.Data,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				From:      tt.fields.From,
			}
			if got := e.PK(); got != tt.want {
				t.Errorf("Event.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_Save(t *testing.T) {
	type fields struct {
		ID        string
		Title     string
		Data      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Status    string
		From      string
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
			e := &Event{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Data:      tt.fields.Data,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				Status:    tt.fields.Status,
				From:      tt.fields.From,
			}
			if err := e.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Event.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConsume(t *testing.T) {
	type args struct {
		ie IEvent
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Consume(tt.args.ie)
		})
	}
}

func TestImageBuiltEvent_msg(t *testing.T) {
	type fields struct {
		ImageName string
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
			ibe := ImageBuiltEvent{
				ImageName: tt.fields.ImageName,
			}
			if got := ibe.msg(); got != tt.want {
				t.Errorf("ImageBuiltEvent.msg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageBuiltEvent_toEvent(t *testing.T) {
	type fields struct {
		ImageName string
	}
	tests := []struct {
		name   string
		fields fields
		want   Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ibe := ImageBuiltEvent{
				ImageName: tt.fields.ImageName,
			}
			if got := ibe.toEvent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageBuiltEvent.toEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
