package schema

import (
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Employee schema.
type Employee struct {
    ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
    return []ent.Field{
        field.String("EMPLOYEENAME"),
        field.String("EMPLOYEEPASSWORD"),
	}	
}
// Edges of the Customer.
func (Employee) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("checkout3", Checkout.Type),
    }
}