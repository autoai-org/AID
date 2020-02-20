// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Process is used for shell/cmd command
type Process struct {
	proc               *exec.Cmd
	cancellationSignal chan uint8
	done               chan error
	returnCode         chan error
	started            bool
	stdOutRead         *io.PipeReader
	stdOutWrite        *io.PipeWriter
	inputWriter        *io.PipeWriter
	inputStreamSet     bool
	outputStreamSet    bool
	completed          bool
	timeout            time.Duration
	// Access to completed MUST capture the lock.
	mutex sync.RWMutex
}

// NewProcess creates a new process for the specific command
func NewProcess(command string, envs string, args ...string) *Process {
	cmd := exec.Command(command, args...)
	// load system environments
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, envs)
	// parse package related envs
	process := &Process{
		cmd,
		make(chan uint8, 1),
		make(chan error, 1),
		make(chan error, 1),
		false,
		&io.PipeReader{},
		&io.PipeWriter{},
		&io.PipeWriter{},
		false,
		false,
		false,
		0,
		sync.RWMutex{}}
	return process
}

// Start runs the process with go-routine
func (p *Process) Start() *Process {
	if p.timeout > 0 {
		log.Println("Its greater than 0")
		timer := time.NewTimer(p.timeout)
		go func() {
			<-timer.C
			p.Kill()
		}()
	}
	p.started = true
	//Call the other functions to stream stdin and stdout
	err := p.proc.Start()
	if err != nil {
		panic(err)
	}
	go p.awaitOutput()
	go p.finishTimeOutOrDie()
	return p
}

// SetTimeout set a timeout duration for the process
// Started process cannot be set a timeout.
func (p *Process) SetTimeout(d time.Duration) {
	if p.started {
		panic("Can not set timeout after process started")
	}
	p.timeout = d
}

// Wait returns the returncode
func (p *Process) Wait() error {
	return <-p.returnCode
}

func (p *Process) awaitOutput() {
	//send the exit code to the done channel to signify the command finished
	p.done <- p.proc.Wait()
}

// Kill kills the process immediately
func (p *Process) Kill() {
	p.mutex.Lock()
	if !p.completed {
		p.cancellationSignal <- 1
	}
	p.mutex.Unlock()
}

// OpenInputStream returns a stream for the input
func (p *Process) OpenInputStream() (io.WriteCloser, error) {
	if p.inputStreamSet {
		panic("Input stream already set")
	}
	if p.started {
		panic("process already started")
	}
	stdIn, err := p.proc.StdinPipe()
	p.inputStreamSet = true
	return stdIn, err

}

// StreamOutput streams the output
func (p *Process) StreamOutput() *bufio.Scanner {
	//pipe both stdout and stderr into the same pipe
	//panics if you do streamoutput after process starting or
	//if the output stream is already set
	if p.started {
		panic("Cant set output stream after starting")
	}
	if p.outputStreamSet {
		panic("output stream already set!")
	}
	p.stdOutRead, p.stdOutWrite = io.Pipe()
	p.proc.Stdout = p.stdOutWrite
	p.proc.Stderr = p.stdOutWrite
	p.outputStreamSet = true
	//return a scanner which they can read from till empty
	return bufio.NewScanner(p.stdOutRead)
}

func (p *Process) finishTimeOutOrDie() {
	defer p.cleanup()
	var result error
	select {
	case result = <-p.done:
	case <-p.cancellationSignal:
		log.Println("received cancellationSignal")
		//NOT PORTABLE TO WINDOWS
		err := p.proc.Process.Kill()
		if err != nil {
			log.Println(err)
		}
	}
	p.returnCode <- result
}

func (p *Process) cleanup() {
	p.mutex.Lock()
	p.completed = true
	p.mutex.Unlock()
	if p.outputStreamSet {
		p.stdOutRead.Close()
		p.stdOutWrite.Close()
	}
	close(p.done)
	close(p.cancellationSignal)
}
