package config

type Prefix struct {
	Name string
	Path any
}

type Router struct {
	Prefix  Prefix
	Recover bool
}
