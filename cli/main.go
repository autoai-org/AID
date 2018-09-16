package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	// sessionToken := getCache("session-token")
	var currentUser User
	// if sessionToken != "" {
	//		currentUser = User{"", "", sessionToken}
	// currentUser.become()
	// }
	cvpm := cli.NewApp()
	cvpm.Name = "CVPM"
	cvpm.Usage = "Computer Vision Package Manager"
	cvpm.Commands = []cli.Command{
		{
			Name: "login",
			Action: func(c *cli.Context) error {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Username: ")
				username, _ := reader.ReadString('\n')
				username = strings.TrimSpace(username)
				fmt.Printf("Password: ")
				bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
				password := strings.TrimSpace(string(bytePassword))
				u := User{username, password, ""}
				currentUser = u.login()
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
		log.Fatal(err)
	}
}
