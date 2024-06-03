package tempest

type SizingClass interface {
	W(size any, modifiers ...Modifier) Class
	MinW(size any, modifiers ...Modifier) Class
	MaxW(size any, modifiers ...Modifier) Class
	H(size any, modifiers ...Modifier) Class
	MinH(size any, modifiers ...Modifier) Class
	MaxH(size any, modifiers ...Modifier) Class
	Size(size any, modifiers ...Modifier) Class
}

func (b *Builder) W(size any, modifiers ...Modifier) Class {
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

func (b *Builder) MinW(size any, modifiers ...Modifier) Class {
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

func (b *Builder) MaxW(size any, modifiers ...Modifier) Class {
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

func (b *Builder) H(size any, modifiers ...Modifier) Class {
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

func (b *Builder) MinH(size any, modifiers ...Modifier) Class {
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

func (b *Builder) MaxH(size any, modifiers ...Modifier) Class {
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

func (b *Builder) Size(size any, modifiers ...Modifier) Class {
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
