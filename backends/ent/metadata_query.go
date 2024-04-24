// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/backends/ent/document"
	"github.com/protobom/storage/backends/ent/documenttype"
	"github.com/protobom/storage/backends/ent/metadata"
	"github.com/protobom/storage/backends/ent/person"
	"github.com/protobom/storage/backends/ent/predicate"
	"github.com/protobom/storage/backends/ent/tool"
)

// MetadataQuery is the builder for querying Metadata entities.
type MetadataQuery struct {
	config
	ctx               *QueryContext
	order             []metadata.OrderOption
	inters            []Interceptor
	predicates        []predicate.Metadata
	withTools         *ToolQuery
	withAuthors       *PersonQuery
	withDocumentTypes *DocumentTypeQuery
	withDocument      *DocumentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MetadataQuery builder.
func (mq *MetadataQuery) Where(ps ...predicate.Metadata) *MetadataQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MetadataQuery) Limit(limit int) *MetadataQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MetadataQuery) Offset(offset int) *MetadataQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MetadataQuery) Unique(unique bool) *MetadataQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MetadataQuery) Order(o ...metadata.OrderOption) *MetadataQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryTools chains the current query on the "tools" edge.
func (mq *MetadataQuery) QueryTools() *ToolQuery {
	query := (&ToolClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(tool.Table, tool.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metadata.ToolsTable, metadata.ToolsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthors chains the current query on the "authors" edge.
func (mq *MetadataQuery) QueryAuthors() *PersonQuery {
	query := (&PersonClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metadata.AuthorsTable, metadata.AuthorsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocumentTypes chains the current query on the "document_types" edge.
func (mq *MetadataQuery) QueryDocumentTypes() *DocumentTypeQuery {
	query := (&DocumentTypeClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(documenttype.Table, documenttype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, metadata.DocumentTypesTable, metadata.DocumentTypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocument chains the current query on the "document" edge.
func (mq *MetadataQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(metadata.Table, metadata.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, metadata.DocumentTable, metadata.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Metadata entity from the query.
// Returns a *NotFoundError when no Metadata was found.
func (mq *MetadataQuery) First(ctx context.Context) (*Metadata, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{metadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MetadataQuery) FirstX(ctx context.Context) *Metadata {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Metadata ID from the query.
// Returns a *NotFoundError when no Metadata ID was found.
func (mq *MetadataQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{metadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MetadataQuery) FirstIDX(ctx context.Context) string {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Metadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Metadata entity is found.
// Returns a *NotFoundError when no Metadata entities are found.
func (mq *MetadataQuery) Only(ctx context.Context) (*Metadata, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{metadata.Label}
	default:
		return nil, &NotSingularError{metadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MetadataQuery) OnlyX(ctx context.Context) *Metadata {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Metadata ID in the query.
// Returns a *NotSingularError when more than one Metadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MetadataQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{metadata.Label}
	default:
		err = &NotSingularError{metadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MetadataQuery) OnlyIDX(ctx context.Context) string {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MetadataSlice.
func (mq *MetadataQuery) All(ctx context.Context) ([]*Metadata, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Metadata, *MetadataQuery]()
	return withInterceptors[[]*Metadata](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MetadataQuery) AllX(ctx context.Context) []*Metadata {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Metadata IDs.
func (mq *MetadataQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(metadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MetadataQuery) IDsX(ctx context.Context) []string {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MetadataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MetadataQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MetadataQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MetadataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MetadataQuery) Clone() *MetadataQuery {
	if mq == nil {
		return nil
	}
	return &MetadataQuery{
		config:            mq.config,
		ctx:               mq.ctx.Clone(),
		order:             append([]metadata.OrderOption{}, mq.order...),
		inters:            append([]Interceptor{}, mq.inters...),
		predicates:        append([]predicate.Metadata{}, mq.predicates...),
		withTools:         mq.withTools.Clone(),
		withAuthors:       mq.withAuthors.Clone(),
		withDocumentTypes: mq.withDocumentTypes.Clone(),
		withDocument:      mq.withDocument.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithTools tells the query-builder to eager-load the nodes that are connected to
// the "tools" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithTools(opts ...func(*ToolQuery)) *MetadataQuery {
	query := (&ToolClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withTools = query
	return mq
}

// WithAuthors tells the query-builder to eager-load the nodes that are connected to
// the "authors" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithAuthors(opts ...func(*PersonQuery)) *MetadataQuery {
	query := (&PersonClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withAuthors = query
	return mq
}

// WithDocumentTypes tells the query-builder to eager-load the nodes that are connected to
// the "document_types" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithDocumentTypes(opts ...func(*DocumentTypeQuery)) *MetadataQuery {
	query := (&DocumentTypeClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withDocumentTypes = query
	return mq
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MetadataQuery) WithDocument(opts ...func(*DocumentQuery)) *MetadataQuery {
	query := (&DocumentClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withDocument = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version string `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Metadata.Query().
//		GroupBy(metadata.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MetadataQuery) GroupBy(field string, fields ...string) *MetadataGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MetadataGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = metadata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version string `json:"version,omitempty"`
//	}
//
//	client.Metadata.Query().
//		Select(metadata.FieldVersion).
//		Scan(ctx, &v)
func (mq *MetadataQuery) Select(fields ...string) *MetadataSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MetadataSelect{MetadataQuery: mq}
	sbuild.label = metadata.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MetadataSelect configured with the given aggregations.
func (mq *MetadataQuery) Aggregate(fns ...AggregateFunc) *MetadataSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MetadataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !metadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Metadata, error) {
	var (
		nodes       = []*Metadata{}
		_spec       = mq.querySpec()
		loadedTypes = [4]bool{
			mq.withTools != nil,
			mq.withAuthors != nil,
			mq.withDocumentTypes != nil,
			mq.withDocument != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Metadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Metadata{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withTools; query != nil {
		if err := mq.loadTools(ctx, query, nodes,
			func(n *Metadata) { n.Edges.Tools = []*Tool{} },
			func(n *Metadata, e *Tool) { n.Edges.Tools = append(n.Edges.Tools, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withAuthors; query != nil {
		if err := mq.loadAuthors(ctx, query, nodes,
			func(n *Metadata) { n.Edges.Authors = []*Person{} },
			func(n *Metadata, e *Person) { n.Edges.Authors = append(n.Edges.Authors, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withDocumentTypes; query != nil {
		if err := mq.loadDocumentTypes(ctx, query, nodes,
			func(n *Metadata) { n.Edges.DocumentTypes = []*DocumentType{} },
			func(n *Metadata, e *DocumentType) { n.Edges.DocumentTypes = append(n.Edges.DocumentTypes, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withDocument; query != nil {
		if err := mq.loadDocument(ctx, query, nodes, nil,
			func(n *Metadata, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MetadataQuery) loadTools(ctx context.Context, query *ToolQuery, nodes []*Metadata, init func(*Metadata), assign func(*Metadata, *Tool)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Metadata)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Tool(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(metadata.ToolsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.metadata_tools
		if fk == nil {
			return fmt.Errorf(`foreign-key "metadata_tools" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "metadata_tools" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MetadataQuery) loadAuthors(ctx context.Context, query *PersonQuery, nodes []*Metadata, init func(*Metadata), assign func(*Metadata, *Person)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Metadata)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Person(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(metadata.AuthorsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.metadata_authors
		if fk == nil {
			return fmt.Errorf(`foreign-key "metadata_authors" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "metadata_authors" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MetadataQuery) loadDocumentTypes(ctx context.Context, query *DocumentTypeQuery, nodes []*Metadata, init func(*Metadata), assign func(*Metadata, *DocumentType)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Metadata)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DocumentType(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(metadata.DocumentTypesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.metadata_document_types
		if fk == nil {
			return fmt.Errorf(`foreign-key "metadata_document_types" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "metadata_document_types" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MetadataQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*Metadata, init func(*Metadata), assign func(*Metadata, *Document)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Metadata)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Document(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(metadata.DocumentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.metadata_document
		if fk == nil {
			return fmt.Errorf(`foreign-key "metadata_document" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "metadata_document" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for i := range fields {
			if fields[i] != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(metadata.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = metadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MetadataGroupBy is the group-by builder for Metadata entities.
type MetadataGroupBy struct {
	selector
	build *MetadataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MetadataGroupBy) Aggregate(fns ...AggregateFunc) *MetadataGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MetadataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MetadataQuery, *MetadataGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MetadataGroupBy) sqlScan(ctx context.Context, root *MetadataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MetadataSelect is the builder for selecting fields of Metadata entities.
type MetadataSelect struct {
	*MetadataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MetadataSelect) Aggregate(fns ...AggregateFunc) *MetadataSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MetadataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MetadataQuery, *MetadataSelect](ctx, ms.MetadataQuery, ms, ms.inters, v)
}

func (ms *MetadataSelect) sqlScan(ctx context.Context, root *MetadataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
