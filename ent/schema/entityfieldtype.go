package schema

import "entgo.io/ent"

// EntityFieldType holds the schema definition for the EntityFieldType entity.
type EntityFieldType struct {
	ent.Schema
}

// Fields of the EntityFieldType.
func (EntityFieldType) Fields() []ent.Field {
	return nil
}

// Edges of the EntityFieldType.
func (EntityFieldType) Edges() []ent.Edge {
	return nil
}
