package schema

import "entgo.io/ent"

// Key holds the schema definition for the Key entity.
type Key struct {
	ent.Schema
}

// Fields of the Key.
func (Key) Fields() []ent.Field {
	return nil
}

// Edges of the Key.
func (Key) Edges() []ent.Edge {
	return nil
}
