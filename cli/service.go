package main

import (
	"github.com/kardianos/service"
	"log"
	"os/user"
	"os"
	"runtime"
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

func getCVPMDConfig() *service.Config {
	currentUser, _ := user.Current()
	var realUsername string
	if currentUser.Username == "root" && runtime.GOOS != "windows" {
		realUsername = os.Getenv("SUDO_USER")	
	} else {
		realUsername = currentUser.Username
	}
	srvConf := &service.Config{
		Name:        "cvpmd",
		DisplayName: "CVPM Daemon",
		Description: "Computer Vision Package Manager[Daemon]",
		Arguments:   []string{"daemon", "run"},
		UserName: realUsername,
	}
	return srvConf
}

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
