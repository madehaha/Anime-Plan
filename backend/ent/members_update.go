// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/collection"
	"backend/ent/members"
	"backend/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MembersUpdate is the builder for updating Members entities.
type MembersUpdate struct {
	config
	hooks    []Hook
	mutation *MembersMutation
}

// Where appends a list predicates to the MembersUpdate builder.
func (mu *MembersUpdate) Where(ps ...predicate.Members) *MembersUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUsername sets the "username" field.
func (mu *MembersUpdate) SetUsername(s string) *MembersUpdate {
	mu.mutation.SetUsername(s)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MembersUpdate) SetEmail(s string) *MembersUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetPassword sets the "password" field.
func (mu *MembersUpdate) SetPassword(s string) *MembersUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MembersUpdate) SetNickname(s string) *MembersUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MembersUpdate) SetAvatar(s string) *MembersUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MembersUpdate) SetNillableAvatar(s *string) *MembersUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// SetGid sets the "gid" field.
func (mu *MembersUpdate) SetGid(u uint8) *MembersUpdate {
	mu.mutation.ResetGid()
	mu.mutation.SetGid(u)
	return mu
}

// SetNillableGid sets the "gid" field if the given value is not nil.
func (mu *MembersUpdate) SetNillableGid(u *uint8) *MembersUpdate {
	if u != nil {
		mu.SetGid(*u)
	}
	return mu
}

// AddGid adds u to the "gid" field.
func (mu *MembersUpdate) AddGid(u int8) *MembersUpdate {
	mu.mutation.AddGid(u)
	return mu
}

// AddCollectionIDs adds the "collections" edge to the Collection entity by IDs.
func (mu *MembersUpdate) AddCollectionIDs(ids ...uint32) *MembersUpdate {
	mu.mutation.AddCollectionIDs(ids...)
	return mu
}

