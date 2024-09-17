// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExamCategoryQuery is the builder for querying ExamCategory entities.
type ExamCategoryQuery struct {
	config
	ctx        *QueryContext
	order      []examcategory.OrderOption
	inters     []Interceptor
	predicates []predicate.ExamCategory
	withExams  *ExamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamCategoryQuery builder.
func (ecq *ExamCategoryQuery) Where(ps ...predicate.ExamCategory) *ExamCategoryQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *ExamCategoryQuery) Limit(limit int) *ExamCategoryQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *ExamCategoryQuery) Offset(offset int) *ExamCategoryQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *ExamCategoryQuery) Unique(unique bool) *ExamCategoryQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *ExamCategoryQuery) Order(o ...examcategory.OrderOption) *ExamCategoryQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryExams chains the current query on the "exams" edge.
func (ecq *ExamCategoryQuery) QueryExams() *ExamQuery {
	query := (&ExamClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examcategory.Table, examcategory.FieldID, selector),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examcategory.ExamsTable, examcategory.ExamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExamCategory entity from the query.
// Returns a *NotFoundError when no ExamCategory was found.
func (ecq *ExamCategoryQuery) First(ctx context.Context) (*ExamCategory, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *ExamCategoryQuery) FirstX(ctx context.Context) *ExamCategory {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamCategory ID from the query.
// Returns a *NotFoundError when no ExamCategory ID was found.
func (ecq *ExamCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *ExamCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamCategory entity is found.
// Returns a *NotFoundError when no ExamCategory entities are found.
func (ecq *ExamCategoryQuery) Only(ctx context.Context) (*ExamCategory, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examcategory.Label}
	default:
		return nil, &NotSingularError{examcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *ExamCategoryQuery) OnlyX(ctx context.Context) *ExamCategory {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamCategory ID in the query.
// Returns a *NotSingularError when more than one ExamCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *ExamCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examcategory.Label}
	default:
		err = &NotSingularError{examcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *ExamCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamCategories.
func (ecq *ExamCategoryQuery) All(ctx context.Context) ([]*ExamCategory, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryAll)
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamCategory, *ExamCategoryQuery]()
	return withInterceptors[[]*ExamCategory](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *ExamCategoryQuery) AllX(ctx context.Context) []*ExamCategory {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamCategory IDs.
func (ecq *ExamCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryIDs)
	if err = ecq.Select(examcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *ExamCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *ExamCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryCount)
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*ExamCategoryQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *ExamCategoryQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *ExamCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryExist)
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *ExamCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *ExamCategoryQuery) Clone() *ExamCategoryQuery {
	if ecq == nil {
		return nil
	}
	return &ExamCategoryQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]examcategory.OrderOption{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.ExamCategory{}, ecq.predicates...),
		withExams:  ecq.withExams.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithExams tells the query-builder to eager-load the nodes that are connected to
// the "exams" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *ExamCategoryQuery) WithExams(opts ...func(*ExamQuery)) *ExamCategoryQuery {
	query := (&ExamClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withExams = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name examcategory.Name `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamCategory.Query().
//		GroupBy(examcategory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *ExamCategoryQuery) GroupBy(field string, fields ...string) *ExamCategoryGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamCategoryGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = examcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name examcategory.Name `json:"name,omitempty"`
//	}
//
//	client.ExamCategory.Query().
//		Select(examcategory.FieldName).
//		Scan(ctx, &v)
func (ecq *ExamCategoryQuery) Select(fields ...string) *ExamCategorySelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &ExamCategorySelect{ExamCategoryQuery: ecq}
	sbuild.label = examcategory.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamCategorySelect configured with the given aggregations.
func (ecq *ExamCategoryQuery) Aggregate(fns ...AggregateFunc) *ExamCategorySelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *ExamCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !examcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *ExamCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamCategory, error) {
	var (
		nodes       = []*ExamCategory{}
		_spec       = ecq.querySpec()
		loadedTypes = [1]bool{
			ecq.withExams != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamCategory{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecq.withExams; query != nil {
		if err := ecq.loadExams(ctx, query, nodes,
			func(n *ExamCategory) { n.Edges.Exams = []*Exam{} },
			func(n *ExamCategory, e *Exam) { n.Edges.Exams = append(n.Edges.Exams, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *ExamCategoryQuery) loadExams(ctx context.Context, query *ExamQuery, nodes []*ExamCategory, init func(*ExamCategory), assign func(*ExamCategory, *Exam)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ExamCategory)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examcategory.ExamsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_category_exams
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_category_exams" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_category_exams" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ecq *ExamCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *ExamCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examcategory.Table, examcategory.Columns, sqlgraph.NewFieldSpec(examcategory.FieldID, field.TypeInt))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examcategory.FieldID)
		for i := range fields {
			if fields[i] != examcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *ExamCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(examcategory.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = examcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamCategoryGroupBy is the group-by builder for ExamCategory entities.
type ExamCategoryGroupBy struct {
	selector
	build *ExamCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *ExamCategoryGroupBy) Aggregate(fns ...AggregateFunc) *ExamCategoryGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *ExamCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, ent.OpQueryGroupBy)
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCategoryQuery, *ExamCategoryGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *ExamCategoryGroupBy) sqlScan(ctx context.Context, root *ExamCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamCategorySelect is the builder for selecting fields of ExamCategory entities.
type ExamCategorySelect struct {
	*ExamCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *ExamCategorySelect) Aggregate(fns ...AggregateFunc) *ExamCategorySelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *ExamCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, ent.OpQuerySelect)
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamCategoryQuery, *ExamCategorySelect](ctx, ecs.ExamCategoryQuery, ecs, ecs.inters, v)
}

func (ecs *ExamCategorySelect) sqlScan(ctx context.Context, root *ExamCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
