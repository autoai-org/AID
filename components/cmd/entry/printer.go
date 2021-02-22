// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	"github.com/flosch/pongo2"
)

func renderFile(templateFile string, context map[string]interface{}) error {
	fileContent, err := utilities.ReadFileContent(templateFile)
	utilities.CheckError(err, "Cannot read file "+templateFile)
	tpl, err := pongo2.FromString(fileContent)
	utilities.CheckError(err, "Failed to get template.")
	out, err := tpl.Execute(context)
	utilities.CheckError(err, "Failed to generate file.")
	utilities.WriteContentToFile(templateFile, out)
	return err
}
