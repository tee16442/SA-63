package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Problemtype holds the schema definition for the Problemtype entity.
type Problemtype struct {
	ent.Schema
}

// Fields of the Problemtype.
func (Problemtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("PROBLEMTYPE"),
			
	}
}

// Edges of the Problemtype.
func (Problemtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Problem", Problem.Type),
	}
}
