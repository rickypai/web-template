// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rickypai/web-template/api/ent/manufacturer"
	"github.com/rickypai/web-template/api/ent/operatingsystem"
	"github.com/rickypai/web-template/api/ent/phone"
	"github.com/rickypai/web-template/api/ent/predicate"
)

// PhoneUpdate is the builder for updating Phone entities.
type PhoneUpdate struct {
	config
	hooks    []Hook
	mutation *PhoneMutation
}

// Where adds a new predicate for the PhoneUpdate builder.
func (pu *PhoneUpdate) Where(ps ...predicate.Phone) *PhoneUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PhoneUpdate) SetName(s string) *PhoneUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetManufacturerID sets the "manufacturer" edge to the Manufacturer entity by ID.
func (pu *PhoneUpdate) SetManufacturerID(id int) *PhoneUpdate {
	pu.mutation.SetManufacturerID(id)
	return pu
}

// SetManufacturer sets the "manufacturer" edge to the Manufacturer entity.
func (pu *PhoneUpdate) SetManufacturer(m *Manufacturer) *PhoneUpdate {
	return pu.SetManufacturerID(m.ID)
}

// SetOperatingSystemID sets the "operating_system" edge to the OperatingSystem entity by ID.
func (pu *PhoneUpdate) SetOperatingSystemID(id int) *PhoneUpdate {
	pu.mutation.SetOperatingSystemID(id)
	return pu
}

// SetOperatingSystem sets the "operating_system" edge to the OperatingSystem entity.
func (pu *PhoneUpdate) SetOperatingSystem(o *OperatingSystem) *PhoneUpdate {
	return pu.SetOperatingSystemID(o.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (pu *PhoneUpdate) Mutation() *PhoneMutation {
	return pu.mutation
}

// ClearManufacturer clears the "manufacturer" edge to the Manufacturer entity.
func (pu *PhoneUpdate) ClearManufacturer() *PhoneUpdate {
	pu.mutation.ClearManufacturer()
	return pu
}

// ClearOperatingSystem clears the "operating_system" edge to the OperatingSystem entity.
func (pu *PhoneUpdate) ClearOperatingSystem() *PhoneUpdate {
	pu.mutation.ClearOperatingSystem()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PhoneUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PhoneUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PhoneUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PhoneUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PhoneUpdate) defaults() {
	if _, ok := pu.mutation.ModifiedAt(); !ok {
		v := phone.UpdateDefaultModifiedAt()
		pu.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PhoneUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := phone.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pu.mutation.ManufacturerID(); pu.mutation.ManufacturerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"manufacturer\"")
	}
	if _, ok := pu.mutation.OperatingSystemID(); pu.mutation.OperatingSystemCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"operating_system\"")
	}
	return nil
}

func (pu *PhoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   phone.Table,
			Columns: phone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: phone.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: phone.FieldName,
		})
	}
	if value, ok := pu.mutation.ModifiedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: phone.FieldModifiedAt,
		})
	}
	if pu.mutation.ManufacturerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.ManufacturerTable,
			Columns: []string{phone.ManufacturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manufacturer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ManufacturerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.ManufacturerTable,
			Columns: []string{phone.ManufacturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manufacturer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OperatingSystemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.OperatingSystemTable,
			Columns: []string{phone.OperatingSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operatingsystem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OperatingSystemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.OperatingSystemTable,
			Columns: []string{phone.OperatingSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operatingsystem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PhoneUpdateOne is the builder for updating a single Phone entity.
type PhoneUpdateOne struct {
	config
	hooks    []Hook
	mutation *PhoneMutation
}

// SetName sets the "name" field.
func (puo *PhoneUpdateOne) SetName(s string) *PhoneUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetManufacturerID sets the "manufacturer" edge to the Manufacturer entity by ID.
func (puo *PhoneUpdateOne) SetManufacturerID(id int) *PhoneUpdateOne {
	puo.mutation.SetManufacturerID(id)
	return puo
}

// SetManufacturer sets the "manufacturer" edge to the Manufacturer entity.
func (puo *PhoneUpdateOne) SetManufacturer(m *Manufacturer) *PhoneUpdateOne {
	return puo.SetManufacturerID(m.ID)
}

// SetOperatingSystemID sets the "operating_system" edge to the OperatingSystem entity by ID.
func (puo *PhoneUpdateOne) SetOperatingSystemID(id int) *PhoneUpdateOne {
	puo.mutation.SetOperatingSystemID(id)
	return puo
}

// SetOperatingSystem sets the "operating_system" edge to the OperatingSystem entity.
func (puo *PhoneUpdateOne) SetOperatingSystem(o *OperatingSystem) *PhoneUpdateOne {
	return puo.SetOperatingSystemID(o.ID)
}

// Mutation returns the PhoneMutation object of the builder.
func (puo *PhoneUpdateOne) Mutation() *PhoneMutation {
	return puo.mutation
}

// ClearManufacturer clears the "manufacturer" edge to the Manufacturer entity.
func (puo *PhoneUpdateOne) ClearManufacturer() *PhoneUpdateOne {
	puo.mutation.ClearManufacturer()
	return puo
}

// ClearOperatingSystem clears the "operating_system" edge to the OperatingSystem entity.
func (puo *PhoneUpdateOne) ClearOperatingSystem() *PhoneUpdateOne {
	puo.mutation.ClearOperatingSystem()
	return puo
}

// Save executes the query and returns the updated Phone entity.
func (puo *PhoneUpdateOne) Save(ctx context.Context) (*Phone, error) {
	var (
		err  error
		node *Phone
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PhoneUpdateOne) SaveX(ctx context.Context) *Phone {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PhoneUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PhoneUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PhoneUpdateOne) defaults() {
	if _, ok := puo.mutation.ModifiedAt(); !ok {
		v := phone.UpdateDefaultModifiedAt()
		puo.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PhoneUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := phone.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := puo.mutation.ManufacturerID(); puo.mutation.ManufacturerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"manufacturer\"")
	}
	if _, ok := puo.mutation.OperatingSystemID(); puo.mutation.OperatingSystemCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"operating_system\"")
	}
	return nil
}

func (puo *PhoneUpdateOne) sqlSave(ctx context.Context) (_node *Phone, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   phone.Table,
			Columns: phone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: phone.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Phone.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: phone.FieldName,
		})
	}
	if value, ok := puo.mutation.ModifiedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: phone.FieldModifiedAt,
		})
	}
	if puo.mutation.ManufacturerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.ManufacturerTable,
			Columns: []string{phone.ManufacturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manufacturer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ManufacturerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.ManufacturerTable,
			Columns: []string{phone.ManufacturerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manufacturer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OperatingSystemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.OperatingSystemTable,
			Columns: []string{phone.OperatingSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operatingsystem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OperatingSystemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   phone.OperatingSystemTable,
			Columns: []string{phone.OperatingSystemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operatingsystem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Phone{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phone.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
