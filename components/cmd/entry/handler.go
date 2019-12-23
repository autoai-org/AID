package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
)

func build () {
	// Read Dockerfile
	dockerfile := utilities.ReadFileContent("Dockerfile")
	dockerClient := runtime.NewDockerRuntime()
	// Read Base Image Name
	dockerClient.Build("imagename", dockerfile)
}