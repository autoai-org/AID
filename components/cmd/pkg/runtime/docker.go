// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"bufio"
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/mholt/archiver/v3"
	"golang.org/x/net/context"
	"io"
	"os"
	"path"
	"path/filepath"
)

var logger = utilities.NewDefaultLogger()

// DockerRuntime is the basic class for manage docker
type DockerRuntime struct {
	client *client.Client
}

// NewDockerRuntime returns a DockerRuntime Instance
func NewDockerRuntime() *DockerRuntime {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logger.Error("Cannot Create New Docker Runtime")
	}
	return &DockerRuntime{
		client: cli,
	}
}

// Pull will pull an exisiting package
func (docker *DockerRuntime) Pull(imageName string) error {
	reader, err := docker.client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		logger.Error("Cannot pull image " + imageName)
		logger.Error(err.Error())
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

// Create will create a container from an image
func (docker *DockerRuntime) Create(imageName string) container.ContainerCreateCreatedBody {
	resp, err := docker.client.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		logger.Error("Cannot create container from image " + imageName)
		logger.Error(err.Error())
	}
	return resp
}

// Start will start a container
func (docker *DockerRuntime) Start(containerID string) error {
	if err := docker.client.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
		logger.Error("Cannot start container " + containerID)
		logger.Error(err.Error())
		return err
	}
	return nil
}

// Build will build a new image from dockerfile
func (docker *DockerRuntime) Build(imageName string, dockerfile string) error {
	logger.Info("Starting Build Image: " + imageName + "...")
	err := archiver.Archive([]string{path.Dir(dockerfile)}, "archive.tar")
	dockerBuildContext, err := os.Open(filepath.Join(path.Dir(dockerfile), "archive.tar"))
	defer dockerBuildContext.Close()
	buildResponse, err := docker.client.ImageBuild(context.Background(), dockerBuildContext, types.ImageBuildOptions{
		Tags:       []string{"aiflow-" + imageName},
		Dockerfile: dockerfile,
	})
	if err != nil {
		logger.Error("Cannot build image " + imageName)
		logger.Error(err.Error())
	}
	reader := buildResponse.Body
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	var logPath = filepath.Join(utilities.GetBasePath(), "logs", "builds", imageName)
	entities.NewLogObject("build-"+imageName, logPath, "docker-build")
	buildLogger := utilities.NewLogger(logPath)
	for scanner.Scan() {
		buildLogger.Info(scanner.Text())
	}
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		logger.Error("Cannot build image " + imageName)
		logger.Error(err.Error())
	}
	os.Remove(filepath.Join(path.Dir(dockerfile), "archive.tar"))
	return err
}

// ListImages returns all images that have been installed on the host
func (docker *DockerRuntime) ListImages() []types.ImageSummary {
	images, err := docker.client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		logger.Error("Cannot List Images")
		logger.Error(err.Error())
	}
	return images
}

// ListContainers returns all containers that have been built.
func (docker *DockerRuntime) ListContainers() []types.Container {
	containers, err := docker.client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		logger.Error("Cannot List Images")
		logger.Error(err.Error())
	}
	return containers
}

// GenerateDockerFiles returns a DockerFile string that could be used to build image.
func GenerateDockerFiles(baseFilePath string) {
	configFile := filepath.Join(baseFilePath, "cvpm.toml")
	tomlString := utilities.ReadFileContent(configFile)
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	for _, solver := range packageInfo.Solvers {
		RenderDockerfile(solver.Name, baseFilePath)
	}
}

// FetchContainerLogs returns the logs in the container
func (docker *DockerRuntime) FetchContainerLogs(containerID string) {
	var logPath = filepath.Join(utilities.GetBasePath(), "logs", "container", containerID)
	entities.NewLogObject("container-"+containerID, logPath, "container-log")
	logger := utilities.NewLogger(logPath)
	reader, err := docker.client.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Timestamps: false,
		Follow:     true,
		Tail:       "40",
	})
	if err != nil {
		utilities.CheckError(err, "Cannot establish error log reader")
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		logger.Info(scanner.Text())
	}
}
