// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BalanceColumns holds the columns for the "balance" table.
	BalanceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "locked", Type: field.TypeFloat64, Default: 0},
	}
	// BalanceTable holds the schema information for the "balance" table.
	BalanceTable = &schema.Table{
		Name:       "balance",
		Columns:    BalanceColumns,
		PrimaryKey: []*schema.Column{BalanceColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BalanceTable,
	}
)

func init() {
	BalanceTable.Annotation = &entsql.Annotation{
		Table: "balance",
	}
}
