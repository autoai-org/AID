// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"reflect"
	"testing"
)

func TestNewGitClient(t *testing.T) {
	tests := []struct {
		name string
		want GitClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitClient_Clone(t *testing.T) {
	type args struct {
		remoteURL    string
		targetFolder string
	}
	tests := []struct {
		name      string
		gitclient *GitClient
		args      args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gitclient := &GitClient{}
			gitclient.Clone(tt.args.remoteURL, tt.args.targetFolder)
		})
	}
}
