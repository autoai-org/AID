// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"path/filepath"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/requests"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/manifoldco/promptui"
)

func runAgent() {
	startImageReceiver()
}

func pushImage(imageName string, nodeName string) {
	imageFile := filepath.Join(utilities.GetBasePath(), "temp", imageName+".aidimg")
	node := entities.GetNodeByName(nodeName)
	httpClient := requests.NewClient()
	resp := httpClient.Upload("http://"+node.Address+":10591/upload", imageFile)
	fmt.Printf(resp.String())
}

func addNode() {
	promptName := promptui.Prompt{
		Label: "Your node name",
	}
	nodeName, err := promptName.Run()
	promptAddr := promptui.Prompt{
		Label: "Your node ip/domain address",
	}
	nodeAddr, err := promptAddr.Run()
	if err != nil {
		fmt.Println(err)
	}
	node := entities.Node{Name: nodeName, Address: nodeAddr}
	err = node.Save()
	utilities.CheckError(err, "Cannot save node into the database")
}
