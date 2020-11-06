package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Books schema.
type Books struct {
	ent.Schema
}

// Fields of the Books.
func (Books) Fields() []ent.Field {
	return []ent.Field{
		field.Time("Checkin"),
		field.Time("Checkout"),
		
	}
}

// Edges of the Books.
func (Books) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("books").Unique(),
		edge.From("adult", Adult.Type).Ref("books").Unique(),
		edge.From("kid", Kid.Type).Ref("books").Unique(),
		edge.From("roomamount", Roomamount.Type).Ref("books").Unique(),
		edge.From("room", Room.Type).Ref("books").Unique(),
		
	}
}