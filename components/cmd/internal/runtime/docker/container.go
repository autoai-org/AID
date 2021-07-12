// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"context"
	"os"

	entContainer "github.com/autoai-org/aid/ent/generated/container"
	entImage "github.com/autoai-org/aid/ent/generated/image"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

type GPURequest struct {
	NeedGPU bool
}

// Create creates a docker container
func Create(imageUID string, hostPort string, gpu GPURequest) (container.ContainerCreateCreatedBody, error) {
	image, err := database.NewDefaultDB().Image.Query().Where(entImage.UID(imageUID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch image " + imageUID + ", Aborted")
		os.Exit(3)
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
	}
	var env []string
	if gpu.NeedGPU {
		utilities.Formatter.Info("GPU Enabled")
		hostConfig.Resources = container.Resources{
			DeviceRequests: []container.DeviceRequest{
				{
					Capabilities: [][]string{{"gpu"}},
				},
			},
		}
		env = append(env, "LD_LIBRARY_PATH=/usr/local/cuda/extras/CUPTI/lib64:/usr/local/cuda/lib64:/usr/local/nvidia/lib:/usr/local/nvidia/lib64")
		env = append(env, "NVIDIA_VISIBLE_DEVICES=all")
		env = append(env, "NVIDIA_DRIVER_CAPABILITIES=all")
		env = append(env, "NVIDIA_REQUIRE_CUDA=cuda>=11.2 brand=tesla,driver>=418,driver<419 brand=tesla,driver>=440,driver<441 driver>=450")
		hostConfig.Privileged = true
	} else {
		utilities.Formatter.Info("GPU Disabled")
	}
	resp, err := NewDockerRuntime().ContainerCreate(context.Background(), &container.Config{
		Env:   env,
		Image: image.UID,
		Tty:   true,
		ExposedPorts: nat.PortSet{
			"8080/tcp": struct{}{},
		},
	}, hostConfig, nil, nil, image.UID)
	if err != nil {
		utilities.ReportError(err, "Cannot create container from image "+image.UID)
		return resp, err
	}
	newContainerEntity, err := database.NewDefaultDB().Container.Create().SetUID(resp.ID[0:8]).SetPort(hostPort).SetImage(image).Save(context.Background())
	image.Edges.Container = newContainerEntity
	image.Update()
	utilities.ReportError(err, "Cannot save "+resp.ID)
	utilities.Formatter.Info("Successfully created container for " + image.Title)
	utilities.Formatter.Info("The reference for the created container is " + resp.ID[0:8])
	return resp, err
}

// Start will start a docker container
func Start(containerID string) error {
	containerEnt, err := database.NewDefaultDB().Container.Query().Where(entContainer.UID(containerID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch container: " + err.Error())
		os.Exit(3)
	}
	if containerEnt.Running {
		utilities.Formatter.Error("The requested container has already been started: " + containerEnt.UID)
		os.Exit(4)
	}
	if err := NewDockerRuntime().ContainerStart(context.Background(), containerEnt.UID, types.ContainerStartOptions{}); err != nil {
		utilities.ReportError(err, "Cannot start container "+containerID)
		return err
	}
	containerEnt.Update().SetRunning(true).Save(context.Background())
	return nil
}

// Stop will stop a docker container
func Stop(containerID string) error {
	containerEnt, err := database.NewDefaultDB().Container.Query().Where(entContainer.UID(containerID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch container: " + err.Error())
		os.Exit(3)
	}
	if !containerEnt.Running {
		utilities.Formatter.Error("The requested container is not running: " + containerEnt.UID)
		os.Exit(4)
	}
	if err := NewDockerRuntime().ContainerStop(context.Background(), containerEnt.UID, nil); err != nil {
		utilities.ReportError(err, "Cannot start container "+containerID)
		return err
	}
	_, err = containerEnt.Update().SetRunning(false).Save(context.Background())
	if err != nil {
		utilities.ReportError(err, "Cannot stop the container")
		return err
	}
	utilities.Formatter.Info("Successfully stopped " + containerID)
	return nil
}

// RemoveContainer will remove a docker container
func RemoveContainer(containerID string) error {
	containerEnt, err := database.NewDefaultDB().Container.Query().Where(entContainer.UID(containerID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch container: " + err.Error())
		os.Exit(3)
	}
	if containerEnt.Running {
		utilities.Formatter.Error("The requested container is running: " + containerEnt.UID + ". You must stop it first.")
		os.Exit(4)
	}
	if err := NewDockerRuntime().ContainerRemove(context.Background(), containerEnt.UID, types.ContainerRemoveOptions{}); err != nil {
		utilities.ReportError(err, "Cannot remove container "+containerID)
		return err
	}
	_, err = database.NewDefaultDB().Container.Delete().Where(entContainer.UID(containerID)).Exec(context.Background())
	if err != nil {
		utilities.ReportError(err, "Cannot remove container")
		return err
	}
	utilities.Formatter.Info("Successfully removed the container " + containerID)
	return err
}
