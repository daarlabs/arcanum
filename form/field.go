package form

type Field[T any] struct {
	Type      string
	DataType  string
	Id        string
	Name      string
	Label     string
	Text      string
	Value     T
	Messages  []string
	Autofocus bool
	Disabled  bool
	Required  bool
	Multiple  bool
}
