package workflow

import "github.com/autoai-org/aid/internal/runtime/docker"

// BuildDockerImage builds the docker image
func BuildDockerImage(vendorName string, packageName string, solverName string) {
	docker.BuildImage(vendorName, packageName, solverName)
}
