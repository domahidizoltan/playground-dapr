// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
)

// Balance is the model entity for the Balance schema.
type Balance struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked       float64 `json:"locked,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Balance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case balance.FieldBalance, balance.FieldLocked:
			values[i] = new(sql.NullFloat64)
		case balance.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Balance fields.
func (b *Balance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case balance.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				b.ID = value.String
			}
		case balance.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				b.Balance = value.Float64
			}
		case balance.FieldLocked:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				b.Locked = value.Float64
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Balance.
// This includes values selected through modifiers, order, etc.
func (b *Balance) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// Update returns a builder for updating this Balance.
// Note that you need to call Balance.Unwrap() before calling this method if this Balance
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Balance) Update() *BalanceUpdateOne {
	return NewBalanceClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Balance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Balance) Unwrap() *Balance {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("generated: Balance is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Balance) String() string {
	var builder strings.Builder
	builder.WriteString("Balance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", b.Balance))
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", b.Locked))
	builder.WriteByte(')')
	return builder.String()
}

// Balances is a parsable slice of Balance.
type Balances []*Balance
