// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/disism/karma/ent/dirs"
	"github.com/disism/karma/ent/users"
)

// Dirs is the model entity for the Dirs schema.
type Dirs struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DirsQuery when eager-loading is set.
	Edges        DirsEdges `json:"edges"`
	users_dirs   *uint64
	selectValues sql.SelectValues
}

// DirsEdges holds the relations/edges for other nodes in the graph.
type DirsEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Users `json:"owner,omitempty"`
	// Saves holds the value of the saves edge.
	Saves []*Saves `json:"saves,omitempty"`
	// Subdir holds the value of the subdir edge.
	Subdir []*Dirs `json:"subdir,omitempty"`
	// Pdir holds the value of the pdir edge.
	Pdir []*Dirs `json:"pdir,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DirsEdges) OwnerOrErr() (*Users, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: users.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// SavesOrErr returns the Saves value or an error if the edge
// was not loaded in eager-loading.
func (e DirsEdges) SavesOrErr() ([]*Saves, error) {
	if e.loadedTypes[1] {
		return e.Saves, nil
	}
	return nil, &NotLoadedError{edge: "saves"}
}

// SubdirOrErr returns the Subdir value or an error if the edge
// was not loaded in eager-loading.
func (e DirsEdges) SubdirOrErr() ([]*Dirs, error) {
	if e.loadedTypes[2] {
		return e.Subdir, nil
	}
	return nil, &NotLoadedError{edge: "subdir"}
}

// PdirOrErr returns the Pdir value or an error if the edge
// was not loaded in eager-loading.
func (e DirsEdges) PdirOrErr() ([]*Dirs, error) {
	if e.loadedTypes[3] {
		return e.Pdir, nil
	}
	return nil, &NotLoadedError{edge: "pdir"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dirs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dirs.FieldID:
			values[i] = new(sql.NullInt64)
		case dirs.FieldName:
			values[i] = new(sql.NullString)
		case dirs.FieldCreateTime, dirs.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case dirs.ForeignKeys[0]: // users_dirs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dirs fields.
func (d *Dirs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dirs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = uint64(value.Int64)
		case dirs.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case dirs.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = value.Time
			}
		case dirs.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case dirs.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field users_dirs", value)
			} else if value.Valid {
				d.users_dirs = new(uint64)
				*d.users_dirs = uint64(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dirs.
// This includes values selected through modifiers, order, etc.
func (d *Dirs) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Dirs entity.
func (d *Dirs) QueryOwner() *UsersQuery {
	return NewDirsClient(d.config).QueryOwner(d)
}

// QuerySaves queries the "saves" edge of the Dirs entity.
func (d *Dirs) QuerySaves() *SavesQuery {
	return NewDirsClient(d.config).QuerySaves(d)
}

// QuerySubdir queries the "subdir" edge of the Dirs entity.
func (d *Dirs) QuerySubdir() *DirsQuery {
	return NewDirsClient(d.config).QuerySubdir(d)
}

// QueryPdir queries the "pdir" edge of the Dirs entity.
func (d *Dirs) QueryPdir() *DirsQuery {
	return NewDirsClient(d.config).QueryPdir(d)
}

// Update returns a builder for updating this Dirs.
// Note that you need to call Dirs.Unwrap() before calling this method if this Dirs
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dirs) Update() *DirsUpdateOne {
	return NewDirsClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dirs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dirs) Unwrap() *Dirs {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dirs is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dirs) String() string {
	var builder strings.Builder
	builder.WriteString("Dirs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(d.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteByte(')')
	return builder.String()
}

// DirsSlice is a parsable slice of Dirs.
type DirsSlice []*Dirs
