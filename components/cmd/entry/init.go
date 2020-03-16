package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"

	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/requests"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() {
	db := storage.GetDefaultDB()
	db.Connect()
	db.CreateTables()
	logger.Info("Successfully initialized database")
}

// put everything (logs, packages, etc) under ~/.autoai/.aid
func initFolder() {
	utilities.InitFolders()
	logger.Info("Successfully created required folders")
}

// insert default values to the db
func addDefaultToDB() {
	entities.SaveInitLogObject()
	logger.Info("Successfully add required objects into database")
}

func initRepository() {
	prompt := promptui.Prompt{
		Label: "Your Package Name",
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}
	gitClient := requests.NewGitClient()
	gitClient.Clone("https://github.com/aidmodels/bolierplate", result)
	// rename {package_name} to real package name
	pyFolderName := filepath.Join(result, "{package_name}")
	err = os.Rename(pyFolderName, filepath.Join(result, result))
	if err != nil {
		fmt.Println(err)
	}
}
