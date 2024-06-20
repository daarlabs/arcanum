package mystiq

import (
	"strings"
)

type Query struct {
	Table  string
	Fields map[string]string
	Value  string
	Alias  string
}

func CreateSelect(table, textColumn string, valueColumn ...string) Query {
	var alias string
	for _, item := range strings.Split(table, "-") {
		alias += strings.ToLower(item[0:1])
	}
	value := "id"
	if len(valueColumn) > 0 {
		value = valueColumn[0]
	}
	return Query{
		Table: table,
		Fields: map[string]string{
			"text":  alias + "." + textColumn,
			"value": alias + "." + value,
		},
		Value: value,
		Alias: alias,
	}
}

func (q Query) CanUse() bool {
	return q.Table != "" && q.Value != "" && q.Alias != "" && len(q.Fields) > 0
}
