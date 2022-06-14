// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent/predicate"
	"github.com/Faroukhamadi/Banketeer/teller_service/tellerdb/ent/teller"
)

// TellerUpdate is the builder for updating Teller entities.
type TellerUpdate struct {
	config
	hooks    []Hook
	mutation *TellerMutation
}

// Where appends a list predicates to the TellerUpdate builder.
func (tu *TellerUpdate) Where(ps ...predicate.Teller) *TellerUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TellerUpdate) SetUpdateTime(t time.Time) *TellerUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetUsername sets the "username" field.
func (tu *TellerUpdate) SetUsername(s string) *TellerUpdate {
	tu.mutation.SetUsername(s)
	return tu
}

// SetPassword sets the "password" field.
func (tu *TellerUpdate) SetPassword(s string) *TellerUpdate {
	tu.mutation.SetPassword(s)
	return tu
}

// SetRole sets the "role" field.
func (tu *TellerUpdate) SetRole(t teller.Role) *TellerUpdate {
	tu.mutation.SetRole(t)
	return tu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (tu *TellerUpdate) SetNillableRole(t *teller.Role) *TellerUpdate {
	if t != nil {
		tu.SetRole(*t)
	}
	return tu
}

// Mutation returns the TellerMutation object of the builder.
func (tu *TellerUpdate) Mutation() *TellerMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TellerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TellerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TellerUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TellerUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TellerUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TellerUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := teller.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TellerUpdate) check() error {
	if v, ok := tu.mutation.Role(); ok {
		if err := teller.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Teller.role": %w`, err)}
		}
	}
	return nil
}

func (tu *TellerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teller.Table,
			Columns: teller.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: teller.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teller.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teller.FieldUsername,
		})
	}
	if value, ok := tu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teller.FieldPassword,
		})
	}
	if value, ok := tu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: teller.FieldRole,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teller.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TellerUpdateOne is the builder for updating a single Teller entity.
type TellerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TellerMutation
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TellerUpdateOne) SetUpdateTime(t time.Time) *TellerUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetUsername sets the "username" field.
func (tuo *TellerUpdateOne) SetUsername(s string) *TellerUpdateOne {
	tuo.mutation.SetUsername(s)
	return tuo
}

// SetPassword sets the "password" field.
func (tuo *TellerUpdateOne) SetPassword(s string) *TellerUpdateOne {
	tuo.mutation.SetPassword(s)
	return tuo
}

// SetRole sets the "role" field.
func (tuo *TellerUpdateOne) SetRole(t teller.Role) *TellerUpdateOne {
	tuo.mutation.SetRole(t)
	return tuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (tuo *TellerUpdateOne) SetNillableRole(t *teller.Role) *TellerUpdateOne {
	if t != nil {
		tuo.SetRole(*t)
	}
	return tuo
}

// Mutation returns the TellerMutation object of the builder.
func (tuo *TellerUpdateOne) Mutation() *TellerMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TellerUpdateOne) Select(field string, fields ...string) *TellerUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teller entity.
func (tuo *TellerUpdateOne) Save(ctx context.Context) (*Teller, error) {
	var (
		err  error
		node *Teller
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TellerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Teller)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TellerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TellerUpdateOne) SaveX(ctx context.Context) *Teller {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TellerUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TellerUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TellerUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := teller.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TellerUpdateOne) check() error {
	if v, ok := tuo.mutation.Role(); ok {
		if err := teller.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Teller.role": %w`, err)}
		}
	}
	return nil
}

func (tuo *TellerUpdateOne) sqlSave(ctx context.Context) (_node *Teller, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teller.Table,
			Columns: teller.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: teller.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Teller.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teller.FieldID)
		for _, f := range fields {
			if !teller.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teller.FieldID {
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
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teller.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teller.FieldUsername,
		})
	}
	if value, ok := tuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teller.FieldPassword,
		})
	}
	if value, ok := tuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: teller.FieldRole,
		})
	}
	_node = &Teller{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teller.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}