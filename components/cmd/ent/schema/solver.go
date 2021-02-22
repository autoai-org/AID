// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Solver Schema
type Solver struct {
	ent.Schema
}

// Fields of the solver
func (Solver) Fields() []ent.Field {
	return []ent.Field{
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
			Unique(),
		edge.From("image", Image.Type).Unique().Ref("solver"),
	}
}
