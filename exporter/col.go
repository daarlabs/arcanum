package exporter

type Col interface {
	Value(value any) Col
}

type col struct {
	value any
}

func (c *col) Value(value any) Col {
	c.value = value
	return c
}
