package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Customertype holds the schema definition for the Customertype entity.
type Customertype struct {
	ent.Schema
}

// Fields of the user.
func (Customertype) Fields() []ent.Field {
	return []ent.Field{
		field.String("CUSTOMERTYPE").
			NotEmpty().
			Unique(),
	}
}
// Edges of the user.
func (Customertype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer3", Customer.Type).StorageKey(edge.Column("customertype")),
	}
}
