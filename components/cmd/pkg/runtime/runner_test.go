// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"testing"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
)

func Test_getTpl(t *testing.T) {
	type args struct {
		tplName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTpl(tt.args.tplName); got != tt.want {
				t.Errorf("getTpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderRunnerTpl(t *testing.T) {
	type args struct {
		tempFilePath string
		mySolvers    entities.Solvers
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderRunnerTpl(tt.args.tempFilePath, tt.args.mySolvers)
		})
	}
}

func TestRenderDockerfile(t *testing.T) {
	type args struct {
		solvername     string
		targetFilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderDockerfile(tt.args.solvername, tt.args.targetFilePath)
		})
	}
}
