package config

type Prefix struct {
	Name   string
	Proxy  string
	Cookie string
	Path   any
}

type Router struct {
	Prefix  Prefix
	Recover bool
}
