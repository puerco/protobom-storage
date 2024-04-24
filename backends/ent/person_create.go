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
	"github.com/protobom/storage/backends/ent/metadata"
	"github.com/protobom/storage/backends/ent/node"
	"github.com/protobom/storage/backends/ent/person"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PersonCreate) SetName(s string) *PersonCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIsOrg sets the "is_org" field.
func (pc *PersonCreate) SetIsOrg(b bool) *PersonCreate {
	pc.mutation.SetIsOrg(b)
	return pc
}

// SetEmail sets the "email" field.
func (pc *PersonCreate) SetEmail(s string) *PersonCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetURL sets the "url" field.
func (pc *PersonCreate) SetURL(s string) *PersonCreate {
	pc.mutation.SetURL(s)
	return pc
}

// SetPhone sets the "phone" field.
func (pc *PersonCreate) SetPhone(s string) *PersonCreate {
	pc.mutation.SetPhone(s)
	return pc
}

// SetContactOwnerID sets the "contact_owner" edge to the Person entity by ID.
func (pc *PersonCreate) SetContactOwnerID(id int) *PersonCreate {
	pc.mutation.SetContactOwnerID(id)
	return pc
}

// SetNillableContactOwnerID sets the "contact_owner" edge to the Person entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableContactOwnerID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetContactOwnerID(*id)
	}
	return pc
}

// SetContactOwner sets the "contact_owner" edge to the Person entity.
func (pc *PersonCreate) SetContactOwner(p *Person) *PersonCreate {
	return pc.SetContactOwnerID(p.ID)
}

// AddContactIDs adds the "contacts" edge to the Person entity by IDs.
func (pc *PersonCreate) AddContactIDs(ids ...int) *PersonCreate {
	pc.mutation.AddContactIDs(ids...)
	return pc
}

// AddContacts adds the "contacts" edges to the Person entity.
func (pc *PersonCreate) AddContacts(p ...*Person) *PersonCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddContactIDs(ids...)
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (pc *PersonCreate) SetMetadataID(id string) *PersonCreate {
	pc.mutation.SetMetadataID(id)
	return pc
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableMetadataID(id *string) *PersonCreate {
	if id != nil {
		pc = pc.SetMetadataID(*id)
	}
	return pc
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (pc *PersonCreate) SetMetadata(m *Metadata) *PersonCreate {
	return pc.SetMetadataID(m.ID)
}

// SetNodeID sets the "node" edge to the Node entity by ID.
func (pc *PersonCreate) SetNodeID(id string) *PersonCreate {
	pc.mutation.SetNodeID(id)
	return pc
}

// SetNillableNodeID sets the "node" edge to the Node entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableNodeID(id *string) *PersonCreate {
	if id != nil {
		pc = pc.SetNodeID(*id)
	}
	return pc
}

// SetNode sets the "node" edge to the Node entity.
func (pc *PersonCreate) SetNode(n *Node) *PersonCreate {
	return pc.SetNodeID(n.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Person.name"`)}
	}
	if _, ok := pc.mutation.IsOrg(); !ok {
		return &ValidationError{Name: "is_org", err: errors.New(`ent: missing required field "Person.is_org"`)}
	}
	if _, ok := pc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Person.email"`)}
	}
	if _, ok := pc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Person.url"`)}
	}
	if _, ok := pc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Person.phone"`)}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.IsOrg(); ok {
		_spec.SetField(person.FieldIsOrg, field.TypeBool, value)
		_node.IsOrg = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.URL(); ok {
		_spec.SetField(person.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := pc.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := pc.mutation.ContactOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ContactOwnerTable,
			Columns: []string{person.ContactOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.person_contacts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MetadataTable,
			Columns: []string{person.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.metadata_authors = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.NodeTable,
			Columns: []string{person.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_originators = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Person.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PersonUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PersonCreate) OnConflict(opts ...sql.ConflictOption) *PersonUpsertOne {
	pc.conflict = opts
	return &PersonUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Person.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PersonCreate) OnConflictColumns(columns ...string) *PersonUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PersonUpsertOne{
		create: pc,
	}
}

type (
	// PersonUpsertOne is the builder for "upsert"-ing
	//  one Person node.
	PersonUpsertOne struct {
		create *PersonCreate
	}

	// PersonUpsert is the "OnConflict" setter.
	PersonUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PersonUpsert) SetName(v string) *PersonUpsert {
	u.Set(person.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PersonUpsert) UpdateName() *PersonUpsert {
	u.SetExcluded(person.FieldName)
	return u
}

// SetIsOrg sets the "is_org" field.
func (u *PersonUpsert) SetIsOrg(v bool) *PersonUpsert {
	u.Set(person.FieldIsOrg, v)
	return u
}

// UpdateIsOrg sets the "is_org" field to the value that was provided on create.
func (u *PersonUpsert) UpdateIsOrg() *PersonUpsert {
	u.SetExcluded(person.FieldIsOrg)
	return u
}

// SetEmail sets the "email" field.
func (u *PersonUpsert) SetEmail(v string) *PersonUpsert {
	u.Set(person.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *PersonUpsert) UpdateEmail() *PersonUpsert {
	u.SetExcluded(person.FieldEmail)
	return u
}

// SetURL sets the "url" field.
func (u *PersonUpsert) SetURL(v string) *PersonUpsert {
	u.Set(person.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PersonUpsert) UpdateURL() *PersonUpsert {
	u.SetExcluded(person.FieldURL)
	return u
}

// SetPhone sets the "phone" field.
func (u *PersonUpsert) SetPhone(v string) *PersonUpsert {
	u.Set(person.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *PersonUpsert) UpdatePhone() *PersonUpsert {
	u.SetExcluded(person.FieldPhone)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Person.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PersonUpsertOne) UpdateNewValues() *PersonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Person.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PersonUpsertOne) Ignore() *PersonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PersonUpsertOne) DoNothing() *PersonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PersonCreate.OnConflict
// documentation for more info.
func (u *PersonUpsertOne) Update(set func(*PersonUpsert)) *PersonUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PersonUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PersonUpsertOne) SetName(v string) *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PersonUpsertOne) UpdateName() *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateName()
	})
}

