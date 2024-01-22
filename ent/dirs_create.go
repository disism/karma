// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/disism/karma/ent/dirs"
	"github.com/disism/karma/ent/saves"
	"github.com/disism/karma/ent/users"
)

// DirsCreate is the builder for creating a Dirs entity.
type DirsCreate struct {
	config
	mutation *DirsMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DirsCreate) SetCreateTime(t time.Time) *DirsCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DirsCreate) SetNillableCreateTime(t *time.Time) *DirsCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DirsCreate) SetUpdateTime(t time.Time) *DirsCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DirsCreate) SetNillableUpdateTime(t *time.Time) *DirsCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DirsCreate) SetName(s string) *DirsCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DirsCreate) SetID(u uint64) *DirsCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetOwnerID sets the "owner" edge to the Users entity by ID.
func (dc *DirsCreate) SetOwnerID(id uint64) *DirsCreate {
	dc.mutation.SetOwnerID(id)
	return dc
}

// SetNillableOwnerID sets the "owner" edge to the Users entity by ID if the given value is not nil.
func (dc *DirsCreate) SetNillableOwnerID(id *uint64) *DirsCreate {
	if id != nil {
		dc = dc.SetOwnerID(*id)
	}
	return dc
}

// SetOwner sets the "owner" edge to the Users entity.
func (dc *DirsCreate) SetOwner(u *Users) *DirsCreate {
	return dc.SetOwnerID(u.ID)
}

// AddSafeIDs adds the "saves" edge to the Saves entity by IDs.
func (dc *DirsCreate) AddSafeIDs(ids ...uint64) *DirsCreate {
	dc.mutation.AddSafeIDs(ids...)
	return dc
}

// AddSaves adds the "saves" edges to the Saves entity.
func (dc *DirsCreate) AddSaves(s ...*Saves) *DirsCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSafeIDs(ids...)
}

// AddSubdirIDs adds the "subdir" edge to the Dirs entity by IDs.
func (dc *DirsCreate) AddSubdirIDs(ids ...uint64) *DirsCreate {
	dc.mutation.AddSubdirIDs(ids...)
	return dc
}

// AddSubdir adds the "subdir" edges to the Dirs entity.
func (dc *DirsCreate) AddSubdir(d ...*Dirs) *DirsCreate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddSubdirIDs(ids...)
}

// AddPdirIDs adds the "pdir" edge to the Dirs entity by IDs.
func (dc *DirsCreate) AddPdirIDs(ids ...uint64) *DirsCreate {
	dc.mutation.AddPdirIDs(ids...)
	return dc
}

// AddPdir adds the "pdir" edges to the Dirs entity.
func (dc *DirsCreate) AddPdir(d ...*Dirs) *DirsCreate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddPdirIDs(ids...)
}

// Mutation returns the DirsMutation object of the builder.
func (dc *DirsCreate) Mutation() *DirsMutation {
	return dc.mutation
}

// Save creates the Dirs in the database.
func (dc *DirsCreate) Save(ctx context.Context) (*Dirs, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DirsCreate) SaveX(ctx context.Context) *Dirs {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DirsCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DirsCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DirsCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := dirs.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := dirs.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DirsCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Dirs.create_time"`)}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Dirs.update_time"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Dirs.name"`)}
	}
	return nil
}

func (dc *DirsCreate) sqlSave(ctx context.Context) (*Dirs, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DirsCreate) createSpec() (*Dirs, *sqlgraph.CreateSpec) {
	var (
		_node = &Dirs{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dirs.Table, sqlgraph.NewFieldSpec(dirs.FieldID, field.TypeUint64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.SetField(dirs.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.SetField(dirs.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(dirs.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dirs.OwnerTable,
			Columns: []string{dirs.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.users_dirs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SavesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dirs.SavesTable,
			Columns: dirs.SavesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(saves.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SubdirIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   dirs.SubdirTable,
			Columns: dirs.SubdirPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dirs.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PdirIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dirs.PdirTable,
			Columns: dirs.PdirPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dirs.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DirsCreateBulk is the builder for creating many Dirs entities in bulk.
type DirsCreateBulk struct {
	config
	err      error
	builders []*DirsCreate
}

// Save creates the Dirs entities in the database.
func (dcb *DirsCreateBulk) Save(ctx context.Context) ([]*Dirs, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dirs, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DirsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DirsCreateBulk) SaveX(ctx context.Context) []*Dirs {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DirsCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DirsCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
