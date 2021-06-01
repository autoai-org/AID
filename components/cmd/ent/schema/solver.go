// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Solver Schema
type Solver struct {
	ent.Schema
}

// Fields of the solver
func (Solver) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.String("name"),
		field.String("class"),
		field.String("status"),
	}
}

// Edges of the solver.
func (Solver) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("solvers").
			Unique().Annotations(entgql.MapsTo("repository")),
		edge.From("image", Image.Type).Unique().Ref("solver").Annotations(entgql.MapsTo("image")),
	}
}
