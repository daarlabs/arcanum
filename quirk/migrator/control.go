package migrator

import "github.com/daarlabs/arcanum/quirk"

type Control interface {
	DB(name ...string) *quirk.DB
}

type control struct {
	*migrator
}

func (c *control) DB(name ...string) *quirk.DB {
	n := mainDbname
	if len(name) > 0 {
		n = name[0]
	}
	d, ok := c.databases[n]
	if !ok {
		panic(ErrorInvalidDatabase)
	}
	return d
}
