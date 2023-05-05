// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
)

// BalanceCreate is the builder for creating a Balance entity.
type BalanceCreate struct {
	config
	mutation *BalanceMutation
	hooks    []Hook
}

// SetBalance sets the "balance" field.
func (bc *BalanceCreate) SetBalance(f float64) *BalanceCreate {
	bc.mutation.SetBalance(f)
	return bc
}

// SetID sets the "id" field.
func (bc *BalanceCreate) SetID(s string) *BalanceCreate {
	bc.mutation.SetID(s)
	return bc
}

// Mutation returns the BalanceMutation object of the builder.
func (bc *BalanceCreate) Mutation() *BalanceMutation {
	return bc.mutation
}

// Save creates the Balance in the database.
func (bc *BalanceCreate) Save(ctx context.Context) (*Balance, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BalanceCreate) SaveX(ctx context.Context) *Balance {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BalanceCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BalanceCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BalanceCreate) check() error {
	if _, ok := bc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`generated: missing required field "Balance.balance"`)}
	}
	if v, ok := bc.mutation.Balance(); ok {
		if err := balance.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`generated: validator failed for field "Balance.balance": %w`, err)}
		}
	}
	if v, ok := bc.mutation.ID(); ok {
		if err := balance.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`generated: validator failed for field "Balance.id": %w`, err)}
		}
	}
	return nil
}

func (bc *BalanceCreate) sqlSave(ctx context.Context) (*Balance, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Balance.ID type: %T", _spec.ID.Value)
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BalanceCreate) createSpec() (*Balance, *sqlgraph.CreateSpec) {
	var (
		_node = &Balance{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(balance.Table, sqlgraph.NewFieldSpec(balance.FieldID, field.TypeString))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Balance(); ok {
		_spec.SetField(balance.FieldBalance, field.TypeFloat64, value)
		_node.Balance = value
	}
	return _node, _spec
}

// BalanceCreateBulk is the builder for creating many Balance entities in bulk.
type BalanceCreateBulk struct {
	config
	builders []*BalanceCreate
}

// Save creates the Balance entities in the database.
func (bcb *BalanceCreateBulk) Save(ctx context.Context) ([]*Balance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Balance, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BalanceMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BalanceCreateBulk) SaveX(ctx context.Context) []*Balance {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BalanceCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BalanceCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
