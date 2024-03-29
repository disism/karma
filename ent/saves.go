// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/disism/karma/ent/files"
	"github.com/disism/karma/ent/saves"
	"github.com/disism/karma/ent/users"
)

// Saves is the model entity for the Saves schema.
type Saves struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// the descriptive text or title of a document, image, or other media element. It is used to provide a short description of the content, characteristics or context of a document.
	Caption string `json:"caption,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SavesQuery when eager-loading is set.
	Edges        SavesEdges `json:"edges"`
	files_saves  *uint64
	users_saves  *uint64
	selectValues sql.SelectValues
}

// SavesEdges holds the relations/edges for other nodes in the graph.
type SavesEdges struct {
	// File holds the value of the file edge.
	File *Files `json:"file,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Users `json:"owner,omitempty"`
	// Dir holds the value of the dir edge.
	Dir []*Dirs `json:"dir,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FileOrErr returns the File value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SavesEdges) FileOrErr() (*Files, error) {
	if e.loadedTypes[0] {
		if e.File == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: files.Label}
		}
		return e.File, nil
	}
	return nil, &NotLoadedError{edge: "file"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SavesEdges) OwnerOrErr() (*Users, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: users.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// DirOrErr returns the Dir value or an error if the edge
// was not loaded in eager-loading.
func (e SavesEdges) DirOrErr() ([]*Dirs, error) {
	if e.loadedTypes[2] {
		return e.Dir, nil
	}
	return nil, &NotLoadedError{edge: "dir"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Saves) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case saves.FieldID:
			values[i] = new(sql.NullInt64)
		case saves.FieldName, saves.FieldCaption:
			values[i] = new(sql.NullString)
		case saves.FieldCreateTime, saves.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case saves.ForeignKeys[0]: // files_saves
			values[i] = new(sql.NullInt64)
		case saves.ForeignKeys[1]: // users_saves
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Saves fields.
func (s *Saves) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case saves.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case saves.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case saves.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case saves.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case saves.FieldCaption:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field caption", values[i])
			} else if value.Valid {
				s.Caption = value.String
			}
		case saves.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field files_saves", value)
			} else if value.Valid {
				s.files_saves = new(uint64)
				*s.files_saves = uint64(value.Int64)
			}
		case saves.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field users_saves", value)
			} else if value.Valid {
				s.users_saves = new(uint64)
				*s.users_saves = uint64(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Saves.
// This includes values selected through modifiers, order, etc.
func (s *Saves) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryFile queries the "file" edge of the Saves entity.
func (s *Saves) QueryFile() *FilesQuery {
	return NewSavesClient(s.config).QueryFile(s)
}

// QueryOwner queries the "owner" edge of the Saves entity.
func (s *Saves) QueryOwner() *UsersQuery {
	return NewSavesClient(s.config).QueryOwner(s)
}

// QueryDir queries the "dir" edge of the Saves entity.
func (s *Saves) QueryDir() *DirsQuery {
	return NewSavesClient(s.config).QueryDir(s)
}

// Update returns a builder for updating this Saves.
// Note that you need to call Saves.Unwrap() before calling this method if this Saves
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Saves) Update() *SavesUpdateOne {
	return NewSavesClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Saves entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Saves) Unwrap() *Saves {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Saves is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Saves) String() string {
	var builder strings.Builder
	builder.WriteString("Saves(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("caption=")
	builder.WriteString(s.Caption)
	builder.WriteByte(')')
	return builder.String()
}

// SavesSlice is a parsable slice of Saves.
type SavesSlice []*Saves
