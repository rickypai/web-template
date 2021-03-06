// Code generated by entc, DO NOT EDIT.

package phone

import (
	"time"
)

const (
	// Label holds the string label denoting the phone type in the database.
	Label = "phone"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// EdgeManufacturer holds the string denoting the manufacturer edge name in mutations.
	EdgeManufacturer = "manufacturer"
	// EdgeOperatingSystem holds the string denoting the operating_system edge name in mutations.
	EdgeOperatingSystem = "operating_system"
	// Table holds the table name of the phone in the database.
	Table = "phones"
	// ManufacturerTable is the table the holds the manufacturer relation/edge.
	ManufacturerTable = "phones"
	// ManufacturerInverseTable is the table name for the Manufacturer entity.
	// It exists in this package in order to avoid circular dependency with the "manufacturer" package.
	ManufacturerInverseTable = "manufacturers"
	// ManufacturerColumn is the table column denoting the manufacturer relation/edge.
	ManufacturerColumn = "manufacturer_id"
	// OperatingSystemTable is the table the holds the operating_system relation/edge.
	OperatingSystemTable = "phones"
	// OperatingSystemInverseTable is the table name for the OperatingSystem entity.
	// It exists in this package in order to avoid circular dependency with the "operatingsystem" package.
	OperatingSystemInverseTable = "operating_systems"
	// OperatingSystemColumn is the table column denoting the operating_system relation/edge.
	OperatingSystemColumn = "operating_system_id"
)

// Columns holds all SQL columns for phone fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "phones"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"manufacturer_id",
	"operating_system_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
)
