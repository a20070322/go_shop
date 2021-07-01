// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/a20070322/shop-go/ent/orderaddress"
)

// OrderAddress is the model entity for the OrderAddress schema.
type OrderAddress struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderAddress) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderaddress.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderAddress", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderAddress fields.
func (oa *OrderAddress) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oa.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this OrderAddress.
// Note that you need to call OrderAddress.Unwrap() before calling this method if this OrderAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (oa *OrderAddress) Update() *OrderAddressUpdateOne {
	return (&OrderAddressClient{config: oa.config}).UpdateOne(oa)
}

// Unwrap unwraps the OrderAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oa *OrderAddress) Unwrap() *OrderAddress {
	tx, ok := oa.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderAddress is not a transactional entity")
	}
	oa.config.driver = tx.drv
	return oa
}

// String implements the fmt.Stringer.
func (oa *OrderAddress) String() string {
	var builder strings.Builder
	builder.WriteString("OrderAddress(")
	builder.WriteString(fmt.Sprintf("id=%v", oa.ID))
	builder.WriteByte(')')
	return builder.String()
}

// OrderAddresses is a parsable slice of OrderAddress.
type OrderAddresses []*OrderAddress

func (oa OrderAddresses) config(cfg config) {
	for _i := range oa {
		oa[_i].config = cfg
	}
}