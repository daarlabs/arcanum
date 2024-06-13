package mirage

import "github.com/daarlabs/arcanum/gox"

type layoutFactory = func(c Ctx, nodes ...gox.Node) gox.Node

type LayoutManager interface {
	Add(name string, layout layoutFactory) LayoutManager
}

type layout struct {
	factories map[string]layoutFactory
}

func createLayout() *layout {
	return &layout{
		factories: make(map[string]layoutFactory),
	}
}

func (l *layout) Add(name string, layout layoutFactory) LayoutManager {
	l.factories[name] = layout
	return l
}
