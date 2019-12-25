package main

import (
	"github.com/alexeyco/simpletable"
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
	"strconv"
	"strings"
)

func build() {
	// Read Dockerfile
	dockerClient := runtime.NewDockerRuntime()
	// Read Base Image Name
	dockerClient.Build("imagename", "./Dockerfile")
}

func printImages() {
	dockerClient := runtime.NewDockerRuntime()
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Image ID"},
			{Align: simpletable.AlignCenter, Text: "Repo Tags"},
			{Align: simpletable.AlignCenter, Text: "Size"},
		},
	}
	images := dockerClient.ListImages()
	for _, image := range images {
		r := []*simpletable.Cell{
			{Text: image.ID},
			{Text: strings.Join(image.RepoTags, ",")},
			{Text: strconv.FormatInt(image.Size, 10)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}

func printContainers() {
	dockerClient := runtime.NewDockerRuntime()
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Container ID"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Image"},
		},
	}
	containers := dockerClient.ListContainers()
	for _, container := range containers {
		r := []*simpletable.Cell{
			{Text: container.ID},
			{Text: container.Status},
			{Text: container.Image},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}
