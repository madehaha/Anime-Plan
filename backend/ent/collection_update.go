// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/collection"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollectionUpdate is the builder for updating Collection entities.
type CollectionUpdate struct {
	config
	hooks    []Hook
	mutation *CollectionMutation
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cu *CollectionUpdate) Where(ps ...predicate.Collection) *CollectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUID sets the "uid" field.
func (cu *CollectionUpdate) SetUID(u uint32) *CollectionUpdate {
	cu.mutation.ResetUID()
	cu.mutation.SetUID(u)
	return cu
}

// AddUID adds u to the "uid" field.
func (cu *CollectionUpdate) AddUID(u int32) *CollectionUpdate {
	cu.mutation.AddUID(u)
	return cu
}

// SetSubID sets the "sub_id" field.
func (cu *CollectionUpdate) SetSubID(i int64) *CollectionUpdate {
	cu.mutation.ResetSubID()
	cu.mutation.SetSubID(i)
	return cu
}

// AddSubID adds i to the "sub_id" field.
func (cu *CollectionUpdate) AddSubID(i int64) *CollectionUpdate {
	cu.mutation.AddSubID(i)
	return cu
}

// SetType sets the "type" field.
func (cu *CollectionUpdate) SetType(u uint8) *CollectionUpdate {
	cu.mutation.ResetType()
	cu.mutation.SetType(u)
	return cu
}

// AddType adds u to the "type" field.
func (cu *CollectionUpdate) AddType(u int8) *CollectionUpdate {
	cu.mutation.AddType(u)
	return cu
}

// SetIfComment sets the "if_comment" field.
func (cu *CollectionUpdate) SetIfComment(b bool) *CollectionUpdate {
	cu.mutation.SetIfComment(b)
	return cu
}

// SetNillableIfComment sets the "if_comment" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableIfComment(b *bool) *CollectionUpdate {
	if b != nil {
		cu.SetIfComment(*b)
	}
	return cu
}

// SetComment sets the "comment" field.
func (cu *CollectionUpdate) SetComment(s string) *CollectionUpdate {
	cu.mutation.SetComment(s)
	return cu
}

