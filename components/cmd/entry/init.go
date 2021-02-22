package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/manifoldco/promptui"

	"github.com/autoai-org/aid/components/cmd/pkg/entities"
	"github.com/autoai-org/aid/components/cmd/pkg/requests"
	"github.com/autoai-org/aid/components/cmd/pkg/storage"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
	aurora "github.com/logrusorgru/aurora"
	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() {
	db := storage.GetDefaultDB()
	db.Connect()
	db.CreateTables()
	utilities.Formatter.Success("initialized database")
}

// put everything (logs, packages, etc) under ~/.autoai/.aid
func initFolder() {
	utilities.InitFolders()
	utilities.Formatter.Success("created required folders")
}

// insert default values to the db
func addDefaultToDB() {
	entities.SaveInitLogObject()
	utilities.Formatter.Success("add required objects into database")
}

func initRepository() {
	promptVendor := promptui.Prompt{
		Label: "Your Vendor Name, it can be your name/company name, etc.",
	}
	vendorName, err := promptVendor.Run()
	promptPackage := promptui.Prompt{
		Label: "Your Package Name, it should be unique inside your vendor scope",
	}
	packageName, err := promptPackage.Run()
	promptDesc := promptui.Prompt{
		Label: "Description of your package",
	}
	description, err := promptDesc.Run()
	if err != nil {
		fmt.Println(err)
	}
	gitClient := requests.NewGitClient()
	gitClient.Clone("https://github.com/aidmodels/bolierplate", packageName)
	// rename {package_name} to real package name
	pyFolderName := filepath.Join(packageName, "{package_name}")
	err = os.Rename(pyFolderName, filepath.Join(packageName, packageName))
	// remove .git folder
	os.Remove(path.Join(packageName, ".git"))
	// render readme file, and aid.toml
	renderContext := map[string]interface{}{"vendorName": vendorName, "packageName": packageName, "description": description}
	renderFile(path.Join(packageName, "README.md"), renderContext)
	renderFile(path.Join(packageName, "aid.toml"), renderContext)
	// hint what's next
	fmt.Println(aurora.Green("Successfully created your package " + vendorName + "/" + packageName))
	fmt.Println(aurora.Green("To make it work, please follow the instructions:"))
	fmt.Printf("\t 1. Modify your solver inside %s \n", aurora.Cyan(packageName+"/"+packageName+"/solver.py"))
	fmt.Printf("\t 2. Modify your aid.toml file with your new solver name\n")
	fmt.Printf("\t 3. Modify your README.md file with your new solvers and its description\n")
	fmt.Printf("\t 4. You are ready to go! You can now upload your package to GitHub or our Model Hub\n")
	fmt.Println(aurora.Green("For more information, please visit https://aid.autoai.org/docs/package/publish-packages"))
	fmt.Println(aurora.Green("If you met problems/issues, feel free to post it to https://github.com/autaoi-org/aid"))
	if err != nil {
		fmt.Println(err)
	}
}
