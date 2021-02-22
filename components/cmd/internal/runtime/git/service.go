package git

import (
	"log"

	"github.com/autoai-org/aid/internal/utilities"
	"github.com/sosedoff/gitkit"
)

// GetService returns the git service
func GetService() *gitkit.Server {
	// Configure git hooks
	hooks := &gitkit.HookScripts{
		PreReceive: `echo "Push Requests Received!"`,
	}
	service := gitkit.New(gitkit.Config{
		Dir:        utilities.GetFolder(utilities.MODELSFOLDER),
		AutoCreate: true,
		AutoHooks:  true,
		Hooks:      hooks,
	})
	if err := service.Setup(); err != nil {
		log.Fatal(err)
	}
	return service
}
