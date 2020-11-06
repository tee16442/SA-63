package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roomamount schema.
type Roomamount struct {
	ent.Schema
}

// Fields of the Roomamount.
func (Roomamount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Amount").Unique(),
	}
}

// Edges of the Roomamount.
func (Roomamount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Books.Type),
		
	}
}