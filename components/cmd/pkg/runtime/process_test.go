// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package runtime

import (
	"bufio"
	"io"
	"os/exec"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewProcess(t *testing.T) {
	type args struct {
		command string
		envs    string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want *Process
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProcess(tt.args.command, tt.args.envs, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcess_Start(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *Process
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			if got := p.Start(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcess_SetTimeout(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			p.SetTimeout(tt.args.d)
		})
	}
}

func TestProcess_Wait(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			if err := p.Wait(); (err != nil) != tt.wantErr {
				t.Errorf("Process.Wait() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProcess_awaitOutput(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			p.awaitOutput()
		})
	}
}

func TestProcess_Kill(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			p.Kill()
		})
	}
}

func TestProcess_OpenInputStream(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    io.WriteCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			got, err := p.OpenInputStream()
			if (err != nil) != tt.wantErr {
				t.Errorf("Process.OpenInputStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process.OpenInputStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcess_StreamOutput(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *bufio.Scanner
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			if got := p.StreamOutput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process.StreamOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcess_finishTimeOutOrDie(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			p.finishTimeOutOrDie()
		})
	}
}

func TestProcess_cleanup(t *testing.T) {
	type fields struct {
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
		mutex              sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				proc:               tt.fields.proc,
				cancellationSignal: tt.fields.cancellationSignal,
				done:               tt.fields.done,
				returnCode:         tt.fields.returnCode,
				started:            tt.fields.started,
				stdOutRead:         tt.fields.stdOutRead,
				stdOutWrite:        tt.fields.stdOutWrite,
				inputWriter:        tt.fields.inputWriter,
				inputStreamSet:     tt.fields.inputStreamSet,
				outputStreamSet:    tt.fields.outputStreamSet,
				completed:          tt.fields.completed,
				timeout:            tt.fields.timeout,
				mutex:              tt.fields.mutex,
			}
			p.cleanup()
		})
	}
}
