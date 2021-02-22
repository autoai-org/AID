// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// ColorPrinter is the struct for colored printting
type ColorPrinter struct {
	// TODO: abstract all printers with color definition
	Name string
}

// Formatter is the extern object that we should use for color printing
var Formatter *ColorPrinter

// Success prints a green "success" with other information
func (cp *ColorPrinter) Success(msg string, rest ...interface{}) {
	s := aurora.Green("success")
	all := append([]interface{}{s}, rest...)
	fmt.Printf("%s "+msg+"\n", all)
}

// Progress prints a green "success" with other information
func (cp *ColorPrinter) Progress(msg string, rest ...interface{}) {
	s := aurora.Cyan("in progress")
	all := append([]interface{}{s}, rest...)
	fmt.Printf("%s "+msg+"\n", all)
}

// Warning prints a green "success" with other information
func (cp *ColorPrinter) Warning(msg string, rest ...interface{}) {
	s := aurora.Yellow("warning")
	all := append([]interface{}{s}, rest...)
	fmt.Printf("%s "+msg+"\n", all)
}

// Progress prints a green "success" with other information
func (cp *ColorPrinter) Error(msg string, rest ...interface{}) {
	s := aurora.Red("error")
	all := append([]interface{}{s}, rest...)
	fmt.Printf("%s "+msg+"\n", all)
}
