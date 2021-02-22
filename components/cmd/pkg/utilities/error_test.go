// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import "testing"

func TestCheckError(t *testing.T) {
	type args struct {
		err          error
		errorMessage string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error_test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckError(tt.args.err, tt.args.errorMessage)
		})
	}
}
