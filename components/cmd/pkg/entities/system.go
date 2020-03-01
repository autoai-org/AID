// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import "time"

// RunningError is used to save all errors that occured during a single running
type RunningError struct {
	Msg        string
	ErrorMsg   string
	OccurredAt time.Time
}

// RunningErrors is an array of RunningErrors. it is saved in the memory,
// so it will be dropped when stopping the program.
var RunningErrors []RunningError

// NewRunningErrors saves a new error into the array
func NewRunningErrors(err error, msg string) {
	RunningErrors = append(RunningErrors, RunningError{Msg: msg, ErrorMsg: err.Error(), OccurredAt: time.Now()})
}
