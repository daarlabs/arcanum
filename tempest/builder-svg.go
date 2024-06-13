package tempest

type SvgClass interface {
	FillCurrent(modifiers ...Modifier) Tempest
	StrokeCurrent(modifiers ...Modifier) Tempest
}

func (b *Builder) FillCurrent(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "fill-current",
			value:     "currentColor",
			fn:        fillClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) StrokeCurrent(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "stroke-current",
			value:     "currentColor",
			fn:        strokeClass,
			modifiers: modifiers,
		},
	)
}
