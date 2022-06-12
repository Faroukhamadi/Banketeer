package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstName"),
		field.String("lastName"),
		field.String("cin").Unique(),
		field.Int("phone").Unique(),
	}
}

func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Unique(),
	}
}