// AddCollections adds the "collections" edges to the Collection entity.
func (mu *MembersUpdate) AddCollections(c ...*Collection) *MembersUpdate {
	ids := make([]uint32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.AddCollectionIDs(ids...)
}

// Mutation returns the MembersMutation object of the builder.
func (mu *MembersUpdate) Mutation() *MembersMutation {
	return mu.mutation
}

// ClearCollections clears all "collections" edges to the Collection entity.
func (mu *MembersUpdate) ClearCollections() *MembersUpdate {
	mu.mutation.ClearCollections()
	return mu
}

// RemoveCollectionIDs removes the "collections" edge to Collection entities by IDs.
func (mu *MembersUpdate) RemoveCollectionIDs(ids ...uint32) *MembersUpdate {
	mu.mutation.RemoveCollectionIDs(ids...)
	return mu
}

// RemoveCollections removes "collections" edges to Collection entities.
func (mu *MembersUpdate) RemoveCollections(c ...*Collection) *MembersUpdate {
	ids := make([]uint32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.RemoveCollectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MembersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MembersUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MembersUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MembersUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MembersUpdate) check() error {
	if v, ok := mu.mutation.Username(); ok {
		if err := members.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Members.username": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Email(); ok {
		if err := members.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Members.email": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Nickname(); ok {
		if err := members.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Members.nickname": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Avatar(); ok {
		if err := members.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Members.avatar": %w`, err)}
		}
	}
	return nil
}

func (mu *MembersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(members.Table, members.Columns, sqlgraph.NewFieldSpec(members.FieldID, field.TypeUint32))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Username(); ok {
		_spec.SetField(members.FieldUsername, field.TypeString, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(members.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(members.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.SetField(members.FieldNickname, field.TypeString, value)
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.SetField(members.FieldAvatar, field.TypeString, value)
	}
	if value, ok := mu.mutation.Gid(); ok {
		_spec.SetField(members.FieldGid, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.AddedGid(); ok {
		_spec.AddField(members.FieldGid, field.TypeUint8, value)
	}
	if mu.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedCollectionsIDs(); len(nodes) > 0 && !mu.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CollectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{members.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MembersUpdateOne is the builder for updating a single Members entity.
type MembersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MembersMutation
}

// SetUsername sets the "username" field.
func (muo *MembersUpdateOne) SetUsername(s string) *MembersUpdateOne {
	muo.mutation.SetUsername(s)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MembersUpdateOne) SetEmail(s string) *MembersUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetPassword sets the "password" field.
func (muo *MembersUpdateOne) SetPassword(s string) *MembersUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MembersUpdateOne) SetNickname(s string) *MembersUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MembersUpdateOne) SetAvatar(s string) *MembersUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MembersUpdateOne) SetNillableAvatar(s *string) *MembersUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// SetGid sets the "gid" field.
func (muo *MembersUpdateOne) SetGid(u uint8) *MembersUpdateOne {
	muo.mutation.ResetGid()
	muo.mutation.SetGid(u)
	return muo
}

// SetNillableGid sets the "gid" field if the given value is not nil.
func (muo *MembersUpdateOne) SetNillableGid(u *uint8) *MembersUpdateOne {
	if u != nil {
		muo.SetGid(*u)
	}
	return muo
}

// AddGid adds u to the "gid" field.
func (muo *MembersUpdateOne) AddGid(u int8) *MembersUpdateOne {
	muo.mutation.AddGid(u)
	return muo
}

// AddCollectionIDs adds the "collections" edge to the Collection entity by IDs.
func (muo *MembersUpdateOne) AddCollectionIDs(ids ...uint32) *MembersUpdateOne {
	muo.mutation.AddCollectionIDs(ids...)
	return muo
}

// AddCollections adds the "collections" edges to the Collection entity.
func (muo *MembersUpdateOne) AddCollections(c ...*Collection) *MembersUpdateOne {
	ids := make([]uint32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.AddCollectionIDs(ids...)
}

// Mutation returns the MembersMutation object of the builder.
func (muo *MembersUpdateOne) Mutation() *MembersMutation {
	return muo.mutation
}

// ClearCollections clears all "collections" edges to the Collection entity.
func (muo *MembersUpdateOne) ClearCollections() *MembersUpdateOne {
	muo.mutation.ClearCollections()
	return muo
}

// RemoveCollectionIDs removes the "collections" edge to Collection entities by IDs.
func (muo *MembersUpdateOne) RemoveCollectionIDs(ids ...uint32) *MembersUpdateOne {
	muo.mutation.RemoveCollectionIDs(ids...)
	return muo
}

// RemoveCollections removes "collections" edges to Collection entities.
func (muo *MembersUpdateOne) RemoveCollections(c ...*Collection) *MembersUpdateOne {
	ids := make([]uint32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.RemoveCollectionIDs(ids...)
}

// Where appends a list predicates to the MembersUpdate builder.
func (muo *MembersUpdateOne) Where(ps ...predicate.Members) *MembersUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MembersUpdateOne) Select(field string, fields ...string) *MembersUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Members entity.
func (muo *MembersUpdateOne) Save(ctx context.Context) (*Members, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MembersUpdateOne) SaveX(ctx context.Context) *Members {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MembersUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MembersUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MembersUpdateOne) check() error {
	if v, ok := muo.mutation.Username(); ok {
		if err := members.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Members.username": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Email(); ok {
		if err := members.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Members.email": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Nickname(); ok {
		if err := members.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Members.nickname": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Avatar(); ok {
		if err := members.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Members.avatar": %w`, err)}
		}
	}
	return nil
}

func (muo *MembersUpdateOne) sqlSave(ctx context.Context) (_node *Members, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(members.Table, members.Columns, sqlgraph.NewFieldSpec(members.FieldID, field.TypeUint32))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Members.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, members.FieldID)
		for _, f := range fields {
			if !members.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != members.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Username(); ok {
		_spec.SetField(members.FieldUsername, field.TypeString, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(members.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(members.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.SetField(members.FieldNickname, field.TypeString, value)
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.SetField(members.FieldAvatar, field.TypeString, value)
	}
	if value, ok := muo.mutation.Gid(); ok {
		_spec.SetField(members.FieldGid, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.AddedGid(); ok {
		_spec.AddField(members.FieldGid, field.TypeUint8, value)
	}
	if muo.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedCollectionsIDs(); len(nodes) > 0 && !muo.mutation.CollectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CollectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   members.CollectionsTable,
			Columns: []string{members.CollectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Members{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{members.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
