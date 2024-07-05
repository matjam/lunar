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
		field.UUID("created_by", uuid.UUID{}).Optional().Comment("The user who created the entity."),
		field.UUID("updated_by", uuid.UUID{}).Optional().Comment("The user who updated the entity."),
		field.UUID("deleted_by", uuid.UUID{}).Optional().Comment("The user who deleted the entity."),
	}
}

// Edges of the Entity.
func (Entity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("entitytype", EntityType.Type).Unique().Comment("The type of the entity."),
		edge.To("entityschema", EntitySchema.Type).Unique().Comment("The schema of the entity."),
		edge.To("namespace", Namespace.Type).Unique().Comment("The namespace of the entity."),
		edge.To("created_by", User.Type).Unique().Comment("The user who created the entity."),
		edge.To("updated_by", User.Type).Unique().Comment("The user who updated the entity."),
		edge.To("deleted_by", User.Type).Unique().Comment("The user who deleted the entity."),
		edge.To("write_permissions", Permission.Type).Comment("The write permissions of the entity."),
		edge.To("read_permissions", Permission.Type).Comment("The read permissions of the entity."),
		edge.From("entityschema_versions", EntitySchemaVersion.Type).Ref("entities").Comment("The versions of the entityschema."),
	}
}
