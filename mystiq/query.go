package mystiq

type Query struct {
	Table  string
	Fields map[string]string
	Value  string
	Alias  string
}

func (q Query) CanUse() bool {
	return q.Table != "" && q.Value != "" && q.Alias != "" && len(q.Fields) > 0
}
