package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Version will be automatically inserted when using build.sh
	Version string
	// Build will be automatically inserted when using build.sh
	Build string
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s build=%s\n", c.App.Version, Build)
	}
	app := &cli.App{
		Name:    "AIPM",
		Version: Version,
		Usage:   "The Package Manager for A.I. Models",
		Commands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build Image",
				Action: func(c *cli.Context) error {
					build()
					return nil
				},
			},
			{
				Name:  "images",
				Usage: "List Image",
				Action: func(c *cli.Context) error {
					printImages()
					return nil
				},
			},
			{
				Name:  "containers",
				Usage: "List Containers",
				Action: func(c *cli.Context) error {
					printContainers()
					return nil
				},
			},
			{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "Interactive Mode",
				Action: func(c *cli.Context) error {
					interactiveMode()
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
