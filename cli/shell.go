package main

import (
	"log"
	"os/exec"
)

func pip(args []string) {
	config := readConfig()
	localPip := config.Local.Pip
	_ = _execCommand(localPip, args)
}

func _execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Printf(string(output))
	return true
}
