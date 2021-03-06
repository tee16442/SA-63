// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ADMIN/app/ent/customer"
	"github.com/ADMIN/app/ent/customertype"
	"github.com/ADMIN/app/ent/gender"
	"github.com/ADMIN/app/ent/title"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NAME holds the value of the "NAME" field.
	NAME string `json:"NAME,omitempty"`
	// AGE holds the value of the "AGE" field.
	AGE int `json:"AGE,omitempty"`
	// EMAIL holds the value of the "EMAIL" field.
	EMAIL string `json:"EMAIL,omitempty"`
	// USERNAME holds the value of the "USERNAME" field.
	USERNAME string `json:"USERNAME,omitempty"`
	// PASSWORD holds the value of the "PASSWORD" field.
	PASSWORD string `json:"PASSWORD,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges        CustomerEdges `json:"edges"`
	customertype *int
	gender       *int
	title        *int
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Customertype holds the value of the customertype edge.
	Customertype *Customertype
	// Title holds the value of the title edge.
	Title *Title
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// CustomertypeOrErr returns the Customertype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) CustomertypeOrErr() (*Customertype, error) {
	if e.loadedTypes[1] {
		if e.Customertype == nil {
			// The edge customertype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customertype.Label}
		}
		return e.Customertype, nil
	}
	return nil, &NotLoadedError{edge: "customertype"}
}

// TitleOrErr returns the Title value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) TitleOrErr() (*Title, error) {
	if e.loadedTypes[2] {
		if e.Title == nil {
			// The edge title was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: title.Label}
		}
		return e.Title, nil
	}
	return nil, &NotLoadedError{edge: "title"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NAME
		&sql.NullInt64{},  // AGE
		&sql.NullString{}, // EMAIL
		&sql.NullString{}, // USERNAME
		&sql.NullString{}, // PASSWORD
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Customer) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // customertype
		&sql.NullInt64{}, // gender
		&sql.NullInt64{}, // title
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
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field AGE", values[1])
	} else if value.Valid {
		c.AGE = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMAIL", values[2])
	} else if value.Valid {
		c.EMAIL = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USERNAME", values[3])
	} else if value.Valid {
		c.USERNAME = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PASSWORD", values[4])
	} else if value.Valid {
		c.PASSWORD = value.String
	}
	values = values[5:]
	if len(values) == len(customer.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customertype", value)
		} else if value.Valid {
			c.customertype = new(int)
			*c.customertype = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender", value)
		} else if value.Valid {
			c.gender = new(int)
			*c.gender = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field title", value)
		} else if value.Valid {
			c.title = new(int)
			*c.title = int(value.Int64)
		}
	}
	return nil
}

// QueryGender queries the gender edge of the Customer.
func (c *Customer) QueryGender() *GenderQuery {
	return (&CustomerClient{config: c.config}).QueryGender(c)
}

// QueryCustomertype queries the customertype edge of the Customer.
func (c *Customer) QueryCustomertype() *CustomertypeQuery {
	return (&CustomerClient{config: c.config}).QueryCustomertype(c)
}

// QueryTitle queries the title edge of the Customer.
func (c *Customer) QueryTitle() *TitleQuery {
	return (&CustomerClient{config: c.config}).QueryTitle(c)
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
	builder.WriteString(", AGE=")
	builder.WriteString(fmt.Sprintf("%v", c.AGE))
	builder.WriteString(", EMAIL=")
	builder.WriteString(c.EMAIL)
	builder.WriteString(", USERNAME=")
	builder.WriteString(c.USERNAME)
	builder.WriteString(", PASSWORD=")
	builder.WriteString(c.PASSWORD)
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
