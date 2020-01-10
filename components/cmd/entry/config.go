// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

var logger = utilities.NewDefaultLogger()

// readConfig checks if all the required folder exists,
// If not, it will create the folder

var userConfig *utilities.Config

func readConfig() {
	userConfig = utilities.NewConfig("./config.json")
	utilities.DefaultConfig = userConfig
}