// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent/schema"
	"github.com/Faroukhamadi/Banketeer/transaction_service/transactiondb/ent/transaction"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	transactionMixin := schema.Transaction{}.Mixin()
	transactionMixinFields0 := transactionMixin[0].Fields()
	_ = transactionMixinFields0
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCreateTime is the schema descriptor for create_time field.
	transactionDescCreateTime := transactionMixinFields0[0].Descriptor()
	// transaction.DefaultCreateTime holds the default value on creation for the create_time field.
	transaction.DefaultCreateTime = transactionDescCreateTime.Default.(func() time.Time)
	// transactionDescUpdateTime is the schema descriptor for update_time field.
	transactionDescUpdateTime := transactionMixinFields0[1].Descriptor()
	// transaction.DefaultUpdateTime holds the default value on creation for the update_time field.
	transaction.DefaultUpdateTime = transactionDescUpdateTime.Default.(func() time.Time)
	// transaction.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	transaction.UpdateDefaultUpdateTime = transactionDescUpdateTime.UpdateDefault.(func() time.Time)
}
