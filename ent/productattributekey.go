// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"ice-mall/ent/productattributekey"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ProductAttributeKey is the model entity for the ProductAttributeKey schema.
type ProductAttributeKey struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductAttributeKey) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productattributekey.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductAttributeKey", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductAttributeKey fields.
func (pak *ProductAttributeKey) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productattributekey.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pak.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this ProductAttributeKey.
// Note that you need to call ProductAttributeKey.Unwrap() before calling this method if this ProductAttributeKey
// was returned from a transaction, and the transaction was committed or rolled back.
func (pak *ProductAttributeKey) Update() *ProductAttributeKeyUpdateOne {
	return (&ProductAttributeKeyClient{config: pak.config}).UpdateOne(pak)
}

// Unwrap unwraps the ProductAttributeKey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pak *ProductAttributeKey) Unwrap() *ProductAttributeKey {
	tx, ok := pak.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductAttributeKey is not a transactional entity")
	}
	pak.config.driver = tx.drv
	return pak
}

// String implements the fmt.Stringer.
func (pak *ProductAttributeKey) String() string {
	var builder strings.Builder
	builder.WriteString("ProductAttributeKey(")
	builder.WriteString(fmt.Sprintf("id=%v", pak.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ProductAttributeKeys is a parsable slice of ProductAttributeKey.
type ProductAttributeKeys []*ProductAttributeKey

func (pak ProductAttributeKeys) config(cfg config) {
	for _i := range pak {
		pak[_i].config = cfg
	}
}
