// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"context"
	"os"

	entContainer "github.com/autoai-org/aid/ent/generated/container"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/levigross/grequests"
)

var defaultHTTPClient *HTTPClient

// HTTPClient is the basic structure for performing http requests
type HTTPClient struct {
}

// NewHTTPClient returns a new HTTP Client
func NewHTTPClient() *HTTPClient {
	if defaultHTTPClient == nil {
		defaultHTTPClient = &HTTPClient{}
	}
	return defaultHTTPClient
}

// Infer makes a http request to the given solver
func (httpclient *HTTPClient) Infer(containerID string, params map[string]string) (*grequests.Response, error) {
	containerEnt, err := database.NewDefaultDB().Container.Query().Where(entContainer.UID(containerID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch container: " + err.Error())
		os.Exit(3)
	}
	if containerEnt.Running != true {
		utilities.Formatter.Error("Container is not running, Aborted")
		os.Exit(4)
	}
	resp, err := grequests.Post("http://127.0.0.1:"+containerEnt.Port+"/infer", &grequests.RequestOptions{
		Data: params,
	})
	return resp, err
}
