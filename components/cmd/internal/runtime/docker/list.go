// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package docker

import (
	"context"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/docker/docker/api/types"
)

// ListImages returns all images that have been installed on the host
func ListImages() []types.ImageSummary {
	images, err := Client.ImageList(context.Background(), types.ImageListOptions{})
	utilities.ReportError(err, "cannot list images")
	return images
}
