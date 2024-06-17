package crest

import (
	"reflect"
)

type entity interface {
	Table() string
	Alias() string
	Fields() []Field
}

type Override struct {
	Table string
	Alias string
}

func Entity[E entity](override ...Override) *E {
	e := new(E)
	p := any(e).(entity)
	table := p.Table()
	prefix := p.Alias()
	if len(override) > 0 {
		if override[0].Table != "" {
			table = override[0].Table
		}
		if override[0].Alias != "" {
			prefix = override[0].Alias
		}
	}
	v := reflect.ValueOf(e)
	f := v.Elem().FieldByName(entityBuilderFieldName)
	if !f.IsValid() {
		return e
	}
	f.Set(reflect.ValueOf(EntityBuilder{table: table, prefix: prefix}))
	return e
}
