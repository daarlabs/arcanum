package tempest

type FlexClass interface {
	FlexSize(size string, modifiers ...Modifier) Tempest
	FlexNone(modifiers ...Modifier) Tempest
	FlexRow(modifiers ...Modifier) Tempest
	FlexCol(modifiers ...Modifier) Tempest
	FlexRowReverse(modifiers ...Modifier) Tempest
	FlexColReverse(modifiers ...Modifier) Tempest
	FlexWrap(modifiers ...Modifier) Tempest
	FlexNoWrap(modifiers ...Modifier) Tempest
}

func (b *Builder) FlexSize(size string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-",
			value:     size,
			fn:        flexClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexNone(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-none",
			value:     "none",
			fn:        flexClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexCol(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-col",
			value:     "column",
			fn:        flexDirectionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexRow(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-row",
			value:     "row",
			fn:        flexDirectionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexRowReverse(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-row-reverse",
			value:     "row-reverse",
			fn:        flexDirectionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexColReverse(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-col-reverse",
			value:     "column-reverse",
			fn:        flexDirectionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexWrap(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-wrap",
			value:     "wrap",
			fn:        flexWrapClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexNoWrap(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex-nowrap",
			value:     "nowrap",
			fn:        flexWrapClass,
			modifiers: modifiers,
		},
	)
}
