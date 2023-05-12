// Code generated by ent, DO NOT EDIT.

package balance

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the balance type in the database.
	Label = "balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// Table holds the table name of the balance in the database.
	Table = "balance"
)

// Columns holds all SQL columns for balance fields.
var Columns = []string{
	FieldID,
	FieldBalance,
	FieldLocked,
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
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance float64
	// BalanceValidator is a validator for the "balance" field. It is called by the builders before save.
	BalanceValidator func(float64) error
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked float64
	// LockedValidator is a validator for the "locked" field. It is called by the builders before save.
	LockedValidator func(float64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Balance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBalance orders the results by the balance field.
func ByBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBalance, opts...).ToFunc()
}

// ByLocked orders the results by the locked field.
func ByLocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocked, opts...).ToFunc()
}
