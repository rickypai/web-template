// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rickypai/web-template/api/ent/manufacturer"
	"github.com/rickypai/web-template/api/ent/operatingsystem"
	"github.com/rickypai/web-template/api/ent/phone"
)

// Phone is the model entity for the Phone schema.
type Phone struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhoneQuery when eager-loading is set.
	Edges               PhoneEdges `json:"edges"`
	manufacturer_id     *int
	operating_system_id *int
}

// PhoneEdges holds the relations/edges for other nodes in the graph.
type PhoneEdges struct {
	// Manufacturer holds the value of the manufacturer edge.
	Manufacturer *Manufacturer `json:"manufacturer,omitempty"`
	// OperatingSystem holds the value of the operating_system edge.
	OperatingSystem *OperatingSystem `json:"operating_system,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ManufacturerOrErr returns the Manufacturer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneEdges) ManufacturerOrErr() (*Manufacturer, error) {
	if e.loadedTypes[0] {
		if e.Manufacturer == nil {
			// The edge manufacturer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: manufacturer.Label}
		}
		return e.Manufacturer, nil
	}
	return nil, &NotLoadedError{edge: "manufacturer"}
}

// OperatingSystemOrErr returns the OperatingSystem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhoneEdges) OperatingSystemOrErr() (*OperatingSystem, error) {
	if e.loadedTypes[1] {
		if e.OperatingSystem == nil {
			// The edge operating_system was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: operatingsystem.Label}
		}
		return e.OperatingSystem, nil
	}
	return nil, &NotLoadedError{edge: "operating_system"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Phone) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case phone.FieldID:
			values[i] = &sql.NullInt64{}
		case phone.FieldName:
			values[i] = &sql.NullString{}
		case phone.FieldCreatedAt, phone.FieldModifiedAt:
			values[i] = &sql.NullTime{}
		case phone.ForeignKeys[0]: // manufacturer_id
			values[i] = &sql.NullInt64{}
		case phone.ForeignKeys[1]: // operating_system_id
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Phone", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Phone fields.
func (ph *Phone) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case phone.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ph.ID = int(value.Int64)
		case phone.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ph.Name = value.String
			}
		case phone.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ph.CreatedAt = value.Time
			}
		case phone.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				ph.ModifiedAt = value.Time
			}
		case phone.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field manufacturer_id", value)
			} else if value.Valid {
				ph.manufacturer_id = new(int)
				*ph.manufacturer_id = int(value.Int64)
			}
		case phone.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field operating_system_id", value)
			} else if value.Valid {
				ph.operating_system_id = new(int)
				*ph.operating_system_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryManufacturer queries the "manufacturer" edge of the Phone entity.
func (ph *Phone) QueryManufacturer() *ManufacturerQuery {
	return (&PhoneClient{config: ph.config}).QueryManufacturer(ph)
}

// QueryOperatingSystem queries the "operating_system" edge of the Phone entity.
func (ph *Phone) QueryOperatingSystem() *OperatingSystemQuery {
	return (&PhoneClient{config: ph.config}).QueryOperatingSystem(ph)
}

// Update returns a builder for updating this Phone.
// Note that you need to call Phone.Unwrap() before calling this method if this Phone
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *Phone) Update() *PhoneUpdateOne {
	return (&PhoneClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the Phone entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *Phone) Unwrap() *Phone {
	tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: Phone is not a transactional entity")
	}
	ph.config.driver = tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *Phone) String() string {
	var builder strings.Builder
	builder.WriteString("Phone(")
	builder.WriteString(fmt.Sprintf("id=%v", ph.ID))
	builder.WriteString(", name=")
	builder.WriteString(ph.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(ph.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", modified_at=")
	builder.WriteString(ph.ModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Phones is a parsable slice of Phone.
type Phones []*Phone

func (ph Phones) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
