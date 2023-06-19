// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/subject"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollectionCreate is the builder for creating a Collection entity.
type CollectionCreate struct {
	config
	mutation *CollectionMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (cc *CollectionCreate) SetType(u uint8) *CollectionCreate {
	cc.mutation.SetType(u)
	return cc
}

// SetHasComment sets the "has_comment" field.
func (cc *CollectionCreate) SetHasComment(b bool) *CollectionCreate {
	cc.mutation.SetHasComment(b)
	return cc
}

// SetNillableHasComment sets the "has_comment" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableHasComment(b *bool) *CollectionCreate {
	if b != nil {
		cc.SetHasComment(*b)
	}
	return cc
}

// SetComment sets the "comment" field.
func (cc *CollectionCreate) SetComment(s string) *CollectionCreate {
	cc.mutation.SetComment(s)
	return cc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableComment(s *string) *CollectionCreate {
	if s != nil {
		cc.SetComment(*s)
	}
	return cc
}

// SetScore sets the "score" field.
func (cc *CollectionCreate) SetScore(u uint8) *CollectionCreate {
	cc.mutation.SetScore(u)
	return cc
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableScore(u *uint8) *CollectionCreate {
	if u != nil {
		cc.SetScore(*u)
	}
	return cc
}

// SetAddTime sets the "add_time" field.
func (cc *CollectionCreate) SetAddTime(s string) *CollectionCreate {
	cc.mutation.SetAddTime(s)
	return cc
}

// SetNillableAddTime sets the "add_time" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableAddTime(s *string) *CollectionCreate {
	if s != nil {
		cc.SetAddTime(*s)
	}
	return cc
}

// SetEpStatus sets the "ep_status" field.
func (cc *CollectionCreate) SetEpStatus(u uint8) *CollectionCreate {
	cc.mutation.SetEpStatus(u)
	return cc
}

// SetNillableEpStatus sets the "ep_status" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableEpStatus(u *uint8) *CollectionCreate {
	if u != nil {
		cc.SetEpStatus(*u)
	}
	return cc
}

// SetMemberID sets the "member_id" field.
func (cc *CollectionCreate) SetMemberID(u uint32) *CollectionCreate {
	cc.mutation.SetMemberID(u)
	return cc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableMemberID(u *uint32) *CollectionCreate {
	if u != nil {
		cc.SetMemberID(*u)
	}
	return cc
}

// SetSubjectID sets the "subject_id" field.
func (cc *CollectionCreate) SetSubjectID(u uint32) *CollectionCreate {
	cc.mutation.SetSubjectID(u)
	return cc
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (cc *CollectionCreate) SetNillableSubjectID(u *uint32) *CollectionCreate {
	if u != nil {
		cc.SetSubjectID(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CollectionCreate) SetID(u uint32) *CollectionCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetMember sets the "member" edge to the Members entity.
func (cc *CollectionCreate) SetMember(m *Members) *CollectionCreate {
	return cc.SetMemberID(m.ID)
}

// SetSubject sets the "subject" edge to the Subject entity.
func (cc *CollectionCreate) SetSubject(s *Subject) *CollectionCreate {
	return cc.SetSubjectID(s.ID)
}

// Mutation returns the CollectionMutation object of the builder.
func (cc *CollectionCreate) Mutation() *CollectionMutation {
	return cc.mutation
}

// Save creates the Collection in the database.
func (cc *CollectionCreate) Save(ctx context.Context) (*Collection, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollectionCreate) SaveX(ctx context.Context) *Collection {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollectionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollectionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CollectionCreate) defaults() {
	if _, ok := cc.mutation.HasComment(); !ok {
		v := collection.DefaultHasComment
		cc.mutation.SetHasComment(v)
	}
	if _, ok := cc.mutation.Comment(); !ok {
		v := collection.DefaultComment
		cc.mutation.SetComment(v)
	}
	if _, ok := cc.mutation.Score(); !ok {
		v := collection.DefaultScore
		cc.mutation.SetScore(v)
	}
	if _, ok := cc.mutation.AddTime(); !ok {
		v := collection.DefaultAddTime
		cc.mutation.SetAddTime(v)
	}
	if _, ok := cc.mutation.EpStatus(); !ok {
		v := collection.DefaultEpStatus
		cc.mutation.SetEpStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollectionCreate) check() error {
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Collection.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := collection.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Collection.type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.HasComment(); !ok {
		return &ValidationError{Name: "has_comment", err: errors.New(`ent: missing required field "Collection.has_comment"`)}
	}
	if _, ok := cc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Collection.comment"`)}
	}
	if v, ok := cc.mutation.Comment(); ok {
		if err := collection.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "Collection.comment": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Collection.score"`)}
	}
	if _, ok := cc.mutation.AddTime(); !ok {
		return &ValidationError{Name: "add_time", err: errors.New(`ent: missing required field "Collection.add_time"`)}
	}
	if _, ok := cc.mutation.EpStatus(); !ok {
		return &ValidationError{Name: "ep_status", err: errors.New(`ent: missing required field "Collection.ep_status"`)}
	}
	return nil
}

func (cc *CollectionCreate) sqlSave(ctx context.Context) (*Collection, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollectionCreate) createSpec() (*Collection, *sqlgraph.CreateSpec) {
	var (
		_node = &Collection{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(collection.Table, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(collection.FieldType, field.TypeUint8, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.HasComment(); ok {
		_spec.SetField(collection.FieldHasComment, field.TypeBool, value)
		_node.HasComment = value
	}
	if value, ok := cc.mutation.Comment(); ok {
		_spec.SetField(collection.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := cc.mutation.Score(); ok {
		_spec.SetField(collection.FieldScore, field.TypeUint8, value)
		_node.Score = value
	}
	if value, ok := cc.mutation.AddTime(); ok {
		_spec.SetField(collection.FieldAddTime, field.TypeString, value)
		_node.AddTime = value
	}
	if value, ok := cc.mutation.EpStatus(); ok {
		_spec.SetField(collection.FieldEpStatus, field.TypeUint8, value)
		_node.EpStatus = value
	}
	if nodes := cc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.MemberTable,
			Columns: []string{collection.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(members.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   collection.SubjectTable,
			Columns: []string{collection.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CollectionCreateBulk is the builder for creating many Collection entities in bulk.
type CollectionCreateBulk struct {
	config
	builders []*CollectionCreate
}

// Save creates the Collection entities in the database.
func (ccb *CollectionCreateBulk) Save(ctx context.Context) ([]*Collection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Collection, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollectionCreateBulk) SaveX(ctx context.Context) []*Collection {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollectionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
