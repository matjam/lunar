package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity. Sessions are used to store user data between requests.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", nil).Unique().Immutable().Comment("The ID of the session."),
		field.JSON("data", map[string]any{}).Default(map[string]any{}).
			Comment("The data stored in the session."),
		field.Time("created_at").Immutable().Comment("The time the session was created."),
		field.Time("updated_at").Nillable().Optional().Comment("The time the session was last updated."),
		field.Time("deleted_at").Nillable().Optional().Comment("The time the session was deleted."),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Unique().Comment("The user that owns the session."),
	}
}
