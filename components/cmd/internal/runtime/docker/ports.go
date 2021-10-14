// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/autoai-org/aid/internal/utilities"
)

// ExportImage save image into dest, for uploading to other nodes, etc.
// Before calling export, we should ask users if they want to overwrite existing file.
func ExportImage(imageName string) error {
	targetFile := filepath.Join(utilities.GetBasePath(), "temp", imageName+".aidimg")
	utilities.Formatter.Info("Exporting your image to " + targetFile)
	// Delete existing file
	if utilities.IsExists(targetFile) {
		err := os.Remove(targetFile)
		utilities.ReportError(err, "Cannot remove existing file "+targetFile)
	}
	file, err := os.Create(targetFile)
	if err != nil {
		utilities.ReportError(err, "Cannot create new image file at "+targetFile)
		return err
	}
	defer file.Close()
	resBody, err := Client.ImageSave(context.Background(), []string{imageName})
	_, err = io.Copy(file, resBody)
	if err != nil {
		utilities.ReportError(err, "Cannot write to the file: "+targetFile)
		return err
	}
	utilities.Formatter.Info("Exported your image to " + targetFile)
	return nil
}

// ImportImage reads images into docker runtime
func ImportImage(imageName string, quiet bool) error {
	targetFile := filepath.Join(utilities.GetBasePath(), "temp", imageName+".aidimg")
	utilities.Formatter.Info("Importing your image to " + targetFile)
	_, err := os.Stat(targetFile)
	if os.IsNotExist(err) {
		utilities.ReportError(err, "Cannot read the file: "+targetFile)
		return err
	}
	file, err := os.Open(targetFile)
	if err != nil {
		utilities.ReportError(err, "Cannot open the file: "+targetFile)
	}
	reader := bufio.NewReader(file)
	_, err = Client.ImageLoad(context.Background(), reader, quiet)
	if err != nil {
		utilities.ReportError(err, "Cannot import image from "+targetFile)
		return err
	}
	utilities.Formatter.Info("Imported your image from " + targetFile)
	return nil
}

// UploadImage uploads an image to a remote server such that the solver will be loaded on that server as well
func UploadImage(remoteAddr string, solverID string) error {
	return nil
}

// DownloadImage downloads an image remotely, such that the solver will be loaded in the local server
func DownloadImage(remoteAddr string, solverID string) error {
	return nil
}
