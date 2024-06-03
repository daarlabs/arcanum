package tempest

type FlexClass interface {
	FlexSize(size string, modifiers ...Modifier) Class
}

func (b *Builder) FlexSize(size string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "flex-",
			value:     size,
			fn:        flexClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FlexNone(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "flex-none",
			value:     "none",
			fn:        flexClass,
			modifiers: modifiers,
		},
	)
}
