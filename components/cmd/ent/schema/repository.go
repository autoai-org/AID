package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Repository schema.
type Repository struct {
	ent.Schema
}

// Fields of the user.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid"),
		field.String("vendor"),
		field.String("name"),
		field.String("status"),
		field.String("remote_url").
			Unique(),
		field.String("localpath").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solvers", Solver.Type),
	}
}
