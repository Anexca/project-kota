// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/ent/exam"
	"common/ent/examcategory"
	"common/ent/examgroup"
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

// ExamGroupQuery is the builder for querying ExamGroup entities.
type ExamGroupQuery struct {
	config
	ctx          *QueryContext
	order        []examgroup.OrderOption
	inters       []Interceptor
	predicates   []predicate.ExamGroup
	withCategory *ExamCategoryQuery
	withExams    *ExamQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExamGroupQuery builder.
func (egq *ExamGroupQuery) Where(ps ...predicate.ExamGroup) *ExamGroupQuery {
	egq.predicates = append(egq.predicates, ps...)
	return egq
}

// Limit the number of records to be returned by this query.
func (egq *ExamGroupQuery) Limit(limit int) *ExamGroupQuery {
	egq.ctx.Limit = &limit
	return egq
}

// Offset to start from.
func (egq *ExamGroupQuery) Offset(offset int) *ExamGroupQuery {
	egq.ctx.Offset = &offset
	return egq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (egq *ExamGroupQuery) Unique(unique bool) *ExamGroupQuery {
	egq.ctx.Unique = &unique
	return egq
}

// Order specifies how the records should be ordered.
func (egq *ExamGroupQuery) Order(o ...examgroup.OrderOption) *ExamGroupQuery {
	egq.order = append(egq.order, o...)
	return egq
}

// QueryCategory chains the current query on the "category" edge.
func (egq *ExamGroupQuery) QueryCategory() *ExamCategoryQuery {
	query := (&ExamCategoryClient{config: egq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := egq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := egq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examgroup.Table, examgroup.FieldID, selector),
			sqlgraph.To(examcategory.Table, examcategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, examgroup.CategoryTable, examgroup.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(egq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExams chains the current query on the "exams" edge.
func (egq *ExamGroupQuery) QueryExams() *ExamQuery {
	query := (&ExamClient{config: egq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := egq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := egq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(examgroup.Table, examgroup.FieldID, selector),
			sqlgraph.To(exam.Table, exam.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, examgroup.ExamsTable, examgroup.ExamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(egq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExamGroup entity from the query.
// Returns a *NotFoundError when no ExamGroup was found.
func (egq *ExamGroupQuery) First(ctx context.Context) (*ExamGroup, error) {
	nodes, err := egq.Limit(1).All(setContextOp(ctx, egq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{examgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (egq *ExamGroupQuery) FirstX(ctx context.Context) *ExamGroup {
	node, err := egq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExamGroup ID from the query.
// Returns a *NotFoundError when no ExamGroup ID was found.
func (egq *ExamGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = egq.Limit(1).IDs(setContextOp(ctx, egq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{examgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (egq *ExamGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := egq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExamGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExamGroup entity is found.
// Returns a *NotFoundError when no ExamGroup entities are found.
func (egq *ExamGroupQuery) Only(ctx context.Context) (*ExamGroup, error) {
	nodes, err := egq.Limit(2).All(setContextOp(ctx, egq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{examgroup.Label}
	default:
		return nil, &NotSingularError{examgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (egq *ExamGroupQuery) OnlyX(ctx context.Context) *ExamGroup {
	node, err := egq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExamGroup ID in the query.
// Returns a *NotSingularError when more than one ExamGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (egq *ExamGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = egq.Limit(2).IDs(setContextOp(ctx, egq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{examgroup.Label}
	default:
		err = &NotSingularError{examgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (egq *ExamGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := egq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExamGroups.
func (egq *ExamGroupQuery) All(ctx context.Context) ([]*ExamGroup, error) {
	ctx = setContextOp(ctx, egq.ctx, ent.OpQueryAll)
	if err := egq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExamGroup, *ExamGroupQuery]()
	return withInterceptors[[]*ExamGroup](ctx, egq, qr, egq.inters)
}

// AllX is like All, but panics if an error occurs.
func (egq *ExamGroupQuery) AllX(ctx context.Context) []*ExamGroup {
	nodes, err := egq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExamGroup IDs.
func (egq *ExamGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if egq.ctx.Unique == nil && egq.path != nil {
		egq.Unique(true)
	}
	ctx = setContextOp(ctx, egq.ctx, ent.OpQueryIDs)
	if err = egq.Select(examgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (egq *ExamGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := egq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (egq *ExamGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, egq.ctx, ent.OpQueryCount)
	if err := egq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, egq, querierCount[*ExamGroupQuery](), egq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (egq *ExamGroupQuery) CountX(ctx context.Context) int {
	count, err := egq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (egq *ExamGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, egq.ctx, ent.OpQueryExist)
	switch _, err := egq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (egq *ExamGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := egq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExamGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (egq *ExamGroupQuery) Clone() *ExamGroupQuery {
	if egq == nil {
		return nil
	}
	return &ExamGroupQuery{
		config:       egq.config,
		ctx:          egq.ctx.Clone(),
		order:        append([]examgroup.OrderOption{}, egq.order...),
		inters:       append([]Interceptor{}, egq.inters...),
		predicates:   append([]predicate.ExamGroup{}, egq.predicates...),
		withCategory: egq.withCategory.Clone(),
		withExams:    egq.withExams.Clone(),
		// clone intermediate query.
		sql:  egq.sql.Clone(),
		path: egq.path,
	}
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (egq *ExamGroupQuery) WithCategory(opts ...func(*ExamCategoryQuery)) *ExamGroupQuery {
	query := (&ExamCategoryClient{config: egq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	egq.withCategory = query
	return egq
}

// WithExams tells the query-builder to eager-load the nodes that are connected to
// the "exams" edge. The optional arguments are used to configure the query builder of the edge.
func (egq *ExamGroupQuery) WithExams(opts ...func(*ExamQuery)) *ExamGroupQuery {
	query := (&ExamClient{config: egq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	egq.withExams = query
	return egq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExamGroup.Query().
//		GroupBy(examgroup.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (egq *ExamGroupQuery) GroupBy(field string, fields ...string) *ExamGroupGroupBy {
	egq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExamGroupGroupBy{build: egq}
	grbuild.flds = &egq.ctx.Fields
	grbuild.label = examgroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ExamGroup.Query().
//		Select(examgroup.FieldName).
//		Scan(ctx, &v)
func (egq *ExamGroupQuery) Select(fields ...string) *ExamGroupSelect {
	egq.ctx.Fields = append(egq.ctx.Fields, fields...)
	sbuild := &ExamGroupSelect{ExamGroupQuery: egq}
	sbuild.label = examgroup.Label
	sbuild.flds, sbuild.scan = &egq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExamGroupSelect configured with the given aggregations.
func (egq *ExamGroupQuery) Aggregate(fns ...AggregateFunc) *ExamGroupSelect {
	return egq.Select().Aggregate(fns...)
}

func (egq *ExamGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range egq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, egq); err != nil {
				return err
			}
		}
	}
	for _, f := range egq.ctx.Fields {
		if !examgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if egq.path != nil {
		prev, err := egq.path(ctx)
		if err != nil {
			return err
		}
		egq.sql = prev
	}
	return nil
}

func (egq *ExamGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExamGroup, error) {
	var (
		nodes       = []*ExamGroup{}
		withFKs     = egq.withFKs
		_spec       = egq.querySpec()
		loadedTypes = [2]bool{
			egq.withCategory != nil,
			egq.withExams != nil,
		}
	)
	if egq.withCategory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, examgroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExamGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExamGroup{config: egq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, egq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := egq.withCategory; query != nil {
		if err := egq.loadCategory(ctx, query, nodes, nil,
			func(n *ExamGroup, e *ExamCategory) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	if query := egq.withExams; query != nil {
		if err := egq.loadExams(ctx, query, nodes,
			func(n *ExamGroup) { n.Edges.Exams = []*Exam{} },
			func(n *ExamGroup, e *Exam) { n.Edges.Exams = append(n.Edges.Exams, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (egq *ExamGroupQuery) loadCategory(ctx context.Context, query *ExamCategoryQuery, nodes []*ExamGroup, init func(*ExamGroup), assign func(*ExamGroup, *ExamCategory)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ExamGroup)
	for i := range nodes {
		if nodes[i].exam_category_groups == nil {
			continue
		}
		fk := *nodes[i].exam_category_groups
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(examcategory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exam_category_groups" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (egq *ExamGroupQuery) loadExams(ctx context.Context, query *ExamQuery, nodes []*ExamGroup, init func(*ExamGroup), assign func(*ExamGroup, *Exam)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ExamGroup)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Exam(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(examgroup.ExamsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.exam_group_exams
		if fk == nil {
			return fmt.Errorf(`foreign-key "exam_group_exams" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "exam_group_exams" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (egq *ExamGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := egq.querySpec()
	_spec.Node.Columns = egq.ctx.Fields
	if len(egq.ctx.Fields) > 0 {
		_spec.Unique = egq.ctx.Unique != nil && *egq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, egq.driver, _spec)
}

func (egq *ExamGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(examgroup.Table, examgroup.Columns, sqlgraph.NewFieldSpec(examgroup.FieldID, field.TypeInt))
	_spec.From = egq.sql
	if unique := egq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if egq.path != nil {
		_spec.Unique = true
	}
	if fields := egq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, examgroup.FieldID)
		for i := range fields {
			if fields[i] != examgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := egq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := egq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := egq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := egq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (egq *ExamGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(egq.driver.Dialect())
	t1 := builder.Table(examgroup.Table)
	columns := egq.ctx.Fields
	if len(columns) == 0 {
		columns = examgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if egq.sql != nil {
		selector = egq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if egq.ctx.Unique != nil && *egq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range egq.predicates {
		p(selector)
	}
	for _, p := range egq.order {
		p(selector)
	}
	if offset := egq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := egq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExamGroupGroupBy is the group-by builder for ExamGroup entities.
type ExamGroupGroupBy struct {
	selector
	build *ExamGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eggb *ExamGroupGroupBy) Aggregate(fns ...AggregateFunc) *ExamGroupGroupBy {
	eggb.fns = append(eggb.fns, fns...)
	return eggb
}

// Scan applies the selector query and scans the result into the given value.
func (eggb *ExamGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eggb.build.ctx, ent.OpQueryGroupBy)
	if err := eggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamGroupQuery, *ExamGroupGroupBy](ctx, eggb.build, eggb, eggb.build.inters, v)
}

func (eggb *ExamGroupGroupBy) sqlScan(ctx context.Context, root *ExamGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eggb.fns))
	for _, fn := range eggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eggb.flds)+len(eggb.fns))
		for _, f := range *eggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExamGroupSelect is the builder for selecting fields of ExamGroup entities.
type ExamGroupSelect struct {
	*ExamGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (egs *ExamGroupSelect) Aggregate(fns ...AggregateFunc) *ExamGroupSelect {
	egs.fns = append(egs.fns, fns...)
	return egs
}

// Scan applies the selector query and scans the result into the given value.
func (egs *ExamGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egs.ctx, ent.OpQuerySelect)
	if err := egs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExamGroupQuery, *ExamGroupSelect](ctx, egs.ExamGroupQuery, egs, egs.inters, v)
}

func (egs *ExamGroupSelect) sqlScan(ctx context.Context, root *ExamGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(egs.fns))
	for _, fn := range egs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*egs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
