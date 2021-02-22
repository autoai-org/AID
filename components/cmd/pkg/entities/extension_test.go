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

func TestWebhook_TableName(t *testing.T) {
	type fields struct {
		ID        string
		RemoteURL string
		Status    string
		Event     string
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
			w := &Webhook{
				ID:        tt.fields.ID,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				Event:     tt.fields.Event,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := w.TableName(); got != tt.want {
				t.Errorf("Webhook.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhook_PK(t *testing.T) {
	type fields struct {
		ID        string
		RemoteURL string
		Status    string
		Event     string
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
			w := &Webhook{
				ID:        tt.fields.ID,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				Event:     tt.fields.Event,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := w.PK(); got != tt.want {
				t.Errorf("Webhook.PK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhook_Save(t *testing.T) {
	type fields struct {
		ID        string
		RemoteURL string
		Status    string
		Event     string
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
			w := &Webhook{
				ID:        tt.fields.ID,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				Event:     tt.fields.Event,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := w.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Webhook.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhook_Trigger(t *testing.T) {
	type fields struct {
		ID        string
		RemoteURL string
		Status    string
		Event     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Webhook{
				ID:        tt.fields.ID,
				RemoteURL: tt.fields.RemoteURL,
				Status:    tt.fields.Status,
				Event:     tt.fields.Event,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			w.Trigger(tt.args.data)
		})
	}
}

func TestFetchWebhooks(t *testing.T) {
	tests := []struct {
		name string
		want []Webhook
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchWebhooks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchWebhooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
