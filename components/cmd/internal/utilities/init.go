package utilities

import (
	"fmt"
	"path/filepath"
)

// InitFolders create all required folders under ~/.autoai/.aid
func InitFolders() {
	vendorDir := filepath.Join(GetHomeDir(), ".autoai")
	CreateFolderIfNotExist(vendorDir)
	aidDir := filepath.Join(vendorDir, "aid")
	CreateFolderIfNotExist(aidDir)
	requiredFolders := [5]string{"logs", "models", "plugins", "datasets", "temp"}
	for _, each := range requiredFolders {
		CreateFolderIfNotExist(filepath.Join(aidDir, each))
	}
	initConfig := SystemConfig{RemoteReport: true}
	SaveConfig(initConfig)
}

func init() {
	fmt.Println("Initialising folders")
	InitFolders()
}
