package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Upload holds the schema definition for the Upload entity.
type Upload struct {
	ent.Schema
}

// Fields of the Upload.
func (Upload) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(70),
		field.UUID("uid", uuid.UUID{}),
		field.String("mime_type"),
		field.String("title"),
		field.Int("size"),
		field.Text("description"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Upload.
func (Upload) Edges() []ent.Edge {
	return nil
}
