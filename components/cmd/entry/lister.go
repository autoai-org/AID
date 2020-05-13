// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/runtime"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

func baseList(headers simpletable.Header, items [][]string) {
	table := simpletable.New()
	table.Header = &headers
	for _, item := range items {
		r := []*simpletable.Cell{}
		for _, each := range item {
			rowcol := simpletable.Cell{
				Align: simpletable.AlignCenter, Text: each,
			}
			r = append(r, &rowcol)
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}

func listableToListItem(reqToList string) (simpletable.Header, [][]string) {
	// Create dockerClient beforehand as several items need dockerclient
	// There will be only one instance and won't cost much time to create one.
	dockerClient := runtime.NewDockerRuntime()
	var headers simpletable.Header
	var items [][]string
	if reqToList == "images" {
		headers = simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Image ID (Short)"},
				{Align: simpletable.AlignCenter, Text: "Repo Names (Vendor/Package/Solver)"},
				{Align: simpletable.AlignCenter, Text: "Size (GB)"},
			},
		}
		for _, image := range dockerClient.ListImages() {
			repoTags := strings.Join(image.RepoTags, ",")
			if strings.HasPrefix(repoTags, "aid-") {
				repoIdentifiers := strings.Split(repoTags, "-")
				repoString := repoIdentifiers[2] + "/" + repoIdentifiers[3] + "/" + strings.TrimRight(repoIdentifiers[4], ":latest") + ":" + repoIdentifiers[1]
				row := []string{
					image.ID[8:16],
					repoString,
					strconv.FormatFloat(float64(image.Size)/float64((1000*1000*1000)), 'f', 2, 64),
				}
				items = append(items, row)
			}
		}
	} else if reqToList == "containers" {
		headers = simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Container ID"},
				{Align: simpletable.AlignCenter, Text: "Image"},
				{Align: simpletable.AlignCenter, Text: "Status"},
			},
		}
		for _, container := range dockerClient.ListContainers() {
			if strings.HasPrefix(container.Image, "aid-") {
				repoIdentifiers := strings.Split(container.Image, "-")
				repoString := repoIdentifiers[2] + "/" + repoIdentifiers[3] + "/" + strings.TrimRight(repoIdentifiers[4], ":latest") + ":" + repoIdentifiers[1]
				row := []string{
					container.ID[8:16],
					repoString,
					container.Status,
				}
				items = append(items, row)
			}
		}
	} else if reqToList == "config" {
		config := utilities.NewDefaultConfig()
		headers = simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Config Name"},
				{Align: simpletable.AlignCenter, Text: "Description"},
				{Align: simpletable.AlignCenter, Text: "Value"},
			},
		}
		row := []string{
			"Remote Report",
			"Report bugs/stats to our server to allow user experience improvements",
			strconv.FormatBool(config.RemoteReport),
		}
		items = append(items, row)
	} else if reqToList == "nodes" {
		nodes := entities.FetchNodes()
		headers = simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Node ID"},
				{Align: simpletable.AlignCenter, Text: "Node Name"},
				{Align: simpletable.AlignCenter, Text: "Node Address"},
			},
		}
		for _, node := range nodes {
			row := []string{
				node.ID,
				node.Name,
				node.Address,
			}
			items = append(items, row)
		}
	}
	return headers, items
}

func listItem(reqToList string) {
	listable := []string{"images", "containers", "config", "nodes"}
	if utilities.StringInSlice(reqToList, listable) {
		headers, items := listableToListItem(reqToList)
		baseList(headers, items)
	} else {
		utilities.Formatter.Error("cannot list " + reqToList + ", please check your parameters.")
	}
}
