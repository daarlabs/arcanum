package crest

import (
	"strings"
)

type SaveRepository[R result] interface {
	QueryBuilder
	Run(runner ...Runner) (R, error)
	MustRun(runner ...Runner) R
}

type saveRepository[E entity, R result] struct {
	*repository[E, R]
	filters         []*filterBuilder
	relationships   []*relationshipBuilder
	selectors       []*selectorBuilder
	temporaries     []*temporaryBuilder
	values          []*valuesBuilder
	primaryKeyValue any
}

const (
	Insert = "INSERT"
	Update = "UPDATE"
)

func (r *saveRepository[E, R]) buildValues() map[string]any {
	values := make(map[string]any)
	for _, vb := range r.values {
		b := vb.Build()
		for k, v := range b.Values {
			values[k] = v
		}
	}
	return values
}

func (r *saveRepository[E, R]) createFieldsValues(operation string, values *map[string]any, fields ...Field) {
	for _, item := range fields {
		f := item.(*field)
		if f.valueFactory == nil {
			continue
		}
		v := f.valueFactory(operation, *values)
		if v == nil {
			continue
		}
		(*values)[f.name] = v
	}
}

func (r *saveRepository[E, R]) Build() BuildResult {
	e := any(r.entity).(entity)
	fields := e.Fields()
	values := r.buildValues()
	primaryKeyField := getPrimaryKeyField(fields...)
	if primaryKeyField == nil {
		panic(ErrorMissingPrimaryKey)
	}
	primaryKeyValue, ok := values[primaryKeyField.name]
	if ok {
		r.primaryKeyValue = primaryKeyValue
	}
	if r.primaryKeyValue == nil {
		r.createFieldsValues(Insert, &values, fields...)
		return r.buildInsert(values)
	}
	r.createFieldsValues(Update, &values, fields...)
	r.appendPrimaryKeyFilterIfNecessary(primaryKeyField)
	return r.buildUpdate(values)
}

func (r *saveRepository[E, R]) Run(runner ...Runner) (R, error) {
	res := new(R)
	if r.db == nil {
		return *res, ErrorMissingDatabase
	}
	if len(runner) > 0 {
		return *res, nil
	}
	b := r.Build()
	if err := r.db.Q(b.Sql, b.Values).Exec(res); err != nil {
		return *res, err
	}
	return *res, nil
}

func (r *saveRepository[E, R]) MustRun(runner ...Runner) R {
	res, err := r.Run(runner...)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *saveRepository[E, R]) buildWith() BuildResult {
	sql := make([]string, len(r.temporaries))
	values := make(map[string]any)
	for i, t := range r.temporaries {
		b := t.Build()
		sql[i] = b.Sql
	}
	return BuildResult{strings.Join(sql, ","), values}
}

func (r *saveRepository[E, R]) buildInsert(values map[string]any) BuildResult {
	selectorsExist := len(r.selectors) > 0
	temporariesExist := len(r.temporaries) > 0
	e := any(r.entity).(entity)
	fields := e.Fields()
	q := createSqlBuilder()
	if temporariesExist {
		tb := r.buildWith()
		q.Q("WITH " + tb.Sql)
		for k, v := range tb.Values {
			values[k] = v
		}
	}
	q.Q("INSERT INTO " + e.Table()).
		Q("AS " + e.Alias()).
		Q("(" + buildFieldsSqlWithoutPrimaryKey(fields...) + ")").
		Q("VALUES (" + createInsertSqlFromValues(fields, values) + ")")
	
	q.If(!selectorsExist, "RETURNING *").
		If(selectorsExist, "RETURNING "+buildFieldsSql(r.selectors...))
	
	return BuildResult{q.Build(), values}
}

func (r *saveRepository[E, R]) buildUpdate(values map[string]any) BuildResult {
	selectorsExist := len(r.selectors) > 0
	temporariesExist := len(r.temporaries) > 0
	e := any(r.entity).(entity)
	fields := e.Fields()
	q := createSqlBuilder()
	if temporariesExist {
		tb := r.buildWith()
		q.Q("WITH " + tb.Sql)
		for k, v := range tb.Values {
			values[k] = v
		}
	}
	q.Q("UPDATE " + e.Table()).
		Q("AS " + e.Alias()).
		Q("SET " + createUpdateSqlFromValues(fields, values))
	
	// Where
	buildBeforeAggregationFilters(q, r.filters, &values)
	
	q.If(!selectorsExist, "RETURNING *").
		If(selectorsExist, "RETURNING "+buildFieldsSql(r.selectors...))
	
	return BuildResult{q.Build(), values}
}

func (r *saveRepository[E, R]) appendPrimaryKeyFilterIfNecessary(primaryKeyField *field) {
	for _, f := range r.filters {
		for _, p := range f.parts {
			if p.sql == primaryKeyField.prefix+"."+primaryKeyField.name {
				return
			}
		}
	}
	r.filters = append(
		r.filters,
		Filter().Field(primaryKeyField).Equal().Value(r.primaryKeyValue, primaryKeyField.name).(*filterBuilder),
	)
}
