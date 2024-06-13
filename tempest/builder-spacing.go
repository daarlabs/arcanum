package tempest

type SpacingClass interface {
	P(size any, modifiers ...Modifier) Tempest
	Px(size any, modifiers ...Modifier) Tempest
	Py(size any, modifiers ...Modifier) Tempest
	Pt(size any, modifiers ...Modifier) Tempest
	Pr(size any, modifiers ...Modifier) Tempest
	Pb(size any, modifiers ...Modifier) Tempest
	Pl(size any, modifiers ...Modifier) Tempest
	M(size any, modifiers ...Modifier) Tempest
	Mx(size any, modifiers ...Modifier) Tempest
	My(size any, modifiers ...Modifier) Tempest
	Mt(size any, modifiers ...Modifier) Tempest
	Mr(size any, modifiers ...Modifier) Tempest
	Mb(size any, modifiers ...Modifier) Tempest
	Ml(size any, modifiers ...Modifier) Tempest
}

func (b *Builder) P(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Px(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Py(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Pt(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Pr(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Pb(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Pl(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) M(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Mx(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) My(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Mt(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Mr(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Mb(size any, modifiers ...Modifier) Tempest {
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

func (b *Builder) Ml(size any, modifiers ...Modifier) Tempest {
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
