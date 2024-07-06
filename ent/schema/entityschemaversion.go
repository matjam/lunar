package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EntitySchemaVersion holds the schema definition for the EntitySchemaVersion entity. These
// are concrete versions of the schema for each EntitySchema. The EntitySchemaVersion provides
// a way to validate that an Entity conforms to the schema that was in place when it was created,
// and provide a way to migrate Entities to a new schema.
type EntitySchemaVersion struct {
	ent.Schema
}

// Fields of the EntitySchemaVersion.
func (EntitySchemaVersion) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("The ID of the entity schema version."),
		field.JSON("schema", map[string]interface{}{}).Comment("The schema of the entity schema version."),
		field.Time("created_at").Immutable().Comment("The time the entity schema version was created. Note that this is not the time the schema was created, but the time the version was created."),
		field.String("note").Optional().Comment("A note about the entity schema version."),
	}
}

// Edges of the EntitySchemaVersion.
func (EntitySchemaVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("created_by", User.Type).Unique().Comment("The user who created the entity schema version."),
		edge.From("entityschema", EntitySchema.Type).
			Ref("entityschemaversions").
			Unique().
			Comment("The schema that this version is for."),
		edge.To("entities", Entity.Type).Comment("The entities that conform to this version of the schema."),
	}
}

// Entities can only reference a single entityschema, and each Entity can only conform to a single version of the schema.
// EntitySchemas can have multiple versions, and each version can have multiple entities that conform to it.
// The EntitySchemaVersion entity is used to track the schema of each version of an EntitySchema.
