package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payment schema.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("PAYDAY"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("payment").
			Unique(),
		edge.From("employee", Employee.Type).
			Ref("payment").
			Unique(),
		edge.From("paymenttype", Paymenttype.Type).
			Ref("payment").
			Unique(),
		edge.From("roomtype", Roomtype.Type).
			Ref("payment").
			Unique(),
	}
}
