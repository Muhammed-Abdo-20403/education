package schema

import "entgo.io/ent"

// Playlist holds the schema definition for the Playlist entity.
type Playlist struct {
	ent.Schema
}

// Fields of the Playlist.
func (Playlist) Fields() []ent.Field {
	return nil
}

// Edges of the Playlist.
func (Playlist) Edges() []ent.Edge {
	return nil
}
