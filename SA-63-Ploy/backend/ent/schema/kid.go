package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Kid schema.
type Kid struct {
	ent.Schema
}

// Fields of the Kid.
func (Kid) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Amount").Unique(),
	}
}

// Edges of the Kid.
func (Kid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Books.Type),
		
	}
}