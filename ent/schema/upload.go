package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int("user_id"),
		field.Int("playlist_id"),
		field.String("name").MaxLen(70),
		field.UUID("uid", uuid.UUID{}),
		field.String("mime_type"),
		field.Int("size"),
		field.String("title"),
		field.Text("description"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Upload.
func (Upload) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Required().Ref("uploads").Field("user_id").Unique(),
		edge.From("playlist", Playlist.Type).Field("playlist_id").Required().Ref("uploads").Unique(),
	}
}
