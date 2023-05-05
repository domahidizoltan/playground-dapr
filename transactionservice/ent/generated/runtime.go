// Code generated by ent, DO NOT EDIT.

package generated

import (
	"github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated/transaction"
	"github.com/domahidizoltan/playground-dapr/transactionservice/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescTnx is the schema descriptor for tnx field.
	transactionDescTnx := transactionFields[1].Descriptor()
	// transaction.TnxValidator is a validator for the "tnx" field. It is called by the builders before save.
	transaction.TnxValidator = transactionDescTnx.Validators[0].(func(string) error)
	// transactionDescSourceAcc is the schema descriptor for source_acc field.
	transactionDescSourceAcc := transactionFields[2].Descriptor()
	// transaction.SourceAccValidator is a validator for the "source_acc" field. It is called by the builders before save.
	transaction.SourceAccValidator = transactionDescSourceAcc.Validators[0].(func(string) error)
	// transactionDescDestAcc is the schema descriptor for dest_acc field.
	transactionDescDestAcc := transactionFields[3].Descriptor()
	// transaction.DestAccValidator is a validator for the "dest_acc" field. It is called by the builders before save.
	transaction.DestAccValidator = transactionDescDestAcc.Validators[0].(func(string) error)
}
