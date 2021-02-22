// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import "testing"

func TestWriteCounter_Write(t *testing.T) {
	type fields struct {
		Total uint64
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := &WriteCounter{
				Total: tt.fields.Total,
			}
			got, err := wc.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteCounter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteCounter_PrintProgress(t *testing.T) {
	type fields struct {
		Total uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wc := WriteCounter{
				Total: tt.fields.Total,
			}
			wc.PrintProgress()
		})
	}
}

func TestDownload(t *testing.T) {
	type args struct {
		url  string
		dest string
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
			if err := Download(tt.args.url, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
