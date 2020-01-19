// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewDefaultLogger()

func readConfig() {
	utilities.DefaultConfig = utilities.NewDefaultConfig()
}

func resetConfig() {
	config := utilities.SystemConfig{
		RemoteReport: true,
	}
	utilities.SaveConfig(config)
}