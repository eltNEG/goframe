package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permissions holds the schema definition for the Permissions entity.
type Permissions struct {
	ent.Schema
}

// Fields of the Permissions.
func (Permissions) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Permissions.
func (Permissions) Edges() []ent.Edge {
	return nil
}
