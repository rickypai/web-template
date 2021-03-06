// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rickypai/web-template/api/ent/operatingsystem"
	"github.com/rickypai/web-template/api/ent/phone"
)

// OperatingSystemCreate is the builder for creating a OperatingSystem entity.
type OperatingSystemCreate struct {
	config
	mutation *OperatingSystemMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (osc *OperatingSystemCreate) SetName(s string) *OperatingSystemCreate {
	osc.mutation.SetName(s)
	return osc
}

// SetCreatedAt sets the "created_at" field.
func (osc *OperatingSystemCreate) SetCreatedAt(t time.Time) *OperatingSystemCreate {
	osc.mutation.SetCreatedAt(t)
	return osc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableCreatedAt(t *time.Time) *OperatingSystemCreate {
	if t != nil {
		osc.SetCreatedAt(*t)
	}
	return osc
}

// SetModifiedAt sets the "modified_at" field.
func (osc *OperatingSystemCreate) SetModifiedAt(t time.Time) *OperatingSystemCreate {
	osc.mutation.SetModifiedAt(t)
	return osc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (osc *OperatingSystemCreate) SetNillableModifiedAt(t *time.Time) *OperatingSystemCreate {
	if t != nil {
		osc.SetModifiedAt(*t)
	}
	return osc
}

// AddPhoneIDs adds the "phones" edge to the Phone entity by IDs.
func (osc *OperatingSystemCreate) AddPhoneIDs(ids ...int) *OperatingSystemCreate {
	osc.mutation.AddPhoneIDs(ids...)
	return osc
}

// AddPhones adds the "phones" edges to the Phone entity.
func (osc *OperatingSystemCreate) AddPhones(p ...*Phone) *OperatingSystemCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return osc.AddPhoneIDs(ids...)
}

// Mutation returns the OperatingSystemMutation object of the builder.
func (osc *OperatingSystemCreate) Mutation() *OperatingSystemMutation {
	return osc.mutation
}

// Save creates the OperatingSystem in the database.
func (osc *OperatingSystemCreate) Save(ctx context.Context) (*OperatingSystem, error) {
	var (
		err  error
		node *OperatingSystem
	)
	osc.defaults()
	if len(osc.hooks) == 0 {
		if err = osc.check(); err != nil {
			return nil, err
		}
		node, err = osc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperatingSystemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osc.check(); err != nil {
				return nil, err
			}
			osc.mutation = mutation
			node, err = osc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osc.hooks) - 1; i >= 0; i-- {
			mut = osc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OperatingSystemCreate) SaveX(ctx context.Context) *OperatingSystem {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (osc *OperatingSystemCreate) defaults() {
	if _, ok := osc.mutation.CreatedAt(); !ok {
		v := operatingsystem.DefaultCreatedAt()
		osc.mutation.SetCreatedAt(v)
	}
	if _, ok := osc.mutation.ModifiedAt(); !ok {
		v := operatingsystem.DefaultModifiedAt()
		osc.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OperatingSystemCreate) check() error {
	if _, ok := osc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := osc.mutation.Name(); ok {
		if err := operatingsystem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := osc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := osc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New("ent: missing required field \"modified_at\"")}
	}
	return nil
}

func (osc *OperatingSystemCreate) sqlSave(ctx context.Context) (*OperatingSystem, error) {
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (osc *OperatingSystemCreate) createSpec() (*OperatingSystem, *sqlgraph.CreateSpec) {
	var (
		_node = &OperatingSystem{config: osc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: operatingsystem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operatingsystem.FieldID,
			},
		}
	)
	if value, ok := osc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operatingsystem.FieldName,
		})
		_node.Name = value
	}
	if value, ok := osc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operatingsystem.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := osc.mutation.ModifiedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operatingsystem.FieldModifiedAt,
		})
		_node.ModifiedAt = value
	}
	if nodes := osc.mutation.PhonesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   operatingsystem.PhonesTable,
			Columns: []string{operatingsystem.PhonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: phone.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OperatingSystemCreateBulk is the builder for creating many OperatingSystem entities in bulk.
type OperatingSystemCreateBulk struct {
	config
	builders []*OperatingSystemCreate
}

// Save creates the OperatingSystem entities in the database.
func (oscb *OperatingSystemCreateBulk) Save(ctx context.Context) ([]*OperatingSystem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OperatingSystem, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperatingSystemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OperatingSystemCreateBulk) SaveX(ctx context.Context) []*OperatingSystem {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
