// Code generated by entc, DO NOT EDIT.

package productattributevalue

const (
	// Label holds the string label denoting the productattributevalue type in the database.
	Label = "product_attribute_value"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the productattributevalue in the database.
	Table = "product_attribute_values"
)

// Columns holds all SQL columns for productattributevalue fields.
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
