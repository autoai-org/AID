package local

import (
	"fmt"
	"log"
)

var (
	PythonPath = "python"
	PipPath    = "pip"
)

// Pip calls the Pip program
func Pip(args []string) {
	proc := NewProcess(PipPath, "", args...)
	out := proc.StreamOutput()
	go func() {
		for out.Scan() {
			fmt.Println(out.Text())
		}
	}()
	proc.Start()
}

// Python calls the Python Program
func Python(args []string, envs string) {
	log.Println(args)
	// handles environment variables
	proc := NewProcess(PythonPath, "", args...)
	out := proc.StreamOutput()
	go func() {
		for out.Scan() {
			fmt.Println(out.Text())
		}
	}()
	proc.Start()
}
