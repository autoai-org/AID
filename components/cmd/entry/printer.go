// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/autoai-org/aid/components/cmd/pkg/runtime"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/flosch/pongo2"
)

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

func printConfig() {
	config := utilities.NewDefaultConfig()
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Config Name"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignCenter, Text: "Value"},
		},
	}
	r := []*simpletable.Cell{
		{Text: "RemoteReport"},
		{Text: "Allows report bugs/stats to our server to allow user experience improvements"},
		{Text: strconv.FormatBool(config.RemoteReport)},
	}
	table.Body.Cells = append(table.Body.Cells, r)
	table.Println()
}

func renderFile(templateFile string, context map[string]interface{}) error {
	fileContent, err := utilities.ReadFileContent(templateFile)
	utilities.CheckError(err, "Cannot read file "+templateFile)
	tpl, err := pongo2.FromString(fileContent)
	utilities.CheckError(err, "Failed to get template.")
	out, err := tpl.Execute(context)
	utilities.CheckError(err, "Failed to generate file.")
	utilities.WriteContentToFile(templateFile, out)
	return err
}
