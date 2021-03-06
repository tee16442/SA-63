// Code generated by entc, DO NOT EDIT.

package roomtype

const (
	// Label holds the string label denoting the roomtype type in the database.
	Label = "roomtype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldROOMPRICE holds the string denoting the roomprice field in the database.
	FieldROOMPRICE = "roomprice"

	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"

	// Table holds the table name of the roomtype in the database.
	Table = "roomtypes"
	// PaymentTable is the table the holds the payment relation/edge.
	PaymentTable = "payments"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payments"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "roomtype_payment"
)

// Columns holds all SQL columns for roomtype fields.
var Columns = []string{
	FieldID,
	FieldROOMPRICE,
}
