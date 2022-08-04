// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"ice-mall/ent/messageread"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// MessageRead is the model entity for the MessageRead schema.
type MessageRead struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageRead) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case messageread.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MessageRead", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageRead fields.
func (mr *MessageRead) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messageread.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mr.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this MessageRead.
// Note that you need to call MessageRead.Unwrap() before calling this method if this MessageRead
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MessageRead) Update() *MessageReadUpdateOne {
	return (&MessageReadClient{config: mr.config}).UpdateOne(mr)
}

// Unwrap unwraps the MessageRead entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MessageRead) Unwrap() *MessageRead {
	tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageRead is not a transactional entity")
	}
	mr.config.driver = tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MessageRead) String() string {
	var builder strings.Builder
	builder.WriteString("MessageRead(")
	builder.WriteString(fmt.Sprintf("id=%v", mr.ID))
	builder.WriteByte(')')
	return builder.String()
}

// MessageReads is a parsable slice of MessageRead.
type MessageReads []*MessageRead

func (mr MessageReads) config(cfg config) {
	for _i := range mr {
		mr[_i].config = cfg
	}
}