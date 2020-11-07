package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Customer schema.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		 field.String("NAME"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("checkout2", Checkout.Type),
    }
}
