package workflow

import "github.com/autoai-org/aid/internal/runtime/docker"

// BuildDockerImage builds the docker image
func BuildDockerImage(vendorName string, packageName string, solverName string, gpu bool) string {
	logID := docker.BuildImage(vendorName, packageName, solverName, gpu)
	return logID
}

// BuildDockerImageWithPath builds the docker image in standalone mode
func BuildDockerImageWithPath(path string, solverName string, autoRemove bool) {
	// todo: now force images to be removed after building
	docker.BuildWithPath(path, solverName, autoRemove)
}
