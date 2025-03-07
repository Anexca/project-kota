// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/examassesment"
	"common/ent/examattempt"
	"common/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamAssesmentQuery is the builder for querying ExamAssesment entities.
type ExamAssesmentQuery struct {
	config
	ctx         *QueryContext
	order       []examassesment.OrderOption
	inters      []Interceptor
	predicates  []predicate.ExamAssesment
	withAttempt *ExamAttemptQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamAssesmentQuery builder.
func (eaq *ExamAssesmentQuery) Where(ps ...predicate.ExamAssesment) *ExamAssesmentQuery {
	eaq.predicates = append(eaq.predicates, ps...)
	return eaq
}

// Limit the number of records to be returned by this query.
func (eaq *ExamAssesmentQuery) Limit(limit int) *ExamAssesmentQuery {
	eaq.ctx.Limit = &limit
	return eaq
}

// Offset to start from.
func (eaq *ExamAssesmentQuery) Offset(offset int) *ExamAssesmentQuery {
	eaq.ctx.Offset = &offset
	return eaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eaq *ExamAssesmentQuery) Unique(unique bool) *ExamAssesmentQuery {
	eaq.ctx.Unique = &unique
	return eaq
}

// Order specifies how the records should be ordered.
func (eaq *ExamAssesmentQuery) Order(o ...examassesment.OrderOption) *ExamAssesmentQuery {
	eaq.order = append(eaq.order, o...)
	return eaq
}

// QueryAttempt chains the current query on the "attempt" edge.
func (eaq *ExamAssesmentQuery) QueryAttempt() *ExamAttemptQuery {
	query := (&ExamAttemptClient{config: eaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examassesment.Table, examassesment.FieldID, selector),
			sqlgraph.To(examattempt.Table, examattempt.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, examassesment.AttemptTable, examassesment.AttemptColumn),
		)
		fromU = sqlgraph.SetNeighbors(eaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExamAssesment entity from the query.
// Returns a *NotFoundError when no ExamAssesment was found.
func (eaq *ExamAssesmentQuery) First(ctx context.Context) (*ExamAssesment, error) {
	nodes, err := eaq.Limit(1).All(setContextOp(ctx, eaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examassesment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) FirstX(ctx context.Context) *ExamAssesment {
	node, err := eaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamAssesment ID from the query.
// Returns a *NotFoundError when no ExamAssesment ID was found.
func (eaq *ExamAssesmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(1).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examassesment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) FirstIDX(ctx context.Context) int {
	id, err := eaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamAssesment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamAssesment entity is found.
// Returns a *NotFoundError when no ExamAssesment entities are found.
func (eaq *ExamAssesmentQuery) Only(ctx context.Context) (*ExamAssesment, error) {
	nodes, err := eaq.Limit(2).All(setContextOp(ctx, eaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examassesment.Label}
	default:
		return nil, &NotSingularError{examassesment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) OnlyX(ctx context.Context) *ExamAssesment {
	node, err := eaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamAssesment ID in the query.
// Returns a *NotSingularError when more than one ExamAssesment ID is found.
// Returns a *NotFoundError when no entities are found.
func (eaq *ExamAssesmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eaq.Limit(2).IDs(setContextOp(ctx, eaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examassesment.Label}
	default:
		err = &NotSingularError{examassesment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := eaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamAssesments.
func (eaq *ExamAssesmentQuery) All(ctx context.Context) ([]*ExamAssesment, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryAll)
	if err := eaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamAssesment, *ExamAssesmentQuery]()
	return withInterceptors[[]*ExamAssesment](ctx, eaq, qr, eaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) AllX(ctx context.Context) []*ExamAssesment {
	nodes, err := eaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamAssesment IDs.
func (eaq *ExamAssesmentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eaq.ctx.Unique == nil && eaq.path != nil {
		eaq.Unique(true)
	}
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryIDs)
	if err = eaq.Select(examassesment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) IDsX(ctx context.Context) []int {
	ids, err := eaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eaq *ExamAssesmentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryCount)
	if err := eaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eaq, querierCount[*ExamAssesmentQuery](), eaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) CountX(ctx context.Context) int {
	count, err := eaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eaq *ExamAssesmentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eaq.ctx, ent.OpQueryExist)
	switch _, err := eaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eaq *ExamAssesmentQuery) ExistX(ctx context.Context) bool {
	exist, err := eaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamAssesmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eaq *ExamAssesmentQuery) Clone() *ExamAssesmentQuery {
	if eaq == nil {
		return nil
	}
	return &ExamAssesmentQuery{
		config:      eaq.config,
		ctx:         eaq.ctx.Clone(),
		order:       append([]examassesment.OrderOption{}, eaq.order...),
		inters:      append([]Interceptor{}, eaq.inters...),
		predicates:  append([]predicate.ExamAssesment{}, eaq.predicates...),
		withAttempt: eaq.withAttempt.Clone(),
		// clone intermediate query.
		sql:  eaq.sql.Clone(),
		path: eaq.path,
	}
}

// WithAttempt tells the query-builder to eager-load the nodes that are connected to
// the "attempt" edge. The optional arguments are used to configure the query builder of the edge.
func (eaq *ExamAssesmentQuery) WithAttempt(opts ...func(*ExamAttemptQuery)) *ExamAssesmentQuery {
	query := (&ExamAttemptClient{config: eaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eaq.withAttempt = query
	return eaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CompletedSeconds int `json:"completed_seconds,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamAssesment.Query().
//		GroupBy(examassesment.FieldCompletedSeconds).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eaq *ExamAssesmentQuery) GroupBy(field string, fields ...string) *ExamAssesmentGroupBy {
	eaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamAssesmentGroupBy{build: eaq}
	grbuild.flds = &eaq.ctx.Fields
	grbuild.label = examassesment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CompletedSeconds int `json:"completed_seconds,omitempty"`
//	}
//
//	client.ExamAssesment.Query().
//		Select(examassesment.FieldCompletedSeconds).
//		Scan(ctx, &v)
func (eaq *ExamAssesmentQuery) Select(fields ...string) *ExamAssesmentSelect {
	eaq.ctx.Fields = append(eaq.ctx.Fields, fields...)
	sbuild := &ExamAssesmentSelect{ExamAssesmentQuery: eaq}
	sbuild.label = examassesment.Label
	sbuild.flds, sbuild.scan = &eaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamAssesmentSelect configured with the given aggregations.
func (eaq *ExamAssesmentQuery) Aggregate(fns ...AggregateFunc) *ExamAssesmentSelect {
	return eaq.Select().Aggregate(fns...)
}

func (eaq *ExamAssesmentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eaq); err != nil {
				return err
			}
		}
	}
	for _, f := range eaq.ctx.Fields {
		if !examassesment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eaq.path != nil {
		prev, err := eaq.path(ctx)
		if err != nil {
			return err
		}
		eaq.sql = prev
	}
	return nil
}

func (eaq *ExamAssesmentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamAssesment, error) {
	var (
		nodes       = []*ExamAssesment{}
		withFKs     = eaq.withFKs
		_spec       = eaq.querySpec()
		loadedTypes = [1]bool{
			eaq.withAttempt != nil,
		}
	)
	if eaq.withAttempt != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, examassesment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamAssesment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamAssesment{config: eaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eaq.withAttempt; query != nil {
		if err := eaq.loadAttempt(ctx, query, nodes, nil,
			func(n *ExamAssesment, e *ExamAttempt) { n.Edges.Attempt = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eaq *ExamAssesmentQuery) loadAttempt(ctx context.Context, query *ExamAttemptQuery, nodes []*ExamAssesment, init func(*ExamAssesment), assign func(*ExamAssesment, *ExamAttempt)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ExamAssesment)
	for i := range nodes {
		if nodes[i].exam_attempt_assesment == nil {
			continue
		}
		fk := *nodes[i].exam_attempt_assesment
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(examattempt.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exam_attempt_assesment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eaq *ExamAssesmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eaq.querySpec()
	_spec.Node.Columns = eaq.ctx.Fields
	if len(eaq.ctx.Fields) > 0 {
		_spec.Unique = eaq.ctx.Unique != nil && *eaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eaq.driver, _spec)
}

func (eaq *ExamAssesmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examassesment.Table, examassesment.Columns, sqlgraph.NewFieldSpec(examassesment.FieldID, field.TypeInt))
	_spec.From = eaq.sql
	if unique := eaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eaq.path != nil {
		_spec.Unique = true
	}
	if fields := eaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examassesment.FieldID)
		for i := range fields {
			if fields[i] != examassesment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eaq *ExamAssesmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eaq.driver.Dialect())
	t1 := builder.Table(examassesment.Table)
	columns := eaq.ctx.Fields
	if len(columns) == 0 {
		columns = examassesment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eaq.sql != nil {
		selector = eaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eaq.ctx.Unique != nil && *eaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eaq.predicates {
		p(selector)
	}
	for _, p := range eaq.order {
		p(selector)
	}
	if offset := eaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamAssesmentGroupBy is the group-by builder for ExamAssesment entities.
type ExamAssesmentGroupBy struct {
	selector
	build *ExamAssesmentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eagb *ExamAssesmentGroupBy) Aggregate(fns ...AggregateFunc) *ExamAssesmentGroupBy {
	eagb.fns = append(eagb.fns, fns...)
	return eagb
}

// Scan applies the selector query and scans the result into the given value.
func (eagb *ExamAssesmentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eagb.build.ctx, ent.OpQueryGroupBy)
	if err := eagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamAssesmentQuery, *ExamAssesmentGroupBy](ctx, eagb.build, eagb, eagb.build.inters, v)
}

func (eagb *ExamAssesmentGroupBy) sqlScan(ctx context.Context, root *ExamAssesmentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eagb.fns))
	for _, fn := range eagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eagb.flds)+len(eagb.fns))
		for _, f := range *eagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamAssesmentSelect is the builder for selecting fields of ExamAssesment entities.
type ExamAssesmentSelect struct {
	*ExamAssesmentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eas *ExamAssesmentSelect) Aggregate(fns ...AggregateFunc) *ExamAssesmentSelect {
	eas.fns = append(eas.fns, fns...)
	return eas
}

// Scan applies the selector query and scans the result into the given value.
func (eas *ExamAssesmentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eas.ctx, ent.OpQuerySelect)
	if err := eas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamAssesmentQuery, *ExamAssesmentSelect](ctx, eas.ExamAssesmentQuery, eas, eas.inters, v)
}

func (eas *ExamAssesmentSelect) sqlScan(ctx context.Context, root *ExamAssesmentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eas.fns))
	for _, fn := range eas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
