package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(70),
		field.String("channel_name").MaxLen(70).Unique(),
		field.String("email").Unique(),
		field.String("specialist").Optional(),
		field.Int("age").Optional().Positive(),
		field.String("phone").Optional(),
		field.String("language").Optional(),
		field.String("country"),
		field.Text("shor_bio").MaxLen(270),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("uploads", Upload.Type),
	}
}
