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
	Build   string
)

func main() {
  cli.VersionPrinter = func(c *cli.Context) {
    fmt.Printf("version=%s build=%s\n", c.App.Version, Build)
  }
  app := &cli.App{
    Name: "AIPM",
    Version: Version,
    Usage: "The Package Manager for A.I. Models",
    Action: func(c *cli.Context) error {
      fmt.Println("Hello friend!")
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}