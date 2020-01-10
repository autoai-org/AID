// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

// CheckError checks if error is not nil, and if it is nil it will logger the information
func CheckError(err error, errorMessage string) {
	if err != nil {
		logger.Error(errorMessage)
		logger.Error(err.Error())
	}
}
