// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

// logContent is used for the daemon to wrap log content as json format
type logContent struct {
	Content string `json:"content"`
}

type installPackageRequest struct {
	RemoteURL string `json:"remote_url"`
}

type messageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
