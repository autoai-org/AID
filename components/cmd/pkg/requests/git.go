// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	git "gopkg.in/src-d/go-git.v4"
	"os"
)

var logger = utilities.NewLogger()

// GitClient is the basic structure for performing git-based operations
type GitClient struct {
}

// Clone downloads remote contents from remoteURL to targetFolder
func (gitclient *GitClient) Clone(remoteURL string, targetFolder string) {
	logger.Info("Cloning from " + remoteURL + " into " + targetFolder)
	_, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL:      remoteURL,
		Progress: os.Stdout,
	})
	utilities.CheckError(err, "Cannot Clone from "+remoteURL)
}
