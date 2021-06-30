package initialization

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func createFolderIfNotExist(folderPath string) {
	exist := IsExists(folderPath)
	if !exist {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Print("error when creating " + folderPath + ": " + err.Error())
		}
	}
}

func init() {
	// The init function will be automatically called when the initialization package is imported.
	vendorDir := filepath.Join(getHomeDir(), ".autoai")
	createFolderIfNotExist(vendorDir)
	aidDir := filepath.Join(vendorDir, "aid")
	createFolderIfNotExist(aidDir)
	requiredFolders := [5]string{"logs", "models", "plugins", "datasets", "temp"}
	for _, each := range requiredFolders {
		createFolderIfNotExist(filepath.Join(aidDir, each))
	}
}
