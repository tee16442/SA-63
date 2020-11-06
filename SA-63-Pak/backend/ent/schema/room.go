package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// Room schema.
type Room struct {
   ent.Schema
}


// Fields of the Room.
func (Room) Fields() []ent.Field {
   return []ent.Field{
      field.String("ROOMNUMBER").
      Unique(),
   }
}
// Edges of the Room.
func (Room) Edges() []ent.Edge {
   return []ent.Edge{
       edge.From("Roomtype", Roomtype.Type).
           Ref("Room1").
           Unique(),
       edge.From("Promotion", Promotion.Type).
           Ref("Room2").
           Unique(),
       edge.From("roomstate", Roomstate.Type).
           Ref("Room3").
           Unique(),
   }  
}
