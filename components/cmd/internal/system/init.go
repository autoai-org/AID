package system

import (
	"context"
	"log"
	"path/filepath"

	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
)

// initFolders create all required folders under ~/.autoai/.aid
func initFolders() {
	vendorDir := filepath.Join(utilities.GetHomeDir(), ".autoai")
	utilities.CreateFolderIfNotExist(vendorDir)
	targetDir := filepath.Join(vendorDir, "aid")
	utilities.CreateFolderIfNotExist(targetDir)
	requiredFolders := [5]string{"logs", "models", "plugins", "datasets", "temp"}
	for _, each := range requiredFolders {
		utilities.CreateFolderIfNotExist(filepath.Join(targetDir, each))
	}
	initConfig := SystemConfig{RemoteReport: true}
	SaveConfig(initConfig)
}

// migrateDB is performed everytime before the system is started
func migrateDB() {
	if err := database.NewDefaultDB().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// InitializeSystem is checked everytime the TUI is called.
func InitializeSystem() {
	migrateDB()
	initFolders()
}
