package docker

import (
	"io"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

// BuildLog defines the stream of logs when building docker images
type BuildLog struct {
	Stream string `json:"stream"`
}

// Client is the basic class for manage docker
var Client *client.Client

// NewDockerRuntime returns a DockerRuntime Instance
func NewDockerRuntime() *client.Client {
	if Client != nil {
		return Client
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	utilities.ReportError(err, "Cannot Create New Docker Runtime")
	return cli
}

func getBuildCtx(dockerPath string) io.Reader {
	ctx, _ := archive.TarWithOptions(dockerPath, &archive.TarOptions{ExcludePatterns: []string{".git"}})
	return ctx
}
