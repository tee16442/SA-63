package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Title holds the schema definition for the User entity.
type Title struct {
	ent.Schema
}

// Fields of the user.
func (Title) Fields() []ent.Field {
	return []ent.Field{
		field.String("TITLETYPE").
			NotEmpty().
			Unique(),
	}
}
// Edges of the user.
func (Title) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer4", Customer.Type).StorageKey(edge.Column("title")),
	}
}
