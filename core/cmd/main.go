// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package CVPM (main) implements the CLI for CVPM.
To get started, use:

cvpm help

to get a detailed explanation.
*/
package main

import (
	"github.com/getsentry/raven-go"
	"github.com/unarxiv/cvpm/pkg/config"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	Version string
	Build   string
)

func main() {
	config.Validate()
	cvpmapp := cli.NewApp()
	cvpmapp.Name = "CVPM"
	cvpmapp.Usage = "Computer Vision Package Manager"
	cvpmapp.Version = "0.0.4@alpha"
	cvpmapp.Commands = []cli.Command{
		{
			Name: "login",
			Action: func(c *cli.Context) error {
				LoginHandler(c)
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
	err := cvpmapp.Run(os.Args)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}
}
