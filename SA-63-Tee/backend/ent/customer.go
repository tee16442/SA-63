// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/tee16/app/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NAME holds the value of the "NAME" field.
	NAME string `json:"NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Problem holds the value of the Problem edge.
	Problem []*Problem
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProblemOrErr returns the Problem value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) ProblemOrErr() ([]*Problem, error) {
	if e.loadedTypes[0] {
		return e.Problem, nil
	}
	return nil, &NotLoadedError{edge: "Problem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field NAME", values[0])
	} else if value.Valid {
		c.NAME = value.String
	}
	return nil
}

// QueryProblem queries the Problem edge of the Customer.
func (c *Customer) QueryProblem() *ProblemQuery {
	return (&CustomerClient{config: c.config}).QueryProblem(c)
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", NAME=")
	builder.WriteString(c.NAME)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
