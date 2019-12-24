package runtime

import (
	"io/ioutil"
	"fmt"
	"os"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"io"
	"github.com/jhoonb/archivex"
	"path/filepath"
	"path"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

var logger = utilities.NewLogger()

// DockerRuntime is the basic class for manage docker
type DockerRuntime struct {
	client *client.Client
}

// NewDockerRuntime returns a DockerRuntime Instance
func NewDockerRuntime () (*DockerRuntime) {
	cli, err := client.NewClientWithOpts(client.FromEnv,client.WithAPIVersionNegotiation())
	if err != nil {
        logger.Error("Cannot Create New Docker Runtime")
    }
	return &DockerRuntime{
		client: cli,
	}
}

func (docker *DockerRuntime) pull (imageName string) {
	reader, err := docker.client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		logger.Error("Cannot pull image " + imageName)
	}
	io.Copy(os.Stdout, reader)
}

// Build will build a new image from dockerfile
func (docker *DockerRuntime) Build (imageName string, dockerfile string) {
	tar := new(archivex.TarFile)
	tar.Create(filepath.Join(path.Dir(dockerfile), "archieve.tar"))
	tar.AddAll(path.Dir(dockerfile), false)
	tar.Close()
	dockerBuildContext, err := os.Open(filepath.Join(path.Dir(dockerfile), "archieve.tar"))
	defer dockerBuildContext.Close()
	buildResponse, err := docker.client.ImageBuild(context.Background(), dockerBuildContext, types.ImageBuildOptions{})
	if err != nil {
		logger.Error(err.Error())
		logger.Error("Cannot pull image " + imageName)
	}
	fmt.Printf("********* %s **********", buildResponse.OSType)
    response, err := ioutil.ReadAll(buildResponse.Body)
    fmt.Println(string(response))
}

// GenerateDockerFile returns a DockerFile String that could be used to build image.
func GenerateDockerFile (baseImageName string, setup string) {

}