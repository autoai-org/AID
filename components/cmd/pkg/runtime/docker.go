package runtime

import (
	"os"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"io"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

var logger = utilities.NewLogger()

// DockerRuntime is the basic class for manage docker
type DockerRuntime struct {
	ctx context.Context
	client *client.Client
}

// NewDockerRuntime returns a DockerRuntime Instance
func NewDockerRuntime () (*DockerRuntime) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
        logger.Error("Cannot Create New Docker Runtime")
    }
	return &DockerRuntime{
		ctx: ctx,
		client: cli,
	}
}

func (docker *DockerRuntime) pull (imageName string) {
	reader, err := docker.client.ImagePull(docker.ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		logger.Error("Cannot pull image " + imageName)
	}
	io.Copy(os.Stdout, reader)
}

// Build will build a new image from dockerfile
func (docker *DockerRuntime) Build (imageName string, dockerfile string) {
	_, err := docker.client.ImageBuild(docker.ctx, nil, types.ImageBuildOptions{Dockerfile: dockerfile})
	if err != nil {
		logger.Error(err.Error())
		logger.Error("Cannot pull image " + imageName)
	}
}

// GenerateDockerFile returns a DockerFile String that could be used to build image.
func GenerateDockerFile (baseImageName string, setup string) {

}