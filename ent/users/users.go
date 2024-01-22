// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the users type in the database.
	Label = "users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeDirs holds the string denoting the dirs edge name in mutations.
	EdgeDirs = "dirs"
	// EdgeSaves holds the string denoting the saves edge name in mutations.
	EdgeSaves = "saves"
	// Table holds the table name of the users in the database.
	Table = "users"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "devices"
	// DevicesInverseTable is the table name for the Devices entity.
	// It exists in this package in order to avoid circular dependency with the "devices" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "users_devices"
	// DirsTable is the table that holds the dirs relation/edge.
	DirsTable = "dirs"
	// DirsInverseTable is the table name for the Dirs entity.
	// It exists in this package in order to avoid circular dependency with the "dirs" package.
	DirsInverseTable = "dirs"
	// DirsColumn is the table column denoting the dirs relation/edge.
	DirsColumn = "users_dirs"
	// SavesTable is the table that holds the saves relation/edge.
	SavesTable = "saves"
	// SavesInverseTable is the table name for the Saves entity.
	// It exists in this package in order to avoid circular dependency with the "saves" package.
	SavesInverseTable = "saves"
	// SavesColumn is the table column denoting the saves relation/edge.
	SavesColumn = "users_saves"
)

// Columns holds all SQL columns for users fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUsername,
	FieldPassword,
	FieldEmail,
	FieldName,
	FieldBio,
	FieldAvatar,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
)

// OrderOption defines the ordering options for the Users queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBio orders the results by the bio field.
func ByBio(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBio, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByDevicesCount orders the results by devices count.
func ByDevicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDevicesStep(), opts...)
	}
}

// ByDevices orders the results by devices terms.
func ByDevices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDevicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDirsCount orders the results by dirs count.
func ByDirsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDirsStep(), opts...)
	}
}

// ByDirs orders the results by dirs terms.
func ByDirs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDirsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySavesCount orders the results by saves count.
func BySavesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSavesStep(), opts...)
	}
}

// BySaves orders the results by saves terms.
func BySaves(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSavesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDevicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DevicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
	)
}
func newDirsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DirsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DirsTable, DirsColumn),
	)
}
func newSavesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SavesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SavesTable, SavesColumn),
	)
}