package schema

import "entgo.io/ent"

// EntityType holds the schema definition for the EntityType entity.
type EntityType struct {
	ent.Schema
}

// Fields of the EntityType.
func (EntityType) Fields() []ent.Field {
	return nil
}

// Edges of the EntityType.
func (EntityType) Edges() []ent.Edge {
	return nil
}
