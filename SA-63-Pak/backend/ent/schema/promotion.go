package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Promotion schema.
type Promotion struct {
   ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
   return []ent.Field{
       field.String("PROMOTIONDETAIL"),
   }
}
// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Room2", Room.Type),        
    }
}
