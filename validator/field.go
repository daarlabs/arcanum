package validator

type SchemaField interface {
	Email() SchemaField
	Required() SchemaField
	Min(min int) SchemaField
	Max(max int) SchemaField
}

func Field() SchemaField {
	return &field{}
}

type field struct {
	min      int
	max      int
	email    bool
	required bool
}

func (f *field) Email() SchemaField {
	f.email = true
	return f
}

func (f *field) Required() SchemaField {
	f.required = true
	return f
}

func (f *field) Min(min int) SchemaField {
	f.min = min
	return f
}

func (f *field) Max(max int) SchemaField {
	f.max = max
	return f
}
