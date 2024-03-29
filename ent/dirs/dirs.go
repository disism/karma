// Code generated by ent, DO NOT EDIT.

package dirs

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dirs type in the database.
	Label = "dirs"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeSaves holds the string denoting the saves edge name in mutations.
	EdgeSaves = "saves"
	// EdgeSubdir holds the string denoting the subdir edge name in mutations.
	EdgeSubdir = "subdir"
	// EdgePdir holds the string denoting the pdir edge name in mutations.
	EdgePdir = "pdir"
	// Table holds the table name of the dirs in the database.
	Table = "dirs"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "dirs"
	// OwnerInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "users_dirs"
	// SavesTable is the table that holds the saves relation/edge. The primary key declared below.
	SavesTable = "dirs_saves"
	// SavesInverseTable is the table name for the Saves entity.
	// It exists in this package in order to avoid circular dependency with the "saves" package.
	SavesInverseTable = "saves"
	// SubdirTable is the table that holds the subdir relation/edge. The primary key declared below.
	SubdirTable = "dirs_subdir"
	// PdirTable is the table that holds the pdir relation/edge. The primary key declared below.
	PdirTable = "dirs_subdir"
)

// Columns holds all SQL columns for dirs fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "dirs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"users_dirs",
}

var (
	// SavesPrimaryKey and SavesColumn2 are the table columns denoting the
	// primary key for the saves relation (M2M).
	SavesPrimaryKey = []string{"dirs_id", "saves_id"}
	// SubdirPrimaryKey and SubdirColumn2 are the table columns denoting the
	// primary key for the subdir relation (M2M).
	SubdirPrimaryKey = []string{"dirs_id", "pdir_id"}
	// PdirPrimaryKey and PdirColumn2 are the table columns denoting the
	// primary key for the pdir relation (M2M).
	PdirPrimaryKey = []string{"dirs_id", "pdir_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
)

// OrderOption defines the ordering options for the Dirs queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
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

// BySubdirCount orders the results by subdir count.
func BySubdirCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubdirStep(), opts...)
	}
}

// BySubdir orders the results by subdir terms.
func BySubdir(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubdirStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPdirCount orders the results by pdir count.
func ByPdirCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPdirStep(), opts...)
	}
}

// ByPdir orders the results by pdir terms.
func ByPdir(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPdirStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newSavesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SavesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SavesTable, SavesPrimaryKey...),
	)
}
func newSubdirStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubdirTable, SubdirPrimaryKey...),
	)
}
func newPdirStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PdirTable, PdirPrimaryKey...),
	)
}
