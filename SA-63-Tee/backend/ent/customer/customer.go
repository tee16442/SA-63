// Code generated by entc, DO NOT EDIT.

package customer

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNAME holds the string denoting the name field in the database.
	FieldNAME = "name"

	// EdgeProblem holds the string denoting the problem edge name in mutations.
	EdgeProblem = "Problem"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// ProblemTable is the table the holds the Problem relation/edge.
	ProblemTable = "problems"
	// ProblemInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	ProblemInverseTable = "problems"
	// ProblemColumn is the table column denoting the Problem relation/edge.
	ProblemColumn = "customer_problem"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldNAME,
}
