package tempest

type SizingClass interface {
	W(size any, modifiers ...Modifier) Tempest
	MinW(size any, modifiers ...Modifier) Tempest
	MaxW(size any, modifiers ...Modifier) Tempest
	H(size any, modifiers ...Modifier) Tempest
	MinH(size any, modifiers ...Modifier) Tempest
	MaxH(size any, modifiers ...Modifier) Tempest
	Size(size any, modifiers ...Modifier) Tempest
}

func (b *Builder) W(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "w-",
			value:     size,
			unit:      Rem,
			fn:        widthClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) MinW(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "min-w-",
			value:     size,
			unit:      Rem,
			fn:        minWidthClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) MaxW(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "max-w-",
			value:     size,
			unit:      Rem,
			fn:        maxWidthClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) H(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "h-",
			value:     size,
			unit:      Rem,
			fn:        heightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) MinH(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "min-h-",
			value:     size,
			unit:      Rem,
			fn:        minHeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) MaxH(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "max-h-",
			value:     size,
			unit:      Rem,
			fn:        maxHeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Size(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "size-",
			value:     size,
			unit:      Rem,
			fn:        sizeClass,
			modifiers: modifiers,
		},
	)
}
