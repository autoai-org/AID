// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"io"
	"reflect"
	"testing"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func Test_prepareEnvs(t *testing.T) {
	type args struct {
		imageID string
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
			if got := prepareEnvs(tt.args.imageID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareEnvs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareVolumes(t *testing.T) {
	type args struct {
		hostPath string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareVolumes(tt.args.hostPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareVolumes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDockerRuntime(t *testing.T) {
	tests := []struct {
		name string
		want *DockerRuntime
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDockerRuntime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDockerRuntime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerRuntime_Pull(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		imageName string
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
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if err := docker.Pull(tt.args.imageName); (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.Pull() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDockerRuntime_Create(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		imageID  string
		hostPort string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    container.ContainerCreateCreatedBody
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			got, err := docker.Create(tt.args.imageID, tt.args.hostPort)
			if (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerRuntime.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerRuntime_Start(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		containerID string
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
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if err := docker.Start(tt.args.containerID); (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getBuildCtx(t *testing.T) {
	type args struct {
		solverPath string
	}
	tests := []struct {
		name string
		args args
		want io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBuildCtx(tt.args.solverPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBuildCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_realBuild(t *testing.T) {
	type args struct {
		docker      *DockerRuntime
		dockerfile  string
		imageName   string
		buildLogger *logrus.Logger
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
			if err := realBuild(tt.args.docker, tt.args.dockerfile, tt.args.imageName, tt.args.buildLogger); (err != nil) != tt.wantErr {
				t.Errorf("realBuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDockerRuntime_Build(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		imageName  string
		dockerfile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Log
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			got, err := docker.Build(tt.args.imageName, tt.args.dockerfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerRuntime.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerRuntime_ListImages(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.ImageSummary
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if got := docker.ListImages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerRuntime.ListImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerRuntime_ListContainers(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if got := docker.ListContainers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerRuntime.ListContainers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateDockerFiles(t *testing.T) {
	type args struct {
		baseFilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateDockerFiles(tt.args.baseFilePath)
		})
	}
}

func TestDockerRuntime_FetchContainerLogs(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		containerID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entities.Log
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if got := docker.FetchContainerLogs(tt.args.containerID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DockerRuntime.FetchContainerLogs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerRuntime_ExportImage(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		imageName string
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
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if err := docker.ExportImage(tt.args.imageName); (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.ExportImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDockerRuntime_ImportImage(t *testing.T) {
	type fields struct {
		client *client.Client
	}
	type args struct {
		imageName string
		quiet     bool
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
			docker := &DockerRuntime{
				client: tt.fields.client,
			}
			if err := docker.ImportImage(tt.args.imageName, tt.args.quiet); (err != nil) != tt.wantErr {
				t.Errorf("DockerRuntime.ImportImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
