// Code generated by ent, DO NOT EDIT.

package store

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the store type in the database.
	Label = "store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// Table holds the table name of the store in the database.
	Table = "stores"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "store_id"
)

// Columns holds all SQL columns for store fields.
var Columns = []string{
	FieldID,
	FieldName,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Store queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}