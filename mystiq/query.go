package mystiq

type Query struct {
	Table  string
	Fields map[string]string
	Name   string
	Alias  string
}

func (q Query) CanUse() bool {
	return q.Table != "" && q.Name != "" && q.Alias != "" && len(q.Fields) > 0
}
