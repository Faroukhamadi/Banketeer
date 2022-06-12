// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb/ent/account"
	"github.com/Faroukhamadi/Banketeer/customer_service/customerdb/ent/customer"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// Cin holds the value of the "cin" field.
	Cin string `json:"cin,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone int `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// The edge account was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID, customer.FieldPhone:
			values[i] = new(sql.NullInt64)
		case customer.FieldFirstName, customer.FieldLastName, customer.FieldCin:
			values[i] = new(sql.NullString)
		case customer.FieldCreateTime, customer.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case customer.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case customer.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstName", values[i])
			} else if value.Valid {
				c.FirstName = value.String
			}
		case customer.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastName", values[i])
			} else if value.Valid {
				c.LastName = value.String
			}
		case customer.FieldCin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cin", values[i])
			} else if value.Valid {
				c.Cin = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAccount queries the "account" edge of the Customer entity.
func (c *Customer) QueryAccount() *AccountQuery {
	return (&CustomerClient{config: c.config}).QueryAccount(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
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
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("firstName=")
	builder.WriteString(c.FirstName)
	builder.WriteString(", ")
	builder.WriteString("lastName=")
	builder.WriteString(c.LastName)
	builder.WriteString(", ")
	builder.WriteString("cin=")
	builder.WriteString(c.Cin)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(fmt.Sprintf("%v", c.Phone))
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
