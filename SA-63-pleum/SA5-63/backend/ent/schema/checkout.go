package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	
)

// Checkout schema.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return []ent.Field{
		field.Time("CHECKOUTDATE"),
		
	}
}

// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("books", Books.Type).Ref("checkout1").Unique(),
		edge.From("customer", Customer.Type).Ref("checkout2").Unique(),
		edge.From("employee", Employee.Type).Ref("checkout3").Unique(),
		edge.From("room", Room.Type).Ref("checkout4").Unique(),
	}
}
