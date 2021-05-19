package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Repository schema.
type Repository struct {
	ent.Schema
}

// Fields of the user.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.String("vendor").Annotations(
			entgql.OrderField("VENDOR"),
		),
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("status").Annotations(
			entgql.OrderField("STATUS"),
		),
		field.String("remote_url").
			Unique().Annotations(
			entgql.OrderField("REMOTE_URL"),
		),
		field.String("localpath").
			Unique().Annotations(
			entgql.OrderField("LOCALPATH"),
		),
		field.Time("created_at").
			Default(time.Now).Annotations(
			entgql.OrderField("CREATEDAT"),
		),
	}
}

// Edges of the repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solvers", Solver.Type),
	}
}