// SetScore sets the "score" field.
func (cu *CollectionUpdate) SetScore(i int8) *CollectionUpdate {
	cu.mutation.ResetScore()
	cu.mutation.SetScore(i)
	return cu
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableScore(i *int8) *CollectionUpdate {
	if i != nil {
		cu.SetScore(*i)
	}
	return cu
}

// AddScore adds i to the "score" field.
func (cu *CollectionUpdate) AddScore(i int8) *CollectionUpdate {
	cu.mutation.AddScore(i)
	return cu
}

// Mutation returns the CollectionMutation object of the builder.
func (cu *CollectionUpdate) Mutation() *CollectionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollectionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CollectionUpdate) check() error {
	if v, ok := cu.mutation.Comment(); ok {
		if err := collection.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "Collection.comment": %w`, err)}
		}
	}
	return nil
}

func (cu *CollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(collection.Table, collection.Columns, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UID(); ok {
		_spec.SetField(collection.FieldUID, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedUID(); ok {
		_spec.AddField(collection.FieldUID, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.SubID(); ok {
		_spec.SetField(collection.FieldSubID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedSubID(); ok {
		_spec.AddField(collection.FieldSubID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(collection.FieldType, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.AddedType(); ok {
		_spec.AddField(collection.FieldType, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.IfComment(); ok {
		_spec.SetField(collection.FieldIfComment, field.TypeBool, value)
	}
	if value, ok := cu.mutation.Comment(); ok {
		_spec.SetField(collection.FieldComment, field.TypeString, value)
	}
	if value, ok := cu.mutation.Score(); ok {
		_spec.SetField(collection.FieldScore, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedScore(); ok {
		_spec.AddField(collection.FieldScore, field.TypeInt8, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CollectionUpdateOne is the builder for updating a single Collection entity.
type CollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollectionMutation
}

// SetUID sets the "uid" field.
func (cuo *CollectionUpdateOne) SetUID(u uint32) *CollectionUpdateOne {
	cuo.mutation.ResetUID()
	cuo.mutation.SetUID(u)
	return cuo
}

// AddUID adds u to the "uid" field.
func (cuo *CollectionUpdateOne) AddUID(u int32) *CollectionUpdateOne {
	cuo.mutation.AddUID(u)
	return cuo
}

// SetSubID sets the "sub_id" field.
func (cuo *CollectionUpdateOne) SetSubID(i int64) *CollectionUpdateOne {
	cuo.mutation.ResetSubID()
	cuo.mutation.SetSubID(i)
	return cuo
}

// AddSubID adds i to the "sub_id" field.
func (cuo *CollectionUpdateOne) AddSubID(i int64) *CollectionUpdateOne {
	cuo.mutation.AddSubID(i)
	return cuo
}

// SetType sets the "type" field.
func (cuo *CollectionUpdateOne) SetType(u uint8) *CollectionUpdateOne {
	cuo.mutation.ResetType()
	cuo.mutation.SetType(u)
	return cuo
}

// AddType adds u to the "type" field.
func (cuo *CollectionUpdateOne) AddType(u int8) *CollectionUpdateOne {
	cuo.mutation.AddType(u)
	return cuo
}

// SetIfComment sets the "if_comment" field.
func (cuo *CollectionUpdateOne) SetIfComment(b bool) *CollectionUpdateOne {
	cuo.mutation.SetIfComment(b)
	return cuo
}

// SetNillableIfComment sets the "if_comment" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableIfComment(b *bool) *CollectionUpdateOne {
	if b != nil {
		cuo.SetIfComment(*b)
	}
	return cuo
}

// SetComment sets the "comment" field.
func (cuo *CollectionUpdateOne) SetComment(s string) *CollectionUpdateOne {
	cuo.mutation.SetComment(s)
	return cuo
}

// SetScore sets the "score" field.
func (cuo *CollectionUpdateOne) SetScore(i int8) *CollectionUpdateOne {
	cuo.mutation.ResetScore()
	cuo.mutation.SetScore(i)
	return cuo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableScore(i *int8) *CollectionUpdateOne {
	if i != nil {
		cuo.SetScore(*i)
	}
	return cuo
}

// AddScore adds i to the "score" field.
func (cuo *CollectionUpdateOne) AddScore(i int8) *CollectionUpdateOne {
	cuo.mutation.AddScore(i)
	return cuo
}

// Mutation returns the CollectionMutation object of the builder.
func (cuo *CollectionUpdateOne) Mutation() *CollectionMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cuo *CollectionUpdateOne) Where(ps ...predicate.Collection) *CollectionUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollectionUpdateOne) Select(field string, fields ...string) *CollectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collection entity.
func (cuo *CollectionUpdateOne) Save(ctx context.Context) (*Collection, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollectionUpdateOne) SaveX(ctx context.Context) *Collection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CollectionUpdateOne) check() error {
	if v, ok := cuo.mutation.Comment(); ok {
		if err := collection.CommentValidator(v); err != nil {
			return &ValidationError{Name: "comment", err: fmt.Errorf(`ent: validator failed for field "Collection.comment": %w`, err)}
		}
	}
	return nil
}

func (cuo *CollectionUpdateOne) sqlSave(ctx context.Context) (_node *Collection, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(collection.Table, collection.Columns, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Collection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collection.FieldID)
		for _, f := range fields {
			if !collection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != collection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UID(); ok {
		_spec.SetField(collection.FieldUID, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedUID(); ok {
		_spec.AddField(collection.FieldUID, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.SubID(); ok {
		_spec.SetField(collection.FieldSubID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedSubID(); ok {
		_spec.AddField(collection.FieldSubID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(collection.FieldType, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.AddedType(); ok {
		_spec.AddField(collection.FieldType, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.IfComment(); ok {
		_spec.SetField(collection.FieldIfComment, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.Comment(); ok {
		_spec.SetField(collection.FieldComment, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Score(); ok {
		_spec.SetField(collection.FieldScore, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedScore(); ok {
		_spec.AddField(collection.FieldScore, field.TypeInt8, value)
	}
	_node = &Collection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
