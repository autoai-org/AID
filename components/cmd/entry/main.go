// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	// Version will be automatically inserted when using build.sh
	Version string
	// Build will be automatically inserted when using build.sh
	Build string
)

func main() {
	readConfig()
	var license bool
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("VERSION:%s\nBUILD:%s\n", c.App.Version, Build)
	}
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "license",
				Usage:       "print the license",
				Destination: &license,
			},
		},
		Name:    "AID",
		Version: Version,
		Usage:   "One-Stop AIOps System.",
		Action: func(c *cli.Context) error {
			if license {
				printLicense()
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:     "db",
				Aliases:  []string{"db"},
				Usage:    "database",
				Category: "storage",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:     "list",
				Aliases:  []string{"ls"},
				Usage:    "aid list [image|container|config]",
				Category: "utility",
				Action: func(c *cli.Context) error {
					listItem(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:     "image",
				Aliases:  []string{"img"},
				Usage:    "Manage images - Export/Import/List prebuilt images",
				Category: "image",
				Subcommands: []*cli.Command{
					{
						Name:     "export",
						Aliases:  []string{"exp"},
						Usage:    "Export image into a tarfile",
						Category: "image",
						Action: func(c *cli.Context) error {
							exportImage(c.Args().Get(0))
							return nil
						},
					},
					{
						Name:     "import",
						Aliases:  []string{"imp"},
						Usage:    "Export image into a tarfile",
						Category: "image",
						Action: func(c *cli.Context) error {
							importImage(c.Args().Get(0), false)
							return nil
						},
					},
				},
			},
			{
				Name:     "cluster",
				Aliases:  []string{"cl"},
				Usage:    "Cluster Management -  Add Node/Upload Image/etc",
				Category: "cluster",
				Subcommands: []*cli.Command{
					{
						Name:     "agent",
						Aliases:  []string{"ag"},
						Usage:    "Run a daemon process to receive requests from master",
						Category: "cluster",
						Action: func(c *cli.Context) error {
							runAgent()
							return nil
						},
					},
					{
						Name:     "node",
						Usage:    "List/Add Nodes",
						Category: "cluster",
						Subcommands: []*cli.Command{
							{
								Name: "add",
								Action: func(c *cli.Context) error {
									addNode()
									return nil
								},
							},
							{
								Name: "list",
								Action: func(c *cli.Context) error {
									printNodes()
									return nil
								},
							},
						},
					},
					{
						Name:     "push",
						Aliases:  []string{"ps"},
						Usage:    "Push Image to target host",
						Category: "cluster",
						Action: func(c *cli.Context) error {
							pushImage(c.Args().Get(0), c.Args().Get(1))
							return nil
						},
					},
				},
			},
			{
				Name:     "config",
				Aliases:  []string{"conf"},
				Usage:    "Reset Configurations to default value",
				Category: "self",
				Subcommands: []*cli.Command{
					{
						Name:     "reset",
						Aliases:  []string{"ls"},
						Usage:    "Reset all configurations to default value",
						Category: "self",
						Action: func(c *cli.Context) error {
							resetConfig()
							return nil
						},
					},
				},
			},
			{
				Name:     "init",
				Usage:    "Initialize all configurations, database and web ui",
				Category: "self",
				Action: func(c *cli.Context) error {
					initFolder()
					initDatabase()
					addDefaultToDB()
					return nil
				},
			},
			{
				Name:     "build",
				Aliases:  []string{"b"},
				Usage:    "Build Image",
				Category: "runtime",
				Action: func(c *cli.Context) error {
					buildContext := c.Args().Get(0)
					if strings.Contains(buildContext, "/") {
						buildInfo := strings.Split(buildContext, "/")
						absoluteBuild(buildInfo[0], buildInfo[1], buildInfo[2])
					} else {
						build(buildContext)
					}
					return nil
				},
			},
			{
				Name:     "create",
				Aliases:  []string{"cc"},
				Usage:    "Create Container",
				Category: "runtime",
				Action: func(c *cli.Context) error {
					createSolverContainer(c.Args().Get(0), c.Args().Get(1))
					return nil
				},
			},
			{
				Name:     "install",
				Aliases:  []string{"i"},
				Usage:    "Install Package",
				Category: "packages",
				Action: func(c *cli.Context) error {
					installPackage(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:     "new",
				Usage:    "Initialize a new package",
				Category: "packages",
				Action: func(c *cli.Context) error {
					initRepository()
					return nil
				},
			},
			{
				Name:     "generate",
				Aliases:  []string{"gen"},
				Usage:    "Generate Runners and Dockerfile",
				Category: "runtime",
				Action: func(c *cli.Context) error {
					generate()
					return nil
				},
			},
			{
				Name:     "interactive",
				Aliases:  []string{"in"},
				Usage:    "Interactive Mode",
				Category: "self",
				Action: func(c *cli.Context) error {
					interactiveMode()
					return nil
				},
			},
			{
				Name:     "up",
				Usage:    "Server Up",
				Category: "daemon",
				Action: func(c *cli.Context) error {
					startServer(c.Args().Get(0))
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
