package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Roomtype schema.
type Roomtype struct {
	ent.Schema
}

// Fields of the Roomtype.
func (Roomtype) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ROOMPRICE"),
	}
}

// Edges of the Roomtype.
func (Roomtype) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("payment", Payment.Type),
    }
}