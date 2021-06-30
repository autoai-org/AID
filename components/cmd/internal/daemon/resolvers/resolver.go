package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	ent "github.com/autoai-org/aid/ent/generated"
	generated "github.com/autoai-org/aid/internal/daemon/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct{ Client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
