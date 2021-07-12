package workflow

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/autoai-org/aid/internal/runtime/requests"
	"github.com/autoai-org/aid/internal/utilities"
)

// InitNewRepo
func InitNewRepo(vendorName string, repoName string, description string) {
	bolierplateURL := "https://github.com/aidmodels/bolierplate.git"
	git := requests.NewGitClient()
	err := git.Clone(bolierplateURL, repoName)
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	pyFolderName := filepath.Join(repoName, "{package_name}")
	err = os.Rename(pyFolderName, filepath.Join(repoName, repoName))
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	readmeContent, err := utilities.ReadFileContent(filepath.Join(repoName, "README.md"))
	if err != nil {
		utilities.Formatter.Error(err.Error())
	}
	output := bytes.Replace([]byte(readmeContent), []byte("{{vendorName}}"), []byte(vendorName), -1)
	output = bytes.Replace(output, []byte("{{packageName}}"), []byte(repoName), -1)
	output = bytes.Replace(output, []byte("{{description}}"), []byte(description), -1)
	if err = ioutil.WriteFile(filepath.Join(repoName, "README.md"), output, 0666); err != nil {
		utilities.Formatter.Error(err.Error())
	}
}
