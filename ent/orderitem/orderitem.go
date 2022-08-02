// Code generated by entc, DO NOT EDIT.

package orderitem

const (
	// Label holds the string label denoting the orderitem type in the database.
	Label = "order_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the orderitem in the database.
	Table = "order_items"
)

// Columns holds all SQL columns for orderitem fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
