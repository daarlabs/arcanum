package tempest

type DisplayClass interface {
	Block(modifiers ...Modifier) Class
	Flex(modifiers ...Modifier) Class
	Grid(modifiers ...Modifier) Class
	Inline(modifiers ...Modifier) Class
	InlineBlock(modifiers ...Modifier) Class
}

func (b *Builder) Block(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "block",
			value:     "block",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Flex(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "flex",
			value:     "flex",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Grid(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "grid",
			value:     "grid",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Inline(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "inline",
			value:     "inline",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) InlineBlock(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "inline-block",
			value:     "inline-block",
			fn:        displayClass,
			modifiers: modifiers,
		},
	)
}
