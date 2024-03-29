// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/category"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/item"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/predicate"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/store"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *ItemUpdate) SetName(s string) *ItemUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNote sets the "note" field.
func (iu *ItemUpdate) SetNote(s string) *ItemUpdate {
	iu.mutation.SetNote(s)
	return iu
}

// SetAmount sets the "amount" field.
func (iu *ItemUpdate) SetAmount(i int) *ItemUpdate {
	iu.mutation.ResetAmount()
	iu.mutation.SetAmount(i)
	return iu
}

// AddAmount adds i to the "amount" field.
func (iu *ItemUpdate) AddAmount(i int) *ItemUpdate {
	iu.mutation.AddAmount(i)
	return iu
}

// SetStatus sets the "status" field.
func (iu *ItemUpdate) SetStatus(s string) *ItemUpdate {
	iu.mutation.SetStatus(s)
	return iu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableStatus(s *string) *ItemUpdate {
	if s != nil {
		iu.SetStatus(*s)
	}
	return iu
}

// SetStoreID sets the "store_id" field.
func (iu *ItemUpdate) SetStoreID(i int) *ItemUpdate {
	iu.mutation.SetStoreID(i)
	return iu
}

// SetNillableStoreID sets the "store_id" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableStoreID(i *int) *ItemUpdate {
	if i != nil {
		iu.SetStoreID(*i)
	}
	return iu
}

// ClearStoreID clears the value of the "store_id" field.
func (iu *ItemUpdate) ClearStoreID() *ItemUpdate {
	iu.mutation.ClearStoreID()
	return iu
}

// SetCategoryID sets the "category_id" field.
func (iu *ItemUpdate) SetCategoryID(i int) *ItemUpdate {
	iu.mutation.SetCategoryID(i)
	return iu
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableCategoryID(i *int) *ItemUpdate {
	if i != nil {
		iu.SetCategoryID(*i)
	}
	return iu
}

// ClearCategoryID clears the value of the "category_id" field.
func (iu *ItemUpdate) ClearCategoryID() *ItemUpdate {
	iu.mutation.ClearCategoryID()
	return iu
}

// SetStore sets the "store" edge to the Store entity.
func (iu *ItemUpdate) SetStore(s *Store) *ItemUpdate {
	return iu.SetStoreID(s.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (iu *ItemUpdate) SetCategory(c *Category) *ItemUpdate {
	return iu.SetCategoryID(c.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearStore clears the "store" edge to the Store entity.
func (iu *ItemUpdate) ClearStore() *ItemUpdate {
	iu.mutation.ClearStore()
	return iu
}

// ClearCategory clears the "category" edge to the Category entity.
func (iu *ItemUpdate) ClearCategory() *ItemUpdate {
	iu.mutation.ClearCategory()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := item.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Item.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Amount(); ok {
		if err := item.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Item.amount": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Status(); ok {
		if err := item.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Item.status": %w`, err)}
		}
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Note(); ok {
		_spec.SetField(item.FieldNote, field.TypeString, value)
	}
	if value, ok := iu.mutation.Amount(); ok {
		_spec.SetField(item.FieldAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedAmount(); ok {
		_spec.AddField(item.FieldAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Status(); ok {
		_spec.SetField(item.FieldStatus, field.TypeString, value)
	}
	if iu.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: []string{item.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: []string{item.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetName sets the "name" field.
func (iuo *ItemUpdateOne) SetName(s string) *ItemUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNote sets the "note" field.
func (iuo *ItemUpdateOne) SetNote(s string) *ItemUpdateOne {
	iuo.mutation.SetNote(s)
	return iuo
}

// SetAmount sets the "amount" field.
func (iuo *ItemUpdateOne) SetAmount(i int) *ItemUpdateOne {
	iuo.mutation.ResetAmount()
	iuo.mutation.SetAmount(i)
	return iuo
}

// AddAmount adds i to the "amount" field.
func (iuo *ItemUpdateOne) AddAmount(i int) *ItemUpdateOne {
	iuo.mutation.AddAmount(i)
	return iuo
}

// SetStatus sets the "status" field.
func (iuo *ItemUpdateOne) SetStatus(s string) *ItemUpdateOne {
	iuo.mutation.SetStatus(s)
	return iuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableStatus(s *string) *ItemUpdateOne {
	if s != nil {
		iuo.SetStatus(*s)
	}
	return iuo
}

// SetStoreID sets the "store_id" field.
func (iuo *ItemUpdateOne) SetStoreID(i int) *ItemUpdateOne {
	iuo.mutation.SetStoreID(i)
	return iuo
}

// SetNillableStoreID sets the "store_id" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableStoreID(i *int) *ItemUpdateOne {
	if i != nil {
		iuo.SetStoreID(*i)
	}
	return iuo
}

// ClearStoreID clears the value of the "store_id" field.
func (iuo *ItemUpdateOne) ClearStoreID() *ItemUpdateOne {
	iuo.mutation.ClearStoreID()
	return iuo
}

// SetCategoryID sets the "category_id" field.
func (iuo *ItemUpdateOne) SetCategoryID(i int) *ItemUpdateOne {
	iuo.mutation.SetCategoryID(i)
	return iuo
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableCategoryID(i *int) *ItemUpdateOne {
	if i != nil {
		iuo.SetCategoryID(*i)
	}
	return iuo
}

// ClearCategoryID clears the value of the "category_id" field.
func (iuo *ItemUpdateOne) ClearCategoryID() *ItemUpdateOne {
	iuo.mutation.ClearCategoryID()
	return iuo
}

// SetStore sets the "store" edge to the Store entity.
func (iuo *ItemUpdateOne) SetStore(s *Store) *ItemUpdateOne {
	return iuo.SetStoreID(s.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (iuo *ItemUpdateOne) SetCategory(c *Category) *ItemUpdateOne {
	return iuo.SetCategoryID(c.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearStore clears the "store" edge to the Store entity.
func (iuo *ItemUpdateOne) ClearStore() *ItemUpdateOne {
	iuo.mutation.ClearStore()
	return iuo
}

// ClearCategory clears the "category" edge to the Category entity.
func (iuo *ItemUpdateOne) ClearCategory() *ItemUpdateOne {
	iuo.mutation.ClearCategory()
	return iuo
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := item.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Item.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Amount(); ok {
		if err := item.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Item.amount": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Status(); ok {
		if err := item.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Item.status": %w`, err)}
		}
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Note(); ok {
		_spec.SetField(item.FieldNote, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Amount(); ok {
		_spec.SetField(item.FieldAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedAmount(); ok {
		_spec.AddField(item.FieldAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Status(); ok {
		_spec.SetField(item.FieldStatus, field.TypeString, value)
	}
	if iuo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: []string{item.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: []string{item.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
