// Copyright (c) 2021 Xiaozhe Yao
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"

	"github.com/autoai-org/aid/internal/initialization"
	"github.com/autoai-org/aid/internal/runtime/docker"
	"github.com/autoai-org/aid/internal/utilities"
)

var (
	// Version will be automatically inserted when building with Makefile
	Version string
	// Build will be automatically inserted when building with Makefile
	Build string
)

func main() {
	if !initialization.Initialized {
		cli.Exit("Failed to start", 1)
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Version: %s Build: %s\n", c.App.Version, Build)
	}
	app := &cli.App{
		Name:                 "AID",
		Version:              Version,
		Usage:                "One-Stop System for Machine Learning",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Value:       false,
				Usage:       "Enable detailed logs",
				Destination: &utilities.Verbose,
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:     "install",
				Usage:    "aid install [remote address]",
				Category: "packages",
				Action: func(c *cli.Context) error {
					installPackage(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:     "new",
				Usage:    "aid new",
				Category: "packages",
				Action: func(c *cli.Context) error {
					initRepo(c)
					return nil
				},
			},
			{
				Name:     "build",
				Usage:    "aid build [vendor]/[package]/[solver] or aid build -p [path]",
				Category: "packages",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Value:   "null",
						Usage:   "Path to the solver folder",
					},
					&cli.StringFlag{
						Name:    "solver",
						Aliases: []string{"sol"},
						Value:   "all",
						Usage:   "Name of the target solver",
					},
					&cli.BoolFlag{
						Name:  "remove",
						Value: false,
						Usage: "Remove the images after building",
					},
					&cli.BoolFlag{
						Name:    "gpu",
						Aliases: []string{"g"},
						Value:   false,
						Usage:   "GPU Required",
					},
				},
				Action: func(c *cli.Context) error {
					if c.String("path") == "null" {
						buildImage(c.Args().Get(0), c.Bool("gpu"))
					} else {
						buildImageByPath(c.String("path"), c.String("solver"), c.Bool("remove"))
					}
					//
					return nil
				},
			},
			{
				Name:     "create",
				Usage:    "aid create [Image Unique ID] [Host Port]",
				Category: "packages",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "gpu",
						Aliases: []string{"g"},
						Value:   false,
						Usage:   "GPU Required",
					},
				},
				Action: func(c *cli.Context) error {
					createContainer(c.Args().Get(0), c.Args().Get(1), c.Bool("gpu"))
					return nil
				},
			},
			{
				Name:     "start",
				Usage:    "aid start [Container Unique ID]",
				Category: "packages",
				Action: func(c *cli.Context) error {
					startContainer(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "infer",
				Usage: "Perform Inference",
				Action: func(c *cli.Context) error {
					infer(c.Args().Get(0), c.Args())
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List Entities",
				Action: func(c *cli.Context) error {
					listObject(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "help",
				Usage: "Package Help",
				Action: func(c *cli.Context) error {
					help(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "Stop a running container",
				Action: func(c *cli.Context) error {
					docker.Stop(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove an entity, i.e. package, container or image.",
				Action: func(c *cli.Context) error {
					remove(c.Args().Get(0), c.Args().Get(1))
					return nil
				},
			},
			{
				Name:     "up",
				Usage:    "Server Up",
				Category: "daemon",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "headless",
						Aliases: []string{"hl"},
						Value:   false,
						Usage:   "Headless mode",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("headless") {
						fmt.Println("headless mode")
						headlessDaemon(c)
					} else {
						startServer(c.Args().Get(0))
					}
					return nil
				},
			},
			{
				Name:  "generate",
				Usage: "Generate required files",
				Action: func(c *cli.Context) error {
					generate(c)
					return nil
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}
