// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"os"
	"log"
	"path/filepath"
	"github.com/unarxiv/cvpm/pkg/config"
)

// Pip calls the Pip program
func Pip(args []string) {
	localConfig := config.Read()
	localPip := localConfig.Local.Pip
	proc := NewProcess(localPip, "", args...)
	out := proc.StreamOutput()
	logFile := filepath.Join(config.GetLogLocation(), "system.log")
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

// Python calls the Python Program
func Python(args []string, envs string) {
	log.Println(args)
	localConfig := config.Read()
	localPython := localConfig.Local.Python
	// handles environment variables
	proc := NewProcess(localPython, "", args...)
	out := proc.StreamOutput()
	logFile := filepath.Join(config.GetLogLocation(), "package.log")
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
