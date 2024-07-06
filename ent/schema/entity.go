package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Entity holds the schema definition for the Entity entity.
type Entity struct {
	ent.Schema
}

// Fields of the Entity.
func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("The ID of the entity."),
		field.JSON("data", map[string]any{}).Default(map[string]any{}).Comment("The data of the entity."),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the Entity.
func (Entity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("collection", Collection.Type).Ref("entities").Comment("The collection that the entity belongs to."),
		edge.From("entitytype", EntityType.Type).Ref("entities").Comment("The entity type of the entity."),
		edge.To("created_by", User.Type).Unique().Comment("The user who created the entity."),
		edge.To("updated_by", User.Type).Unique().Comment("The user who updated the entity."),
		edge.To("deleted_by", User.Type).Unique().Comment("The user who deleted the entity."),
		edge.To("write_permissions", Permission.Type).Comment("The write permissions of the entity."),
		edge.To("read_permissions", Permission.Type).Comment("The read permissions of the entity."),
		edge.From("entityschemaversion", EntitySchemaVersion.Type).Ref("entities").Comment("The versions of the entityschema."),
	}
}
