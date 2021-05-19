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
		field.Text("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.Text("vendor").Annotations(
			entgql.OrderField("VENDOR"),
		),
		field.Text("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.Text("status").Annotations(
			entgql.OrderField("STATUS"),
		),
		field.Text("remote_url").
			Unique().Annotations(
			entgql.OrderField("REMOTE_URL"),
		),
		field.Text("localpath").
			Unique().Annotations(
			entgql.OrderField("LOCALPATH"),
		),
		field.Time("created_at").
			Default(time.Now).Annotations(
			entgql.OrderField("CREATED_AT"),
		),
	}
}

// Edges of the repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solvers", Solver.Type),
	}
}
