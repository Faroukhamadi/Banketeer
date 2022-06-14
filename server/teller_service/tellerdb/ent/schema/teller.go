package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Teller holds the schema definition for the Teller entity.
type Teller struct {
	ent.Schema
}

func (Teller) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Teller.
func (Teller) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.String("password"),
		field.Enum("role").
			Values("customer", "admin").
			Default("customer"),
	}
}

// Edges of the Teller.
func (Teller) Edges() []ent.Edge {
	return nil
}
