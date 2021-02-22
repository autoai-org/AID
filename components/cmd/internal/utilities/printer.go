// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"fmt"

	"github.com/gookit/color"
)

// ColorPrinter is the primary object for printing logs into terminals.
type ColorPrinter struct {
}

// Formatter is the extern object that we should use for color printing
var Formatter *ColorPrinter

// Info is the shortcut for basePrint(info,...)
func (cp *ColorPrinter) Info(msg string) {
	cp.basePrint("INFO", msg)
}

// Error is the shorcut for basePrint(error,...)
func (cp *ColorPrinter) Error(msg string) {
	cp.basePrint("ERROR", msg)
}

// Warn is the shorcut for basePrint(warn,...)
func (cp *ColorPrinter) Warn(msg string) {
	cp.basePrint("WARN", msg)
}

func (cp *ColorPrinter) basePrint(level string, msg string) {
	var lvlColor color.Color
	switch level {
	case "ERROR":
		lvlColor = color.FgRed
	case "INFO":
		lvlColor = color.FgGreen
	case "WARN":
		lvlColor = color.FgYellow
	default:
		lvlColor = color.FgWhite
	}
	fmt.Printf("[%s] %s\n", lvlColor.Render(level), msg)
}
