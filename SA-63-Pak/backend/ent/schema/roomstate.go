package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// RoomState schema.
type Roomstate struct {
   ent.Schema
}

// Fields of the RoomState.
func (Roomstate) Fields() []ent.Field {
   return []ent.Field{
       field.String("ROOMSTATE").
         NotEmpty(),
   }
}
// Edges of the RoomState.
func (Roomstate) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Room3", Room.Type),
    }
}