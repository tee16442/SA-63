package schema

import (
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Books schema.
type Books struct {
    ent.Schema
}

// Fields of the Books.
func (Books) Fields() []ent.Field {
    return []ent.Field{
		field.String("CHECKINDATE"),
        
    }
}

// Edges of the Books.
func (Books) Edges() []ent.Edge {
    return []ent.Edge{
         edge.To("checkout1", Checkout.Type).Unique(),
    }
}