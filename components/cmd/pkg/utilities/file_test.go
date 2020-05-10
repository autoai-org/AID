// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utilities

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWriteContentToFile(t *testing.T) {
	type args struct {
		filepath string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Testing Write Content to File",
			args: args{
				filepath: "./test.test",
				content:  "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteContentToFile(tt.args.filepath, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("WriteContentToFile() error = %v, wantErr %v", err, tt.wantErr)
				os.Remove(tt.args.filepath)
			}
		})
	}
}

func TestReadFileContent(t *testing.T) {
	testString := "test"
	err := WriteContentToFile("./test.test", testString)
	assert.Nil(t, err)
	content, _ := ReadFileContent("./test.test")
	assert.Equal(t, content, testString)
	os.Remove("./test.test")
}

func TestGetRemoteFile(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRemoteFile(tt.args.url); got != tt.want {
				t.Errorf("GetRemoteFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFileIfModified(t *testing.T) {
	type args struct {
		filename string
		lastMod  time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReadFileIfModified(tt.args.filename, tt.args.lastMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileIfModified() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileIfModified() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadFileIfModified() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFileExists(tt.args.filename); got != tt.want {
				t.Errorf("IsFileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
