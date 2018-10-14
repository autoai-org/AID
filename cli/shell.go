package main

import (
	"fmt"
	"log"
	"os/exec"
	"github.com/getsentry/raven-go"
)

func pip(args []string) {
	config := readConfig()
	localPip := config.Local.Pip
	_ = _execCommand(localPip, args)
}

func python(args []string) {
	config := readConfig()
	localPython := config.Local.Python
	_ = _execCommand(localPython, args)
}

func _execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	fmt.Println(cmd.Args)
	output, err := cmd.CombinedOutput()
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
		return false
	}
	log.Printf(string(output))
	return true
}
