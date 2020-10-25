package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME"),
	}
}

// Edges of the customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Problem", Problem.Type),
	}
}
