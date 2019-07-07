// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD & AutoAI.org
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utility

import (
	"os"
	"io"
	"net/http"
)

// DownloadFile will download a url to a local file.
func DownloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}