package tempest

type DisplayClass interface {
	Block(modifiers ...Modifier) Tempest
	Flex(modifiers ...Modifier) Tempest
	Grid(modifiers ...Modifier) Tempest
	Inline(modifiers ...Modifier) Tempest
	InlineFlex(modifiers ...Modifier) Tempest
	InlineBlock(modifiers ...Modifier) Tempest
}

func (b *Builder) Block(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "block",
			value:     "block",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Flex(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "flex",
			value:     "flex",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Grid(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "grid",
			value:     "grid",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Inline(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inline",
			value:     "inline",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) InlineFlex(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inline-flex",
			value:     "inline-flex",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) InlineBlock(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inline-block",
			value:     "inline-block",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}
