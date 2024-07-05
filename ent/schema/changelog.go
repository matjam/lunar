package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Changelog holds the schema definition for the Changelog entity.
type Changelog struct {
	ent.Schema
}

// Fields of the Changelog.
func (Changelog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("The ID of the changelog entry."),
		field.Enum("type").Values("create", "update", "delete").Comment("The type of change that was made."),
		field.Enum("object").Values("user", "role", "permission", "namespace", "entity", "entityschema", "entitytype", "key").Comment("The type of object that was changed."),
		field.UUID("object_id", uuid.UUID{}).Comment("The unique ID of the object that was changed."),
		field.JSON("data", map[string]interface{}{}).Comment("The object data before the change."),
		field.JSON("data_diff", map[string]interface{}{}).Comment("A diff of the object data before and after the change."),
		field.String("ip_address").Comment("The IP address of the user who made the change."),
		field.String("user_agent").Comment("The user agent of the user who made the change."),
		field.String("note").Optional().Comment("A note about the change."),
		field.Bool("is_revert").Default(false).Comment("Whether the change was a revert."),
		field.UUID("reverted_from", uuid.UUID{}).Optional().Comment("The ID of the changelog entry that was reverted."),
		field.Time("created_at").Immutable().Comment("The time the change was made."),
	}
}

// Edges of the Changelog.
func (Changelog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Comment("The user who made the change."),
	}
}
