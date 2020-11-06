package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// RoomType schema.
type Roomtype struct {
   ent.Schema
}

// Fields of the RoomType.
func (Roomtype) Fields() []ent.Field {
   return []ent.Field{
       field.Int("ROOMPRICE"),
	   field.String("TYPEDEATAIL").
	   	 NotEmpty(),
   }
}
// Edges of the RoomType.
func (Roomtype) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Room1", Room.Type),
    }
}
