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

// Container schema
type Container struct {
	ent.Schema
}

// Fields of Container
func (Container) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.String("port").Annotations(
			entgql.OrderField("port"),
		),
		field.Bool("running").Default(false).Annotations(entgql.OrderField("running")),
		field.Time("created_at").Default(time.Now).Annotations(
			entgql.OrderField("created_at"),
		),
	}
}

// Edges of Container.
func (Container) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image", Image.Type).Ref("container").Unique().Annotations(entgql.MapsTo("image")),
	}
}
