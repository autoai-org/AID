// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"context"
	"io"
	"os"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/docker/docker/api/types"
)

// Pull will pull an exisiting package
func Pull(imageName string) error {
	reader, err := Client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		utilities.ReportError(err, "Cannot pull image "+imageName)
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

// Push will push an existing package into a remote docker registry
func Push(imageName string) error {
	reader, err := Client.ImagePush(context.Background(), imageName, types.ImagePushOptions{})
	if err != nil {
		utilities.ReportError(err, "Cannot Push image "+imageName+": "+err.Error())
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}
