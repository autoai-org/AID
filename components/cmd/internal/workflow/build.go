package workflow

import "github.com/autoai-org/aid/internal/runtime/docker"

// BuildDockerImage builds the docker image
func BuildDockerImage(vendorName string, packageName string, solverName string) {
	docker.BuildImage(vendorName, packageName, solverName)
}

func BuildDockerImageWithPath(path string, solverName string, autoRemove bool) {
	// todo: now force images to be removed after building
	docker.BuildWithPath(path, solverName, autoRemove)
}
