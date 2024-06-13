package crest

import (
	"strings"
)

type SaveRepository interface {
	QueryBuilder
	Run(result any, runner ...Runner) error
	MustRun(result any, runner ...Runner)
	ForceInsert() SaveRepository
}

type saveRepository[E entity] struct {
	*repository[E]
	filters         []*filterBuilder
	relationships   []*relationshipBuilder
	selectors       []*selectorBuilder
	temporaries     []*temporaryBuilder
	values          []*valuesBuilder
	primaryKeyValue any
	forceInsert     bool
}

const (
	Insert = "INSERT"
	Update = "UPDATE"
)

func (r *saveRepository[E]) ForceInsert() SaveRepository {
	r.forceInsert = true
	return r
}

func (r *saveRepository[E]) buildValues() map[string]any {
	values := make(map[string]any)
	for _, vb := range r.values {
		b := vb.Build()
		for k, v := range b.Values {
			values[k] = v
		}
	}
	return values
}

func (r *saveRepository[E]) createFieldsValues(operation string, values *map[string]any, fields ...Field) {
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

func (r *saveRepository[E]) Build() BuildResult {
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
	if r.forceInsert || r.primaryKeyValue == nil {
		r.createFieldsValues(Insert, &values, fields...)
		return r.buildInsert(values)
	}
	r.createFieldsValues(Update, &values, fields...)
	r.appendPrimaryKeyFilterIfNecessary(primaryKeyField)
	return r.buildUpdate(values)
}

func (r *saveRepository[E]) Run(result any, runner ...Runner) error {
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

func (r *saveRepository[E]) MustRun(result any, runner ...Runner) {
	err := r.Run(result, runner...)
	if err != nil {
		panic(err)
	}
}

func (r *saveRepository[E]) buildWith() BuildResult {
	sql := make([]string, len(r.temporaries))
	values := make(map[string]any)
	for i, t := range r.temporaries {
		b := t.Build()
		sql[i] = b.Sql
	}
	return BuildResult{strings.Join(sql, ","), values}
}

func (r *saveRepository[E]) buildInsert(values map[string]any) BuildResult {
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
	var queryFields string
	if !r.forceInsert {
		queryFields = buildFieldsSqlWithoutPrimaryKey(fields...)
	}
	if r.forceInsert {
		queryFields = buildFieldsSql(fields...)
	}
	q.Q("INSERT INTO " + e.Table()).
		Q("(" + queryFields + ")").
		Q("VALUES (" + createInsertSqlFromValues(r.forceInsert, fields, values) + ")")
	
	q.If(!selectorsExist, "RETURNING *").
		If(selectorsExist, "RETURNING "+buildFieldsSql(r.selectors...))
	
	return BuildResult{strings.ReplaceAll(q.Build(), e.Alias()+".", ""), values}
}

func (r *saveRepository[E]) buildUpdate(values map[string]any) BuildResult {
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
		Q("SET " + createUpdateSqlFromValues(fields, values))
	
	// Where
	buildBeforeAggregationFilters(q, r.filters, &values)
	
	q.If(!selectorsExist, "RETURNING *").
		If(selectorsExist, "RETURNING "+buildFieldsSql(r.selectors...))
	
	return BuildResult{strings.ReplaceAll(q.Build(), e.Alias()+".", ""), values}
}

func (r *saveRepository[E]) appendPrimaryKeyFilterIfNecessary(primaryKeyField *field) {
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
