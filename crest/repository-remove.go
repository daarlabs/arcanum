package crest

type RemoveRepository interface {
	QueryBuilder
	Run(result any, runner ...Runner) error
	MustRun(result any, runner ...Runner)
}

type removeRepository[E entity] struct {
	*repository[E]
	filters   []*filterBuilder
	selectors []*selectorBuilder
}

func (r *removeRepository[E]) Build() BuildResult {
	values := make(map[string]any)
	selectorsExist := len(r.selectors) > 0
	e := any(r.entity).(entity)
	q := createSqlBuilder().
		Q("DELETE").
		Q("FROM " + e.Table()).
		Q("AS " + e.Alias())
	
	// Where
	buildBeforeAggregationFilters(q, r.filters, &values)
	
	q.If(!selectorsExist, "RETURNING *").
		If(selectorsExist, "RETURNING "+buildFieldsSql(r.selectors...))
	
	return BuildResult{q.Build(), values}
}

func (r *removeRepository[E]) Run(result any, runner ...Runner) error {
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

func (r *removeRepository[E]) MustRun(result any, runner ...Runner) {
	err := r.Run(result, runner...)
	if err != nil {
		panic(err)
	}
}
