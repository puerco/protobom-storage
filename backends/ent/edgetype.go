// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/protobom/storage/backends/ent/edgetype"
	"github.com/protobom/storage/backends/ent/node"
)

// EdgeType is the model entity for the EdgeType schema.
type EdgeType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type edgetype.Type `json:"type,omitempty"`
	// NodeID holds the value of the "node_id" field.
	NodeID string `json:"node_id,omitempty"`
	// ToID holds the value of the "to_id" field.
	ToID string `json:"to_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EdgeTypeQuery when eager-loading is set.
	Edges        EdgeTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EdgeTypeEdges holds the relations/edges for other nodes in the graph.
type EdgeTypeEdges struct {
	// From holds the value of the from edge.
	From *Node `json:"from,omitempty"`
	// To holds the value of the to edge.
	To *Node `json:"to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FromOrErr returns the From value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EdgeTypeEdges) FromOrErr() (*Node, error) {
	if e.From != nil {
		return e.From, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: node.Label}
	}
	return nil, &NotLoadedError{edge: "from"}
}

// ToOrErr returns the To value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EdgeTypeEdges) ToOrErr() (*Node, error) {
	if e.To != nil {
		return e.To, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: node.Label}
	}
	return nil, &NotLoadedError{edge: "to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EdgeType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case edgetype.FieldID:
			values[i] = new(sql.NullInt64)
		case edgetype.FieldType, edgetype.FieldNodeID, edgetype.FieldToID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EdgeType fields.
func (et *EdgeType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case edgetype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			et.ID = int(value.Int64)
		case edgetype.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				et.Type = edgetype.Type(value.String)
			}
		case edgetype.FieldNodeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_id", values[i])
			} else if value.Valid {
				et.NodeID = value.String
			}
		case edgetype.FieldToID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to_id", values[i])
			} else if value.Valid {
				et.ToID = value.String
			}
		default:
			et.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EdgeType.
// This includes values selected through modifiers, order, etc.
func (et *EdgeType) Value(name string) (ent.Value, error) {
	return et.selectValues.Get(name)
}

// QueryFrom queries the "from" edge of the EdgeType entity.
func (et *EdgeType) QueryFrom() *NodeQuery {
	return NewEdgeTypeClient(et.config).QueryFrom(et)
}

// QueryTo queries the "to" edge of the EdgeType entity.
func (et *EdgeType) QueryTo() *NodeQuery {
	return NewEdgeTypeClient(et.config).QueryTo(et)
}

// Update returns a builder for updating this EdgeType.
// Note that you need to call EdgeType.Unwrap() before calling this method if this EdgeType
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *EdgeType) Update() *EdgeTypeUpdateOne {
	return NewEdgeTypeClient(et.config).UpdateOne(et)
}

// Unwrap unwraps the EdgeType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (et *EdgeType) Unwrap() *EdgeType {
	_tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("ent: EdgeType is not a transactional entity")
	}
	et.config.driver = _tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *EdgeType) String() string {
	var builder strings.Builder
	builder.WriteString("EdgeType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", et.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", et.Type))
	builder.WriteString(", ")
	builder.WriteString("node_id=")
	builder.WriteString(et.NodeID)
	builder.WriteString(", ")
	builder.WriteString("to_id=")
	builder.WriteString(et.ToID)
	builder.WriteByte(')')
	return builder.String()
}

// EdgeTypes is a parsable slice of EdgeType.
type EdgeTypes []*EdgeType