package crest

type FindRepository interface {
	QueryBuilder
	Run(result any, runner ...Runner) error
	MustRun(result any, runner ...Runner)
}

type findRepository[E entity] struct {
	*repository[E]
	filters       []*filterBuilder
	relationships []*relationshipBuilder
	selectors     []*selectorBuilder
	shapes        []*shapeBuilder
}

func (r *findRepository[E]) Build() BuildResult {
	var fieldsSql string
	values := make(map[string]any)
	e := any(r.entity).(entity)
	fields := e.Fields()
	selectorsExist := len(r.selectors) > 0
	if !selectorsExist {
		fieldsSql = buildFieldsSql(fields...)
	}
	if selectorsExist {
		fieldsSql = buildFieldsSql(r.selectors...)
	}
	q := createSqlBuilder().
		Q("SELECT").
		If(doesExistDistinct(r.shapes), "DISTINCT").
		Q(fieldsSql).
		Q("FROM " + e.Table()).
		Q("AS " + e.Alias())
	
	// Joins
	buildJoins(q, r.relationships, fields)
	
	// Where
	buildBeforeAggregationFilters(q, r.filters, &values)
	
	// Group shapes
	groupShapes := buildGroupShapes(r.shapes)
	q = q.If(len(groupShapes) > 0, groupShapes)
	
	// Having
	buildAfterAggregationFilters(q, r.filters, &values)
	
	// Order, Limit, Offset
	nonGroupShapes := buildNonGroupShapes(r.shapes)
	q = q.If(len(nonGroupShapes) > 0, nonGroupShapes)
	
	return BuildResult{q.Build(), values}
}

func (r *findRepository[E]) Run(result any, runner ...Runner) error {
	if r.db == nil {
		return ErrorMissingDatabase
	}
	if len(runner) > 0 {
		return nil
	}
	b := r.Build()
	if result == nil {
		if err := r.db.Q(b.Sql, b.Values).Exec(); err != nil {
			return err
		}
	}
	if result != nil {
		if err := r.db.Q(b.Sql, b.Values).Exec(result); err != nil {
			return err
		}
	}
	return nil
}

func (r *findRepository[E]) MustRun(result any, runner ...Runner) {
	err := r.Run(result, runner...)
	if err != nil {
		panic(err)
	}
}
