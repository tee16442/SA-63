// Code generated by entc, DO NOT EDIT.

package roomtype

const (
	// Label holds the string label denoting the roomtype type in the database.
	Label = "roomtype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldROOMPRICE holds the string denoting the roomprice field in the database.
	FieldROOMPRICE = "roomprice"
	// FieldTYPEDEATAIL holds the string denoting the typedeatail field in the database.
	FieldTYPEDEATAIL = "typedeatail"

	// EdgeRoom1 holds the string denoting the room1 edge name in mutations.
	EdgeRoom1 = "Room1"

	// Table holds the table name of the roomtype in the database.
	Table = "roomtypes"
	// Room1Table is the table the holds the Room1 relation/edge.
	Room1Table = "rooms"
	// Room1InverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	Room1InverseTable = "rooms"
	// Room1Column is the table column denoting the Room1 relation/edge.
	Room1Column = "roomtype_room1"
)

// Columns holds all SQL columns for roomtype fields.
var Columns = []string{
	FieldID,
	FieldROOMPRICE,
	FieldTYPEDEATAIL,
}

var (
	// TYPEDEATAILValidator is a validator for the "TYPEDEATAIL" field. It is called by the builders before save.
	TYPEDEATAILValidator func(string) error
)
