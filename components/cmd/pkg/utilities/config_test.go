// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"reflect"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	tests := []struct {
		name string
		want *SystemConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveConfig(t *testing.T) {
	type args struct {
		config SystemConfig
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveConfig(tt.args.config)
		})
	}
}

func TestReadConfig(t *testing.T) {
	tests := []struct {
		name string
		want *SystemConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
