package tempest

type SpacingClass interface {
	P(size any, modifiers ...Modifier) Class
	Px(size any, modifiers ...Modifier) Class
	Py(size any, modifiers ...Modifier) Class
	Pt(size any, modifiers ...Modifier) Class
	Pr(size any, modifiers ...Modifier) Class
	Pb(size any, modifiers ...Modifier) Class
	Pl(size any, modifiers ...Modifier) Class
	M(size any, modifiers ...Modifier) Class
	Mx(size any, modifiers ...Modifier) Class
	My(size any, modifiers ...Modifier) Class
	Mt(size any, modifiers ...Modifier) Class
	Mr(size any, modifiers ...Modifier) Class
	Mb(size any, modifiers ...Modifier) Class
	Ml(size any, modifiers ...Modifier) Class
}

func (b *Builder) P(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "p-",
			value:     size,
			unit:      Rem,
			fn:        paddingClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Px(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "px-",
			value:     size,
			unit:      Rem,
			fn:        paddingXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Py(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "py-",
			value:     size,
			unit:      Rem,
			fn:        paddingYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Pt(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "pt-",
			value:     size,
			unit:      Rem,
			fn:        paddingTopClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Pr(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "pr-",
			value:     size,
			unit:      Rem,
			fn:        paddingRightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Pb(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "pb-",
			value:     size,
			unit:      Rem,
			fn:        paddingBottomClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Pl(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "pl-",
			value:     size,
			unit:      Rem,
			fn:        paddingLeftClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) M(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "m-",
			value:     size,
			unit:      Rem,
			fn:        marginClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Mx(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "mx-",
			value:     size,
			unit:      Rem,
			fn:        marginXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) My(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "my-",
			value:     size,
			unit:      Rem,
			fn:        marginYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Mt(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "mt-",
			value:     size,
			unit:      Rem,
			fn:        marginTopClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Mr(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "mr-",
			value:     size,
			unit:      Rem,
			fn:        marginRightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Mb(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "mb-",
			value:     size,
			unit:      Rem,
			fn:        marginBottomClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Ml(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "ml-",
			value:     size,
			unit:      Rem,
			fn:        marginLeftClass,
			modifiers: modifiers,
		},
	)
}
