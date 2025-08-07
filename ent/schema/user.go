package schema

import (
	

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.

// Fields of the User.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("password").Default(""),

	}
}

