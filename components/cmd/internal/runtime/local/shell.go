package local

import (
	"fmt"
	"os/exec"

	"github.com/autoai-org/aid/internal/utilities"
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
	// handles environment variables
	cmd := exec.Command(PythonPath, args...)
	err := cmd.Run()
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
}