// SetIsOrg sets the "is_org" field.
func (u *PersonUpsertOne) SetIsOrg(v bool) *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.SetIsOrg(v)
	})
}

// UpdateIsOrg sets the "is_org" field to the value that was provided on create.
func (u *PersonUpsertOne) UpdateIsOrg() *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateIsOrg()
	})
}

// SetEmail sets the "email" field.
func (u *PersonUpsertOne) SetEmail(v string) *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *PersonUpsertOne) UpdateEmail() *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateEmail()
	})
}

// SetURL sets the "url" field.
func (u *PersonUpsertOne) SetURL(v string) *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PersonUpsertOne) UpdateURL() *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateURL()
	})
}

// SetPhone sets the "phone" field.
func (u *PersonUpsertOne) SetPhone(v string) *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *PersonUpsertOne) UpdatePhone() *PersonUpsertOne {
	return u.Update(func(s *PersonUpsert) {
		s.UpdatePhone()
	})
}

// Exec executes the query.
func (u *PersonUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PersonCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PersonUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PersonUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PersonUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	err      error
	builders []*PersonCreate
	conflict []sql.ConflictOption
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Person.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PersonUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PersonCreateBulk) OnConflict(opts ...sql.ConflictOption) *PersonUpsertBulk {
	pcb.conflict = opts
	return &PersonUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Person.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PersonCreateBulk) OnConflictColumns(columns ...string) *PersonUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PersonUpsertBulk{
		create: pcb,
	}
}

// PersonUpsertBulk is the builder for "upsert"-ing
// a bulk of Person nodes.
type PersonUpsertBulk struct {
	create *PersonCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Person.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PersonUpsertBulk) UpdateNewValues() *PersonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Person.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PersonUpsertBulk) Ignore() *PersonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PersonUpsertBulk) DoNothing() *PersonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PersonCreateBulk.OnConflict
// documentation for more info.
func (u *PersonUpsertBulk) Update(set func(*PersonUpsert)) *PersonUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PersonUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PersonUpsertBulk) SetName(v string) *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PersonUpsertBulk) UpdateName() *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateName()
	})
}

// SetIsOrg sets the "is_org" field.
func (u *PersonUpsertBulk) SetIsOrg(v bool) *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.SetIsOrg(v)
	})
}

// UpdateIsOrg sets the "is_org" field to the value that was provided on create.
func (u *PersonUpsertBulk) UpdateIsOrg() *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateIsOrg()
	})
}

// SetEmail sets the "email" field.
func (u *PersonUpsertBulk) SetEmail(v string) *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *PersonUpsertBulk) UpdateEmail() *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateEmail()
	})
}

// SetURL sets the "url" field.
func (u *PersonUpsertBulk) SetURL(v string) *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *PersonUpsertBulk) UpdateURL() *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.UpdateURL()
	})
}

// SetPhone sets the "phone" field.
func (u *PersonUpsertBulk) SetPhone(v string) *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *PersonUpsertBulk) UpdatePhone() *PersonUpsertBulk {
	return u.Update(func(s *PersonUpsert) {
		s.UpdatePhone()
	})
}

// Exec executes the query.
func (u *PersonUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PersonCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PersonCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PersonUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
