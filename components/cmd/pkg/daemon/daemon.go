// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/kardianos/service"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// DaemonPort is the default port that the daemon will be listening at
const DaemonPort = "10590"

type sol struct {
}

func (s *sol) Start(srv service.Service) error {
	go RunServer(DaemonPort, filepath.Join("./", "../../entry/server.pem"), filepath.Join("./", "../../entry/server.key"))
	return nil
}

func (s *sol) Stop(srv service.Service) error {
	return nil
}

func getRunUser() string {
	currentUser, _ := user.Current()
	if currentUser.Username == "root" && runtime.GOOS != "windows" {
		return os.Getenv("SUDO_USER")
	}
	return currentUser.Username
}

func getAIDConfig() *service.Config {
	realUsername := getRunUser()
	srvConf := &service.Config{
		Name:        "aidaemon",
		DisplayName: "AID Daemon",
		Description: "Artificial Intelligence Ops System[Daemon]",
		Arguments:   []string{"daemon", "up"},
		UserName:    realUsername,
	}
	return srvConf
}

// Install adds the server into the background service
func Install() {
	srvConfig := getAIDConfig()
	dae := &sol{}
	s, err := service.New(dae, srvConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Install()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// Uninstall removes the background daemon service
func Uninstall() {
	srvConfig := getAIDConfig()
	dae := &sol{}
	s, err := service.New(dae, srvConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Stop()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Uninstall()
	if err != nil {
		log.Fatal(err)
	}
}
