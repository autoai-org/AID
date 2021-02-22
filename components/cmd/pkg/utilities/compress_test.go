// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
package utilities

import "testing"

func TestUncompress(t *testing.T) {
	type args struct {
		compressedFile string
		folderPath     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Uncompress(tt.args.compressedFile, tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("Uncompress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
