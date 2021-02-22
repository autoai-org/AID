// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cargo

import (
	"context"
	"fmt"
	"os"

	entRepository "github.com/autoai-org/aid/ent/generated/repository"
	"github.com/autoai-org/aid/internal/database"
	"github.com/autoai-org/aid/internal/utilities"
)

// RemovePackage deletes the source code and pretrained models of a package
// It also removes the database records.
func RemovePackage(packageID string) error {
	repo, err := database.NewDefaultDB().Repository.Query().Where(entRepository.UID(packageID)).First(context.Background())
	if err != nil {
		utilities.Formatter.Error("Cannot fetch repository " + repo.Vendor + "/" + repo.Name + "(" + fmt.Sprint(packageID) + ") from database.")
		os.Exit(3)
	}
	_, err = database.NewDefaultDB().Repository.Delete().Where(entRepository.UID(packageID)).Exec(context.Background())
	if err != nil {
		utilities.ReportError(err, "Cannot remove container")
		return err
	}
	utilities.Formatter.Info("Repository: " + repo.Vendor + "/" + repo.Name + "(" + fmt.Sprint(packageID) + ") deleted from database.")
	// Now remove files on the disk
	err = os.RemoveAll(repo.Localpath)
	if err != nil {
		utilities.Formatter.Error("Cannot delete the folder " + repo.Localpath + ": " + err.Error())
		os.Exit(6)
	}
	utilities.Formatter.Info("Repository: " + repo.Vendor + "/" + repo.Name + "(" + fmt.Sprint(packageID) + ") deleted from your disk.")
	return err
}
