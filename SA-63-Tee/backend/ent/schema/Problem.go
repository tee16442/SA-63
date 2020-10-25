package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Problem  holds the schema definition for the Problem  entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem .
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.String("PROBLEMDETAIL"),
	}
}

// Edges of the Problem .
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Problemtype", Problemtype.Type).
			Ref("Problem").
			Unique(),
		edge.From("Room", Room.Type).
			Ref("Problem").
			Unique(),
		edge.From("Customer", Customer.Type).
			Ref("Problem").
			Unique(),
	}
}
