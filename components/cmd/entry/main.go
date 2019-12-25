package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

var (
	// Version will be automatically inserted when using build.sh
	Version string
	// Build will be automatically inserted when using build.sh
	Build string
)

func main() {
	var license bool

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s build=%s\n", c.App.Version, Build)
	}
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "license",
				Usage:       "print the license",
				Destination: &license,
			},
		},
		Name:    "AIPM",
		Version: Version,
		Usage:   "The Package Manager for A.I. Models",
		Action: func(c *cli.Context) error {
			if license {
				printLicense()
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:     "build",
				Aliases:  []string{"b"},
				Usage:    "Build Image",
				Category: "runtime",
				Action: func(c *cli.Context) error {
					build()
					return nil
				},
			},
			{
				Name:     "images",
				Usage:    "List Image",
				Category: "runtime",
				Action: func(c *cli.Context) error {
					printImages()
					return nil
				},
			},
			{
				Name:     "containers",
				Usage:    "List Containers",
				Category: "runtime",
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
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
