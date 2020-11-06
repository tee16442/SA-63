package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Adult schema.
type Adult struct {
	ent.Schema
}

// Fields of the Adult.
func (Adult) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Amount").Unique(),
	}
}

// Edges of the Adult.
func (Adult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Books.Type),
		
	}
}