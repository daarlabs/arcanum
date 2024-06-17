package quirk

import (
	"database/sql"
	"reflect"
)

type Field struct {
	Name  string
	Props string
	Value any
}

func Nullable[T any](v T) sql.Null[T] {
	return sql.Null[T]{
		V:     v,
		Valid: !reflect.ValueOf(v).IsZero(),
	}
}
