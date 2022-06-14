// Code generated by entc, DO NOT EDIT.

package teller

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the teller type in the database.
	Label = "teller"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// Table holds the table name of the teller in the database.
	Table = "tellers"
)

// Columns holds all SQL columns for teller fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUsername,
	FieldPassword,
	FieldRole,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// Role defines the type for the "role" enum field.
type Role string

// RoleCustomer is the default value of the Role enum.
const DefaultRole = RoleCustomer

// Role values.
const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleCustomer, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("teller: invalid enum value for role field: %q", r)
	}
}
