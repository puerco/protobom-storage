// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package hashesentry

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/protobom/storage/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldLTE(FieldID, id))
}

// HashData applies equality check predicate on the "hash_data" field. It's identical to HashDataEQ.
func HashData(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEQ(FieldHashData, v))
}

// HashAlgorithmTypeEQ applies the EQ predicate on the "hash_algorithm_type" field.
func HashAlgorithmTypeEQ(v HashAlgorithmType) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEQ(FieldHashAlgorithmType, v))
}

// HashAlgorithmTypeNEQ applies the NEQ predicate on the "hash_algorithm_type" field.
func HashAlgorithmTypeNEQ(v HashAlgorithmType) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNEQ(FieldHashAlgorithmType, v))
}

// HashAlgorithmTypeIn applies the In predicate on the "hash_algorithm_type" field.
func HashAlgorithmTypeIn(vs ...HashAlgorithmType) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldIn(FieldHashAlgorithmType, vs...))
}

// HashAlgorithmTypeNotIn applies the NotIn predicate on the "hash_algorithm_type" field.
func HashAlgorithmTypeNotIn(vs ...HashAlgorithmType) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNotIn(FieldHashAlgorithmType, vs...))
}

// HashDataEQ applies the EQ predicate on the "hash_data" field.
func HashDataEQ(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEQ(FieldHashData, v))
}

// HashDataNEQ applies the NEQ predicate on the "hash_data" field.
func HashDataNEQ(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNEQ(FieldHashData, v))
}

// HashDataIn applies the In predicate on the "hash_data" field.
func HashDataIn(vs ...string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldIn(FieldHashData, vs...))
}

// HashDataNotIn applies the NotIn predicate on the "hash_data" field.
func HashDataNotIn(vs ...string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldNotIn(FieldHashData, vs...))
}

// HashDataGT applies the GT predicate on the "hash_data" field.
func HashDataGT(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldGT(FieldHashData, v))
}

// HashDataGTE applies the GTE predicate on the "hash_data" field.
func HashDataGTE(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldGTE(FieldHashData, v))
}

// HashDataLT applies the LT predicate on the "hash_data" field.
func HashDataLT(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldLT(FieldHashData, v))
}

// HashDataLTE applies the LTE predicate on the "hash_data" field.
func HashDataLTE(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldLTE(FieldHashData, v))
}

// HashDataContains applies the Contains predicate on the "hash_data" field.
func HashDataContains(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldContains(FieldHashData, v))
}

// HashDataHasPrefix applies the HasPrefix predicate on the "hash_data" field.
func HashDataHasPrefix(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldHasPrefix(FieldHashData, v))
}

// HashDataHasSuffix applies the HasSuffix predicate on the "hash_data" field.
func HashDataHasSuffix(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldHasSuffix(FieldHashData, v))
}

// HashDataEqualFold applies the EqualFold predicate on the "hash_data" field.
func HashDataEqualFold(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldEqualFold(FieldHashData, v))
}

// HashDataContainsFold applies the ContainsFold predicate on the "hash_data" field.
func HashDataContainsFold(v string) predicate.HashesEntry {
	return predicate.HashesEntry(sql.FieldContainsFold(FieldHashData, v))
}

// HasExternalReferences applies the HasEdge predicate on the "external_references" edge.
func HasExternalReferences() predicate.HashesEntry {
	return predicate.HashesEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExternalReferencesTable, ExternalReferencesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExternalReferencesWith applies the HasEdge predicate on the "external_references" edge with a given conditions (other predicates).
func HasExternalReferencesWith(preds ...predicate.ExternalReference) predicate.HashesEntry {
	return predicate.HashesEntry(func(s *sql.Selector) {
		step := newExternalReferencesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.HashesEntry {
	return predicate.HashesEntry(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodesTable, NodesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.HashesEntry {
	return predicate.HashesEntry(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HashesEntry) predicate.HashesEntry {
	return predicate.HashesEntry(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HashesEntry) predicate.HashesEntry {
	return predicate.HashesEntry(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HashesEntry) predicate.HashesEntry {
	return predicate.HashesEntry(sql.NotPredicates(p))
}