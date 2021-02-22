package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// SystemLog schema.
type SystemLog struct {
	ent.Schema
}

// Fields of the log.
func (SystemLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("filepath"),
		field.String("source"),
		field.Time("created_at").
			Default(time.Now),
	}
}
