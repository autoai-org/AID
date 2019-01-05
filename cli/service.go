// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"os/user"
	"runtime"

	"github.com/kardianos/service"
)

type sol struct {
}

func (s *sol) Start(srv service.Service) error {
	go runServer(DaemonPort)
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

func getCVPMDConfig() *service.Config {
	realUsername := getRunUser()
	srvConf := &service.Config{
		Name:        "cvpmd",
		DisplayName: "CVPM Daemon",
		Description: "Computer Vision Package Manager[Daemon]",
		Arguments:   []string{"daemon", "run"},
		UserName:    realUsername,
	}
	return srvConf
}

// cvpm daemon install -> install the background daemon service
func InstallService() {
	srvConfig := getCVPMDConfig()
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

// cvpm daemon uninstall -> uninstall the background daemon service
func UninstallService() {
	srvConfig := getCVPMDConfig()
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
