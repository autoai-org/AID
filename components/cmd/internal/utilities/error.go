// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import "os"

// ReportError is used to capture errors, and pass to central analytics
// This function should be disabled, when report=False
// Disable telemetry until final release.
func ReportError(err error, errorMessage string) {
	if err != nil {
		Formatter.Error(errorMessage + ": " + err.Error())
		os.Exit(1)
	}
}
