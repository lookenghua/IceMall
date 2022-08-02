// Code generated by entc, DO NOT EDIT.

package productevaluation

const (
	// Label holds the string label denoting the productevaluation type in the database.
	Label = "product_evaluation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the productevaluation in the database.
	Table = "product_evaluations"
)

// Columns holds all SQL columns for productevaluation fields.
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
