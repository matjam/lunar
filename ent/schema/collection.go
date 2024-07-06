package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("The ID of the collection."),
		field.String("name").NotEmpty().Comment("The name of the collection."),
		field.String("description").Optional().Comment("The description of the collection."),
		field.Time("created_at").Immutable().Comment("The time the collection was created."),
		field.Time("updated_at").Optional().Comment("The time the collection was last updated."),
		field.Time("deleted_at").Optional().Comment("The time the collection was deleted."),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entities", Entity.Type).Comment("The entities in the collection."),
		edge.To("entitytypes", EntityType.Type).Comment("The entity types in the collection."),
		edge.To("created_by", User.Type).Unique().Comment("The user who created the collection."),
		edge.To("updated_by", User.Type).Unique().Comment("The user who updated the collection."),
		edge.To("deleted_by", User.Type).Unique().Comment("The user who deleted the collection."),
	}
}
