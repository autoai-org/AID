package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func pip(args []string) {
	config := readConfig()
	localPip := config.Local.Pip
	// _ = _execCommand(localPip, args)
	proc := NewProcess(localPip, args...)
	out := proc.StreamOutput()
	logFile := filepath.Join(getLogsLocation(), "system.log")
	log.Println(logFile)
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	go func() {
		for out.Scan() {
			log.Println(out.Text())
			_, err := file.WriteString(out.Text() + "\n")
			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			defer file.Close()
			log.Println(out.Text())
		}
	}()
	proc.Start()
}

func python(args []string) {
	config := readConfig()
	localPython := config.Local.Python
	//_ = _execCommand(localPython, args)
	proc := NewProcess(localPython, args...)
	out := proc.StreamOutput()
	logFile := filepath.Join(getLogsLocation(), "package.log")
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	go func() {
		for out.Scan() {
			_, err := file.WriteString(out.Text() + "\n")
			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
			defer file.Close()
			log.Println(out.Text())
		}
	}()
	proc.Start()
}

func _execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	cmd.Stdout = os.Stdout
	fmt.Println(cmd.Args)
	return true
}

