package runtime

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/docker/docker/api/types"
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

var logger = utilities.NewLogger()

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
func (docker *DockerRuntime) Pull(imageName string) {
	reader, err := docker.client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		logger.Error("Cannot pull image " + imageName)
	}
	io.Copy(os.Stdout, reader)
}

// Build will build a new image from dockerfile
func (docker *DockerRuntime) Build(imageName string, dockerfile string) {
	logger.Info("Starting Build Image...")
	err := archiver.Archive([]string{path.Dir(dockerfile)}, "archieve.tar")
	dockerBuildContext, err := os.Open(filepath.Join(path.Dir(dockerfile), "archieve.tar"))
	defer dockerBuildContext.Close()
	buildResponse, err := docker.client.ImageBuild(context.Background(), dockerBuildContext, types.ImageBuildOptions{})
	if err != nil {
		logger.Error("Cannot build image " + imageName)
		logger.Error(err.Error())
	}
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		logger.Error("Cannot build image " + imageName)
		logger.Error(err.Error())
	}
	os.Remove(filepath.Join(path.Dir(dockerfile), "archieve.tar"))
}

// GenerateDockerFile returns a DockerFile String that could be used to build image.
func GenerateDockerFile(baseImageName string, setup string) {

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
