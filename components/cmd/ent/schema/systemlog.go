package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SystemLog schema.
type SystemLog struct {
	ent.Schema
}

// Fields of the log.
func (SystemLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Annotations(
			entgql.OrderField("UID"),
		),
		field.String("title"),
		field.String("filepath"),
		field.String("source"),
		field.Time("created_at").
			Default(time.Now),
	}
}
