package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)

// Customer schema.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME"),
		field.Int("AGE"),
		field.String("EMAIL").
			NotEmpty().
			Unique(),
		field.String("USERNAME").
			NotEmpty().
			Unique(),
		field.String("PASSWORD").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
    return []ent.Edge{
		edge.From("gender", Gender.Type).
			Ref("customer2").
			Unique(),
		edge.From("customertype", Customertype.Type).
			Ref("customer3").
			Unique(),
		edge.From("title", Title.Type).
			Ref("customer4").
			Unique(),
    }
}