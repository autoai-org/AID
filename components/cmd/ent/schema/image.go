// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image schema
type Image struct {
	ent.Schema
}

// Fields of Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.Text("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.Text("title").Annotations(
			entgql.OrderField("TITLE"),
		),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges of Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solver", Solver.Type).Unique().Annotations(entgql.MapsTo("solver")),
		edge.To("container", Container.Type).Unique().Annotations(entgql.MapsTo("container")),
	}
}
