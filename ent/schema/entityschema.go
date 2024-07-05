package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntitySchema holds the schema definition for the EntitySchema entity. This entity
// describes the schema of an EntityType, and ultimately the schema of an Entity.
// Because entities can be created/updated with a schema that is then later changed,
// we need to keep track of every schema that an EntityType has had. This allows us to
// validate that an Entity conforms to the schema that was in place when it was created,
// and provide a way to migrate Entities to a new schema.
type EntitySchema struct {
	ent.Schema
}

// Fields of the EntitySchema.
func (EntitySchema) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("The ID of the entity schema."),
	}
}

// Edges of the EntitySchema.
func (EntitySchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entityschemaversions", EntitySchemaVersion.Type).Comment("The versions of the entity schema."),
	}
}
