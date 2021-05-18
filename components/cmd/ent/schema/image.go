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
		field.String("uid"),
		field.String("title"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("solver", Solver.Type).Unique(),
	}
}

// Container schema
type Container struct {
	ent.Schema
}

// Fields of Container
func (Container) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid"),
		field.String("port"),
		field.Bool("running").Default(false),
		field.Time("created_at").Default(time.Now).Annotations(
			entgql.OrderField("CREATED_AT"),
		),
	}
}

// Edges of Container.
func (Container) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type),
	}
}
