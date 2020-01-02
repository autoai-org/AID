// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "download", Description: "Download"},
		{Text: "deploy", Description: "Deploy"},
		{Text: "build", Description: "Build"},
		{Text: "exit", Description: "Exit"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func interactiveMode() {
	fmt.Println("Welcome to Interactive Mode! Version: " + Version)
	fmt.Println("Type exit to exit.")
	for {
		t := prompt.Input("> ", completer)
		fmt.Println("You selected " + t)
		if t == "exit" {
			break
		}
	}
}
