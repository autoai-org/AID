// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aid/components/cmd/pkg/entities"
)

// logContent is used for the daemon to wrap log content as json format
type logContent struct {
	Content string `json:"content"`
}

// installPackageRequest is used for client to form post requests
type installPackageRequest struct {
	RemoteURL string `json:"remote_url"`
}

// modifySolverDockerfileRequest is used to modify the dockerfile
type modifySolverDockerfileRequest struct {
	Content string `json:"content"`
}

type messageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// metaResponse is used to respond meta information to client
type metaResponse struct {
	Solvers      entities.Solvers     `json:"solvers"`
	Pretraineds  entities.Pretraineds `json:"pretraineds"`
	Readme       string               `json:"readme"`
	Requirements string               `json:"requirements"`
}

type newWebhookRequest struct {
	PayloadURL string `json:"payload"`
	Status     string `json:"status"`
}

type newDatasetRequest struct {
	DatasetPath string `json:"datasetPath"`
	Uncompress  bool   `json:"uncompress"`
	Name        string `json:"name"`
}

type newExperimentRequest struct {
	DatasetPath string `json:"datasetPath"`
	Vendor      string `json:"vendor"`
	Package     string `json:"package"`
	Solver      string `json:"solver"`
}
