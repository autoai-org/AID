// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"
)

func pip(args []string) {
	config := readConfig()
	localPip := config.Local.Pip
	proc := NewProcess(localPip, "", args...)
	out := proc.StreamOutput()
	logFile := filepath.Join(getLogsLocation(), "system.log")
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
		}
	}()
	proc.Start()
}

func python(args []string, envs string) {
	config := readConfig()
	localPython := config.Local.Python
	// handles environment variables
	proc := NewProcess(localPython, "", args...)
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
