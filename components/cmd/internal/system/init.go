package system

import (
	"context"
	"fmt"
	"log"

	"github.com/autoai-org/aid/internal/database"
)

// migrateDB is performed everytime before the system is started
func migrateDB() {
	if err := database.NewDefaultDB().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// InitializeSystem is checked everytime the TUI is called.
func Init() {
	fmt.Println("Initialising system....")
	migrateDB()
}
