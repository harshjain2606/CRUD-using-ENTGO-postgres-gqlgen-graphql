package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Client holds the schema definition for the Client entity.
type Customer struct {
	ent.Schema
}

// Fields of the Client.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("password"),
	}
}
