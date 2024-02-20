package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserDetails holds the schema definition for the UserDetails entity.
type UserDetails struct {
	ent.Schema
}

// Fields of the UserDetails.
func (UserDetails) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("email").Optional(),
	}
}

// Edges of the UserDetails.
func (UserDetails) Edges() []ent.Edge {
	return nil
}
