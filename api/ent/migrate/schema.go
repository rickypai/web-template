// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ManufacturersColumns holds the columns for the "manufacturers" table.
	ManufacturersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "manufacturer_phones", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ManufacturersTable holds the schema information for the "manufacturers" table.
	ManufacturersTable = &schema.Table{
		Name:       "manufacturers",
		Columns:    ManufacturersColumns,
		PrimaryKey: []*schema.Column{ManufacturersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "manufacturers_manufacturers_phones",
				Columns:    []*schema.Column{ManufacturersColumns[4]},
				RefColumns: []*schema.Column{ManufacturersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OperatingSystemsColumns holds the columns for the "operating_systems" table.
	OperatingSystemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "operating_system_phones", Type: field.TypeInt, Nullable: true},
	}
	// OperatingSystemsTable holds the schema information for the "operating_systems" table.
	OperatingSystemsTable = &schema.Table{
		Name:       "operating_systems",
		Columns:    OperatingSystemsColumns,
		PrimaryKey: []*schema.Column{OperatingSystemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "operating_systems_manufacturers_phones",
				Columns:    []*schema.Column{OperatingSystemsColumns[4]},
				RefColumns: []*schema.Column{ManufacturersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhonesColumns holds the columns for the "phones" table.
	PhonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "modified_at", Type: field.TypeTime},
		{Name: "phone_manufacturer", Type: field.TypeInt, Nullable: true},
		{Name: "phone_os", Type: field.TypeInt, Nullable: true},
	}
	// PhonesTable holds the schema information for the "phones" table.
	PhonesTable = &schema.Table{
		Name:       "phones",
		Columns:    PhonesColumns,
		PrimaryKey: []*schema.Column{PhonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "phones_manufacturers_manufacturer",
				Columns:    []*schema.Column{PhonesColumns[4]},
				RefColumns: []*schema.Column{ManufacturersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "phones_operating_systems_os",
				Columns:    []*schema.Column{PhonesColumns[5]},
				RefColumns: []*schema.Column{OperatingSystemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ManufacturersTable,
		OperatingSystemsTable,
		PhonesTable,
	}
)

func init() {
	ManufacturersTable.ForeignKeys[0].RefTable = ManufacturersTable
	OperatingSystemsTable.ForeignKeys[0].RefTable = ManufacturersTable
	PhonesTable.ForeignKeys[0].RefTable = ManufacturersTable
	PhonesTable.ForeignKeys[1].RefTable = OperatingSystemsTable
}
