package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	ent "github.com/autoai-org/aid/ent/generated"
	generated1 "github.com/autoai-org/aid/internal/daemon/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated1.NewExecutableSchema(generated1.Config{
		Resolvers: &Resolver{client},
	})
}
