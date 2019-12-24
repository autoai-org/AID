package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
)

func build () {
	// Read Dockerfile
	dockerClient := runtime.NewDockerRuntime()
	// Read Base Image Name
	dockerClient.Build("imagename", "/home/xzyao/autoai-org/aiflow/components/cmd/examples/face/Dockerfile")
}