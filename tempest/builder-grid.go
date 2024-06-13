package tempest

import (
	"reflect"
)

type GridClass interface {
	GridCols(size any, modifiers ...Modifier) Tempest
	GridRows(size any, modifiers ...Modifier) Tempest
	Gap(size any, modifiers ...Modifier) Tempest
	Order(index int, modifiers ...Modifier) Tempest
}

func (b *Builder) GridCols(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "grid-cols-",
			value:  size,
			fn: func(selector, value string) string {
				return gridTemplateColumnsClass(selector, value, reflect.TypeOf(size).Kind() == reflect.String)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) GridRows(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "grid-rows-",
			value:  size,
			fn: func(selector, value string) string {
				return gridTemplateRowsClass(selector, value, reflect.TypeOf(size).Kind() == reflect.String)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Gap(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "gap-",
			value:     size,
			unit:      Rem,
			fn:        gapClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Order(index int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "order-",
			value:     index,
			fn:        orderClass,
			modifiers: modifiers,
		},
	)
}
