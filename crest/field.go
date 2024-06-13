package crest

import (
	"fmt"
)

type Field interface {
	QueryBuilder
	Default(value any) Field
	NotNull(notNull ...bool) Field
	Prefix(prefix string) Field
	PrimaryKey(primaryKey ...bool) Field
	Type(dataType string) Field
	Unique(unique ...bool) Field
	Relationship(relationship Field) Field
	CreateValue(fn func(operation string, values Map) Value) Field
	Name() string
	TsVector() Field
	Serial() Field
	Bool() Field
	Char(n int) Field
	Varchar(n int) Field
	Jsonb() Field
	Int() Field
	Float() Field
	Timestamp() Field
	Timestampz() Field
}

type field struct {
	Field
	defaultValue string
	dataType     string
	name         string
	notNull      bool
	table        string
	prefix       string
	primaryKey   bool
	relationship *field
	unique       bool
	valueFactory func(operation string, values Map) Value
}

const (
	Id = "id"
)

func (f *field) Name() string {
	return f.name
}

func (f *field) CreateValue(fn func(operation string, values Map) Value) Field {
	f.valueFactory = fn
	return f
}

func (f *field) Default(value any) Field {
	switch v := value.(type) {
	case Safe:
		f.defaultValue = fmt.Sprintf("%s", string(v))
	case string:
		f.defaultValue = fmt.Sprintf("'%s'", v)
	default:
		f.defaultValue = fmt.Sprintf("%v", value)
	}
	return f
}

func (f *field) NotNull(notNull ...bool) Field {
	n := len(notNull)
	if n == 0 {
		f.notNull = true
	}
	if n > 0 {
		f.notNull = notNull[0]
	}
	return f
}

func (f *field) Prefix(prefix string) Field {
	f.prefix = prefix
	return f
}

func (f *field) PrimaryKey(primaryKey ...bool) Field {
	n := len(primaryKey)
	if n == 0 {
		f.primaryKey = true
	}
	if n > 0 {
		f.primaryKey = primaryKey[0]
	}
	f.notNull = true
	return f
}

func (f *field) Relationship(relationship Field) Field {
	f.relationship = relationship.(*field)
	return f
}

func (f *field) Type(dataType string) Field {
	f.dataType = dataType
	return f
}

func (f *field) Serial() Field {
	f.Type("SERIAL")
	return f
}

func (f *field) Bool() Field {
	f.Type("BOOL")
	return f
}

func (f *field) Char(n int) Field {
	f.Type(fmt.Sprintf("CHAR(%d)", n))
	return f
}

func (f *field) Varchar(n int) Field {
	f.Type(fmt.Sprintf("VARCHAR(%d)", n))
	return f
}

func (f *field) Int() Field {
	f.Type("INT")
	return f
}

func (f *field) Float() Field {
	f.Type("FLOAT")
	return f
}

func (f *field) Jsonb() Field {
	f.Type("JSONB")
	return f
}

func (f *field) Timestamp() Field {
	f.Type("TIMESTAMP")
	return f
}

func (f *field) Timestampz() Field {
	f.Type("TIMESTAMPZ")
	return f
}

func (f *field) TsVector() Field {
	f.dataType = TsVectorDataType
	return f
}

func (f *field) Unique(unique ...bool) Field {
	n := len(unique)
	if n == 0 {
		f.unique = true
	}
	if n > 0 {
		f.unique = unique[0]
	}
	return f
}

func (f *field) Build() BuildResult {
	return BuildResult{f.prefix + "." + f.name, nil}
}
