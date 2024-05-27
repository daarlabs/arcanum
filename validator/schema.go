package validator

type Schema interface {
	Add(name string) SchemaField
}

type schema struct {
	shape map[string]*field
}

func Shape() Schema {
	return &schema{shape: make(map[string]*field)}
}

func (s *schema) Add(name string) SchemaField {
	f := &field{}
	s.shape[name] = f
	return f
}
