// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

// LogContent is used for the daemon to wrap log content as json format
type LogContent struct {
	Content string `json:"content"`
}