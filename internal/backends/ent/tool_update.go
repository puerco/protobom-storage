// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/tool"
)

// ToolUpdate is the builder for updating Tool entities.
type ToolUpdate struct {
	config
	hooks    []Hook
	mutation *ToolMutation
}

// Where appends a list predicates to the ToolUpdate builder.
func (tu *ToolUpdate) Where(ps ...predicate.Tool) *ToolUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetMetadataID sets the "metadata_id" field.
func (tu *ToolUpdate) SetMetadataID(s string) *ToolUpdate {
	tu.mutation.SetMetadataID(s)
	return tu
}

// SetNillableMetadataID sets the "metadata_id" field if the given value is not nil.
func (tu *ToolUpdate) SetNillableMetadataID(s *string) *ToolUpdate {
	if s != nil {
		tu.SetMetadataID(*s)
	}
	return tu
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (tu *ToolUpdate) ClearMetadataID() *ToolUpdate {
	tu.mutation.ClearMetadataID()
	return tu
}

// SetName sets the "name" field.
func (tu *ToolUpdate) SetName(s string) *ToolUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *ToolUpdate) SetNillableName(s *string) *ToolUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetVersion sets the "version" field.
func (tu *ToolUpdate) SetVersion(s string) *ToolUpdate {
	tu.mutation.SetVersion(s)
	return tu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tu *ToolUpdate) SetNillableVersion(s *string) *ToolUpdate {
	if s != nil {
		tu.SetVersion(*s)
	}
	return tu
}

// SetVendor sets the "vendor" field.
func (tu *ToolUpdate) SetVendor(s string) *ToolUpdate {
	tu.mutation.SetVendor(s)
	return tu
}

// SetNillableVendor sets the "vendor" field if the given value is not nil.
func (tu *ToolUpdate) SetNillableVendor(s *string) *ToolUpdate {
	if s != nil {
		tu.SetVendor(*s)
	}
	return tu
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (tu *ToolUpdate) SetMetadata(m *Metadata) *ToolUpdate {
	return tu.SetMetadataID(m.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tu *ToolUpdate) Mutation() *ToolMutation {
	return tu.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (tu *ToolUpdate) ClearMetadata() *ToolUpdate {
	tu.mutation.ClearMetadata()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ToolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ToolUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ToolUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ToolUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *ToolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tool.Table, tool.Columns, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tool.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Version(); ok {
		_spec.SetField(tool.FieldVersion, field.TypeString, value)
	}
	if value, ok := tu.mutation.Vendor(); ok {
		_spec.SetField(tool.FieldVendor, field.TypeString, value)
	}
	if tu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tool.MetadataTable,
			Columns: []string{tool.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tool.MetadataTable,
			Columns: []string{tool.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// ToolUpdateOne is the builder for updating a single Tool entity.
type ToolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ToolMutation
}

// SetMetadataID sets the "metadata_id" field.
func (tuo *ToolUpdateOne) SetMetadataID(s string) *ToolUpdateOne {
	tuo.mutation.SetMetadataID(s)
	return tuo
}

// SetNillableMetadataID sets the "metadata_id" field if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableMetadataID(s *string) *ToolUpdateOne {
	if s != nil {
		tuo.SetMetadataID(*s)
	}
	return tuo
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (tuo *ToolUpdateOne) ClearMetadataID() *ToolUpdateOne {
	tuo.mutation.ClearMetadataID()
	return tuo
}

// SetName sets the "name" field.
func (tuo *ToolUpdateOne) SetName(s string) *ToolUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableName(s *string) *ToolUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetVersion sets the "version" field.
func (tuo *ToolUpdateOne) SetVersion(s string) *ToolUpdateOne {
	tuo.mutation.SetVersion(s)
	return tuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableVersion(s *string) *ToolUpdateOne {
	if s != nil {
		tuo.SetVersion(*s)
	}
	return tuo
}

// SetVendor sets the "vendor" field.
func (tuo *ToolUpdateOne) SetVendor(s string) *ToolUpdateOne {
	tuo.mutation.SetVendor(s)
	return tuo
}

// SetNillableVendor sets the "vendor" field if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableVendor(s *string) *ToolUpdateOne {
	if s != nil {
		tuo.SetVendor(*s)
	}
	return tuo
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (tuo *ToolUpdateOne) SetMetadata(m *Metadata) *ToolUpdateOne {
	return tuo.SetMetadataID(m.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tuo *ToolUpdateOne) Mutation() *ToolMutation {
	return tuo.mutation
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (tuo *ToolUpdateOne) ClearMetadata() *ToolUpdateOne {
	tuo.mutation.ClearMetadata()
	return tuo
}

// Where appends a list predicates to the ToolUpdate builder.
func (tuo *ToolUpdateOne) Where(ps ...predicate.Tool) *ToolUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ToolUpdateOne) Select(field string, fields ...string) *ToolUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tool entity.
func (tuo *ToolUpdateOne) Save(ctx context.Context) (*Tool, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ToolUpdateOne) SaveX(ctx context.Context) *Tool {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ToolUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ToolUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *ToolUpdateOne) sqlSave(ctx context.Context) (_node *Tool, err error) {
	_spec := sqlgraph.NewUpdateSpec(tool.Table, tool.Columns, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tool.FieldID)
		for _, f := range fields {
			if !tool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tool.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Version(); ok {
		_spec.SetField(tool.FieldVersion, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Vendor(); ok {
		_spec.SetField(tool.FieldVendor, field.TypeString, value)
	}
	if tuo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tool.MetadataTable,
			Columns: []string{tool.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tool.MetadataTable,
			Columns: []string{tool.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tool{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
