package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Pet.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.String("GENDER").
			NotEmpty().
			Unique(),
	}
}
// Edges of the user.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer2", Customer.Type).StorageKey(edge.Column("gender")),
			
	}
}
