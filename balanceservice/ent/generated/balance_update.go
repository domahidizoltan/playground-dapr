// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/predicate"
)

// BalanceUpdate is the builder for updating Balance entities.
type BalanceUpdate struct {
	config
	hooks    []Hook
	mutation *BalanceMutation
}

// Where appends a list predicates to the BalanceUpdate builder.
func (bu *BalanceUpdate) Where(ps ...predicate.Balance) *BalanceUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBalance sets the "balance" field.
func (bu *BalanceUpdate) SetBalance(f float64) *BalanceUpdate {
	bu.mutation.ResetBalance()
	bu.mutation.SetBalance(f)
	return bu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableBalance(f *float64) *BalanceUpdate {
	if f != nil {
		bu.SetBalance(*f)
	}
	return bu
}

// AddBalance adds f to the "balance" field.
func (bu *BalanceUpdate) AddBalance(f float64) *BalanceUpdate {
	bu.mutation.AddBalance(f)
	return bu
}

// SetPending sets the "pending" field.
func (bu *BalanceUpdate) SetPending(f float64) *BalanceUpdate {
	bu.mutation.ResetPending()
	bu.mutation.SetPending(f)
	return bu
}

// SetNillablePending sets the "pending" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillablePending(f *float64) *BalanceUpdate {
	if f != nil {
		bu.SetPending(*f)
	}
	return bu
}

// AddPending adds f to the "pending" field.
func (bu *BalanceUpdate) AddPending(f float64) *BalanceUpdate {
	bu.mutation.AddPending(f)
	return bu
}

// Mutation returns the BalanceMutation object of the builder.
func (bu *BalanceUpdate) Mutation() *BalanceMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BalanceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BalanceUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BalanceUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BalanceUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BalanceUpdate) check() error {
	if v, ok := bu.mutation.Balance(); ok {
		if err := balance.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`generated: validator failed for field "Balance.balance": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Pending(); ok {
		if err := balance.PendingValidator(v); err != nil {
			return &ValidationError{Name: "pending", err: fmt.Errorf(`generated: validator failed for field "Balance.pending": %w`, err)}
		}
	}
	return nil
}

func (bu *BalanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(balance.Table, balance.Columns, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeString))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Balance(); ok {
		_spec.SetField(balance.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedBalance(); ok {
		_spec.AddField(balance.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.Pending(); ok {
		_spec.SetField(balance.FieldPending, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedPending(); ok {
		_spec.AddField(balance.FieldPending, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BalanceUpdateOne is the builder for updating a single Balance entity.
type BalanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BalanceMutation
}

// SetBalance sets the "balance" field.
func (buo *BalanceUpdateOne) SetBalance(f float64) *BalanceUpdateOne {
	buo.mutation.ResetBalance()
	buo.mutation.SetBalance(f)
	return buo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableBalance(f *float64) *BalanceUpdateOne {
	if f != nil {
		buo.SetBalance(*f)
	}
	return buo
}

// AddBalance adds f to the "balance" field.
func (buo *BalanceUpdateOne) AddBalance(f float64) *BalanceUpdateOne {
	buo.mutation.AddBalance(f)
	return buo
}

// SetPending sets the "pending" field.
func (buo *BalanceUpdateOne) SetPending(f float64) *BalanceUpdateOne {
	buo.mutation.ResetPending()
	buo.mutation.SetPending(f)
	return buo
}

// SetNillablePending sets the "pending" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillablePending(f *float64) *BalanceUpdateOne {
	if f != nil {
		buo.SetPending(*f)
	}
	return buo
}

// AddPending adds f to the "pending" field.
func (buo *BalanceUpdateOne) AddPending(f float64) *BalanceUpdateOne {
	buo.mutation.AddPending(f)
	return buo
}

// Mutation returns the BalanceMutation object of the builder.
func (buo *BalanceUpdateOne) Mutation() *BalanceMutation {
	return buo.mutation
}

// Where appends a list predicates to the BalanceUpdate builder.
func (buo *BalanceUpdateOne) Where(ps ...predicate.Balance) *BalanceUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BalanceUpdateOne) Select(field string, fields ...string) *BalanceUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Balance entity.
func (buo *BalanceUpdateOne) Save(ctx context.Context) (*Balance, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BalanceUpdateOne) SaveX(ctx context.Context) *Balance {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BalanceUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BalanceUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BalanceUpdateOne) check() error {
	if v, ok := buo.mutation.Balance(); ok {
		if err := balance.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`generated: validator failed for field "Balance.balance": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Pending(); ok {
		if err := balance.PendingValidator(v); err != nil {
			return &ValidationError{Name: "pending", err: fmt.Errorf(`generated: validator failed for field "Balance.pending": %w`, err)}
		}
	}
	return nil
}

func (buo *BalanceUpdateOne) sqlSave(ctx context.Context) (_node *Balance, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(balance.Table, balance.Columns, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeString))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Balance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balance.FieldID)
		for _, f := range fields {
			if !balance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != balance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Balance(); ok {
		_spec.SetField(balance.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedBalance(); ok {
		_spec.AddField(balance.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.Pending(); ok {
		_spec.SetField(balance.FieldPending, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedPending(); ok {
		_spec.AddField(balance.FieldPending, field.TypeFloat64, value)
	}
	_node = &Balance{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
