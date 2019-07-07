// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utility

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
	"github.com/gen2brain/go-unarr"
	homedir "github.com/mitchellh/go-homedir"
)

func removeContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

// InstallWebUi -> Download Latest and Unzip
func InstallWebUi() bool {
	webuiUrl := "https://dl.bintray.com/autoai/CVFlow/dist.zip"
	homepath, _ := homedir.Dir()
	webuiFolder := filepath.Join(homepath, "cvpm", "webui")
	webuiZipfile := filepath.Join(webuiFolder, "dist.zip")
	// Clean ~/cvpm/webui
	removeContents(webuiFolder)
	fmt.Println("Downloading webui assets...")
	DownloadFile(webuiZipfile, webuiUrl)
	archiveFile, err := unarr.NewArchive(webuiZipfile)
	fmt.Println("Unzipping...")
	err = archiveFile.Extract(webuiFolder)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Finished! Open http://127.0.0.1:10590 for the webui.")
	return true
}