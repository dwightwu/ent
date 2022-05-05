// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv1/customtype"
	"entgo.io/ent/entc/integration/migrate/entv1/predicate"
	"entgo.io/ent/schema/field"
)

// CustomTypeUpdate is the builder for updating CustomType entities.
type CustomTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CustomTypeMutation
}

// Where appends a list predicates to the CustomTypeUpdate builder.
func (ctu *CustomTypeUpdate) Where(ps ...predicate.CustomType) *CustomTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetCustom sets the "custom" field.
func (ctu *CustomTypeUpdate) SetCustom(s string) *CustomTypeUpdate {
	ctu.mutation.SetCustom(s)
	return ctu
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctu *CustomTypeUpdate) SetNillableCustom(s *string) *CustomTypeUpdate {
	if s != nil {
		ctu.SetCustom(*s)
	}
	return ctu
}

// ClearCustom clears the value of the "custom" field.
func (ctu *CustomTypeUpdate) ClearCustom() *CustomTypeUpdate {
	ctu.mutation.ClearCustom()
	return ctu
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctu *CustomTypeUpdate) Mutation() *CustomTypeMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CustomTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("entv1: uninitialized hook (forgotten import entv1/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CustomTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CustomTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CustomTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *CustomTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customtype.Table,
			Columns: customtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Custom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customtype.FieldCustom,
		})
	}
	if ctu.mutation.CustomCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customtype.FieldCustom,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CustomTypeUpdateOne is the builder for updating a single CustomType entity.
type CustomTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomTypeMutation
}

// SetCustom sets the "custom" field.
func (ctuo *CustomTypeUpdateOne) SetCustom(s string) *CustomTypeUpdateOne {
	ctuo.mutation.SetCustom(s)
	return ctuo
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctuo *CustomTypeUpdateOne) SetNillableCustom(s *string) *CustomTypeUpdateOne {
	if s != nil {
		ctuo.SetCustom(*s)
	}
	return ctuo
}

// ClearCustom clears the value of the "custom" field.
func (ctuo *CustomTypeUpdateOne) ClearCustom() *CustomTypeUpdateOne {
	ctuo.mutation.ClearCustom()
	return ctuo
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctuo *CustomTypeUpdateOne) Mutation() *CustomTypeMutation {
	return ctuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CustomTypeUpdateOne) Select(field string, fields ...string) *CustomTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CustomType entity.
func (ctuo *CustomTypeUpdateOne) Save(ctx context.Context) (*CustomType, error) {
	var (
		err  error
		node *CustomType
	)
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("entv1: uninitialized hook (forgotten import entv1/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CustomType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CustomTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CustomTypeUpdateOne) SaveX(ctx context.Context) *CustomType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CustomTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CustomTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *CustomTypeUpdateOne) sqlSave(ctx context.Context) (_node *CustomType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customtype.Table,
			Columns: customtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv1: missing "CustomType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customtype.FieldID)
		for _, f := range fields {
			if !customtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv1: invalid field %q for query", f)}
			}
			if f != customtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Custom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customtype.FieldCustom,
		})
	}
	if ctuo.mutation.CustomCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customtype.FieldCustom,
		})
	}
	_node = &CustomType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
