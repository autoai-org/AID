package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/autoai-org/aid/internal/initialization"

	"github.com/alexeyco/simpletable"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
)

func baseList(headers simpletable.Header, items [][]*simpletable.Cell) {
	table := simpletable.New()
	table.Header = &headers
	table.Body.Cells = items
	table.Println()
}

func listPackage() {
	repos, err := database.NewDefaultDB().Repository.
		Query().
		All(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch repositories: " + err.Error())
		os.Exit(3)
	}
	headers := simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Unique ID"},
			{Align: simpletable.AlignCenter, Text: "Vendor"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
		},
	}
	var rows [][]*simpletable.Cell
	for idx, repo := range repos {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprint(idx + 1)},
			{Align: simpletable.AlignCenter, Text: repo.UID},
			{Text: repo.Vendor},
			{Text: repo.Name},
			{Text: repo.Status},
			{Align: simpletable.AlignCenter, Text: repo.CreatedAt.Local().Format("2006/01/02 15:04:05")},
		}
		rows = append(rows, r)
	}
	baseList(headers, rows)
}

func listImages() {
	images, err := database.NewDefaultDB().Image.Query().All(context.Background())
	utilities.ReportError(err, "cannot fetch content: images")
	headers := simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Unique ID"},
			{Align: simpletable.AlignCenter, Text: "Name"},
		},
	}
	var rows [][]*simpletable.Cell
	for idx, image := range images {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprint(idx + 1)},
			{Text: image.UID},
			{Text: image.Title},
		}
		rows = append(rows, r)
	}
	baseList(headers, rows)
}

func listContainers() {
	containers, err := database.NewDefaultDB().Container.Query().All(context.Background())
	utilities.ReportError(err, "cannot fetch content: containers")
	headers := simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Unique ID"},
			{Align: simpletable.AlignCenter, Text: "Port"},
			{Align: simpletable.AlignCenter, Text: "Running"},
		},
	}
	var rows [][]*simpletable.Cell
	for idx, container := range containers {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprint(idx + 1)},
			{Text: container.UID},
			{Text: container.Port},
			{Text: fmt.Sprint(container.Running)},
		}
		rows = append(rows, r)
	}
	baseList(headers, rows)
}
func listRepositores() {
	repositories, err := database.NewDefaultDB().Repository.Query().All(context.Background())
	utilities.ReportError(err, "cannot fetch content: containers")
	headers := simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Unique ID"},
			{Align: simpletable.AlignCenter, Text: "Vendor"},
			{Align: simpletable.AlignCenter, Text: "Name"},
		},
	}
	var rows [][]*simpletable.Cell
	for idx, repo := range repositories {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprint(idx + 1)},
			{Text: repo.UID},
			{Text: repo.Vendor},
			{Text: fmt.Sprint(repo.Name)},
		}
		rows = append(rows, r)
	}
	baseList(headers, rows)
}
func listEntity(entityName string) {
	switch entityName {
	case "packages":
		listPackage()
	case "images":
		listImages()
	case "containers":
		listContainers()
	case "repositories":
		listRepositores()
	default:
		utilities.Formatter.Error("Unsupported Entity Name")
	}
}
