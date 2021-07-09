// Copyright (c) 2021 Xiaozhe Yao
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package requests

import (
	"github.com/autoai-org/aid/internal/utilities"
	"github.com/go-git/go-git/v5"
)

var logger = utilities.NewDefaultLogger()

var defaultGitClient *GitClient

// GitClient is the basic structure for performing git-based operations
type GitClient struct {
}

// NewGitClient returns a new Git Client
func NewGitClient() *GitClient {
	if defaultGitClient == nil {
		defaultGitClient = &GitClient{}
	}
	return defaultGitClient
}

// Clone downloads remote contents from remoteURL to targetFolder
func (gitclient *GitClient) Clone(remoteURL string, targetFolder string) error {
	utilities.Formatter.Info("Cloning from " + remoteURL + " into " + targetFolder)
	_, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL: remoteURL,
	})
	utilities.ReportError(err, "Cannot Clone from "+remoteURL)
	return err
}
