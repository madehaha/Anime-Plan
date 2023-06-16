// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/subject"
	"backend/ent/subjectfield"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubjectFieldQuery is the builder for querying SubjectField entities.
type SubjectFieldQuery struct {
	config
	ctx         *QueryContext
	order       []subjectfield.OrderOption
	inters      []Interceptor
	predicates  []predicate.SubjectField
	withSubject *SubjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubjectFieldQuery builder.
func (sfq *SubjectFieldQuery) Where(ps ...predicate.SubjectField) *SubjectFieldQuery {
	sfq.predicates = append(sfq.predicates, ps...)
	return sfq
}

// Limit the number of records to be returned by this query.
func (sfq *SubjectFieldQuery) Limit(limit int) *SubjectFieldQuery {
	sfq.ctx.Limit = &limit
	return sfq
}

// Offset to start from.
func (sfq *SubjectFieldQuery) Offset(offset int) *SubjectFieldQuery {
	sfq.ctx.Offset = &offset
	return sfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sfq *SubjectFieldQuery) Unique(unique bool) *SubjectFieldQuery {
	sfq.ctx.Unique = &unique
	return sfq
}

// Order specifies how the records should be ordered.
func (sfq *SubjectFieldQuery) Order(o ...subjectfield.OrderOption) *SubjectFieldQuery {
	sfq.order = append(sfq.order, o...)
	return sfq
}

// QuerySubject chains the current query on the "subject" edge.
func (sfq *SubjectFieldQuery) QuerySubject() *SubjectQuery {
	query := (&SubjectClient{config: sfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subjectfield.Table, subjectfield.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, subjectfield.SubjectTable, subjectfield.SubjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(sfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SubjectField entity from the query.
// Returns a *NotFoundError when no SubjectField was found.
func (sfq *SubjectFieldQuery) First(ctx context.Context) (*SubjectField, error) {
	nodes, err := sfq.Limit(1).All(setContextOp(ctx, sfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subjectfield.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sfq *SubjectFieldQuery) FirstX(ctx context.Context) *SubjectField {
	node, err := sfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SubjectField ID from the query.
// Returns a *NotFoundError when no SubjectField ID was found.
func (sfq *SubjectFieldQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sfq.Limit(1).IDs(setContextOp(ctx, sfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subjectfield.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sfq *SubjectFieldQuery) FirstIDX(ctx context.Context) int {
	id, err := sfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SubjectField entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SubjectField entity is found.
// Returns a *NotFoundError when no SubjectField entities are found.
func (sfq *SubjectFieldQuery) Only(ctx context.Context) (*SubjectField, error) {
	nodes, err := sfq.Limit(2).All(setContextOp(ctx, sfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subjectfield.Label}
	default:
		return nil, &NotSingularError{subjectfield.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sfq *SubjectFieldQuery) OnlyX(ctx context.Context) *SubjectField {
	node, err := sfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SubjectField ID in the query.
// Returns a *NotSingularError when more than one SubjectField ID is found.
// Returns a *NotFoundError when no entities are found.
func (sfq *SubjectFieldQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sfq.Limit(2).IDs(setContextOp(ctx, sfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subjectfield.Label}
	default:
		err = &NotSingularError{subjectfield.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sfq *SubjectFieldQuery) OnlyIDX(ctx context.Context) int {
	id, err := sfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SubjectFields.
func (sfq *SubjectFieldQuery) All(ctx context.Context) ([]*SubjectField, error) {
	ctx = setContextOp(ctx, sfq.ctx, "All")
	if err := sfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SubjectField, *SubjectFieldQuery]()
	return withInterceptors[[]*SubjectField](ctx, sfq, qr, sfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sfq *SubjectFieldQuery) AllX(ctx context.Context) []*SubjectField {
	nodes, err := sfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SubjectField IDs.
func (sfq *SubjectFieldQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sfq.ctx.Unique == nil && sfq.path != nil {
		sfq.Unique(true)
	}
	ctx = setContextOp(ctx, sfq.ctx, "IDs")
	if err = sfq.Select(subjectfield.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sfq *SubjectFieldQuery) IDsX(ctx context.Context) []int {
	ids, err := sfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sfq *SubjectFieldQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Count")
	if err := sfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sfq, querierCount[*SubjectFieldQuery](), sfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sfq *SubjectFieldQuery) CountX(ctx context.Context) int {
	count, err := sfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sfq *SubjectFieldQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Exist")
	switch _, err := sfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sfq *SubjectFieldQuery) ExistX(ctx context.Context) bool {
	exist, err := sfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubjectFieldQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sfq *SubjectFieldQuery) Clone() *SubjectFieldQuery {
	if sfq == nil {
		return nil
	}
	return &SubjectFieldQuery{
		config:      sfq.config,
		ctx:         sfq.ctx.Clone(),
		order:       append([]subjectfield.OrderOption{}, sfq.order...),
		inters:      append([]Interceptor{}, sfq.inters...),
		predicates:  append([]predicate.SubjectField{}, sfq.predicates...),
		withSubject: sfq.withSubject.Clone(),
		// clone intermediate query.
		sql:  sfq.sql.Clone(),
		path: sfq.path,
	}
}

// WithSubject tells the query-builder to eager-load the nodes that are connected to
// the "subject" edge. The optional arguments are used to configure the query builder of the edge.
func (sfq *SubjectFieldQuery) WithSubject(opts ...func(*SubjectQuery)) *SubjectFieldQuery {
	query := (&SubjectClient{config: sfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sfq.withSubject = query
	return sfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rate1 uint32 `json:"rate_1,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SubjectField.Query().
//		GroupBy(subjectfield.FieldRate1).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sfq *SubjectFieldQuery) GroupBy(field string, fields ...string) *SubjectFieldGroupBy {
	sfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubjectFieldGroupBy{build: sfq}
	grbuild.flds = &sfq.ctx.Fields
	grbuild.label = subjectfield.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rate1 uint32 `json:"rate_1,omitempty"`
//	}
//
//	client.SubjectField.Query().
//		Select(subjectfield.FieldRate1).
//		Scan(ctx, &v)
func (sfq *SubjectFieldQuery) Select(fields ...string) *SubjectFieldSelect {
	sfq.ctx.Fields = append(sfq.ctx.Fields, fields...)
	sbuild := &SubjectFieldSelect{SubjectFieldQuery: sfq}
	sbuild.label = subjectfield.Label
	sbuild.flds, sbuild.scan = &sfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubjectFieldSelect configured with the given aggregations.
func (sfq *SubjectFieldQuery) Aggregate(fns ...AggregateFunc) *SubjectFieldSelect {
	return sfq.Select().Aggregate(fns...)
}

func (sfq *SubjectFieldQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sfq); err != nil {
				return err
			}
		}
	}
	for _, f := range sfq.ctx.Fields {
		if !subjectfield.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sfq.path != nil {
		prev, err := sfq.path(ctx)
		if err != nil {
			return err
		}
		sfq.sql = prev
	}
	return nil
}

func (sfq *SubjectFieldQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SubjectField, error) {
	var (
		nodes       = []*SubjectField{}
		withFKs     = sfq.withFKs
		_spec       = sfq.querySpec()
		loadedTypes = [1]bool{
			sfq.withSubject != nil,
		}
	)
	if sfq.withSubject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subjectfield.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SubjectField).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SubjectField{config: sfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sfq.withSubject; query != nil {
		if err := sfq.loadSubject(ctx, query, nodes, nil,
			func(n *SubjectField, e *Subject) { n.Edges.Subject = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sfq *SubjectFieldQuery) loadSubject(ctx context.Context, query *SubjectQuery, nodes []*SubjectField, init func(*SubjectField), assign func(*SubjectField, *Subject)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*SubjectField)
	for i := range nodes {
		if nodes[i].subject_subject_field == nil {
			continue
		}
		fk := *nodes[i].subject_subject_field
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(subject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subject_subject_field" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sfq *SubjectFieldQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sfq.querySpec()
	_spec.Node.Columns = sfq.ctx.Fields
	if len(sfq.ctx.Fields) > 0 {
		_spec.Unique = sfq.ctx.Unique != nil && *sfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sfq.driver, _spec)
}

func (sfq *SubjectFieldQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subjectfield.Table, subjectfield.Columns, sqlgraph.NewFieldSpec(subjectfield.FieldID, field.TypeInt))
	_spec.From = sfq.sql
	if unique := sfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sfq.path != nil {
		_spec.Unique = true
	}
	if fields := sfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subjectfield.FieldID)
		for i := range fields {
			if fields[i] != subjectfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sfq *SubjectFieldQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sfq.driver.Dialect())
	t1 := builder.Table(subjectfield.Table)
	columns := sfq.ctx.Fields
	if len(columns) == 0 {
		columns = subjectfield.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sfq.sql != nil {
		selector = sfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sfq.ctx.Unique != nil && *sfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sfq.predicates {
		p(selector)
	}
	for _, p := range sfq.order {
		p(selector)
	}
	if offset := sfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubjectFieldGroupBy is the group-by builder for SubjectField entities.
type SubjectFieldGroupBy struct {
	selector
	build *SubjectFieldQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sfgb *SubjectFieldGroupBy) Aggregate(fns ...AggregateFunc) *SubjectFieldGroupBy {
	sfgb.fns = append(sfgb.fns, fns...)
	return sfgb
}

// Scan applies the selector query and scans the result into the given value.
func (sfgb *SubjectFieldGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfgb.build.ctx, "GroupBy")
	if err := sfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubjectFieldQuery, *SubjectFieldGroupBy](ctx, sfgb.build, sfgb, sfgb.build.inters, v)
}

func (sfgb *SubjectFieldGroupBy) sqlScan(ctx context.Context, root *SubjectFieldQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sfgb.fns))
	for _, fn := range sfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sfgb.flds)+len(sfgb.fns))
		for _, f := range *sfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubjectFieldSelect is the builder for selecting fields of SubjectField entities.
type SubjectFieldSelect struct {
	*SubjectFieldQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sfs *SubjectFieldSelect) Aggregate(fns ...AggregateFunc) *SubjectFieldSelect {
	sfs.fns = append(sfs.fns, fns...)
	return sfs
}

// Scan applies the selector query and scans the result into the given value.
func (sfs *SubjectFieldSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfs.ctx, "Select")
	if err := sfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubjectFieldQuery, *SubjectFieldSelect](ctx, sfs.SubjectFieldQuery, sfs, sfs.inters, v)
}

func (sfs *SubjectFieldSelect) sqlScan(ctx context.Context, root *SubjectFieldQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sfs.fns))
	for _, fn := range sfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}