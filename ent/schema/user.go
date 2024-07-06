package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Nillable().Optional(),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("created_by", User.Type).Unique(),
		edge.To("updated_by", User.Type).Unique(),
		edge.To("deleted_by", User.Type).Unique(),
		edge.To("roles", Role.Type),
		edge.To("sessions", Session.Type),
	}
}
