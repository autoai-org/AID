/*
Package CVPM (main) implements the CLI for CVPM.
To get started, use:

cvpm help 

to get a detailed explanation.
*/
package main

import (
	"github.com/urfave/cli"
	"github.com/getsentry/raven-go"
	"log"
	"os"
)

func main() {
	validateConfig()
	var currentUser User
	cvpm := cli.NewApp()
	cvpm.Name = "CVPM"
	cvpm.Usage = "Computer Vision Package Manager"
	cvpm.Commands = []cli.Command{
		{
			Name: "login",
			Action: func(c *cli.Context) error {
				currentUser = LoginHandler(c)
				return nil
			},
		},
		{
			Name: "install",
			Action: func(c *cli.Context) error {
				InstallHandler(c)
				return nil
			},
		},
		{
			Name: "list",
			Action: func(c *cli.Context) error {
				listRepos(c)
				return nil
			},
		},
		{
			Name: "daemon",
			Action: func(c *cli.Context) error {
				DaemonHandler(c)
				return nil
			},
		},
		{
			Name: "repo",
			Action: func(c *cli.Context) error {
				RepoHandler(c)
				return nil
			},
		},
		{
			Name: "config",
			Action: func(c *cli.Context) error {
				ConfigHandler(c)
				return nil
			},
		},
	}
	err := cvpm.Run(os.Args)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
}
