package tempest

type LayoutClass interface {
	Container() Tempest
	Overflow(value string, modifiers ...Modifier) Tempest
	OverflowX(value string, modifiers ...Modifier) Tempest
	OverflowY(value string, modifiers ...Modifier) Tempest
	Position(value string, modifiers ...Modifier) Tempest
	Absolute(modifiers ...Modifier) Tempest
	Relative(modifiers ...Modifier) Tempest
	Static(modifiers ...Modifier) Tempest
	Fixed(modifiers ...Modifier) Tempest
	Sticky(modifiers ...Modifier) Tempest
	Top(value any, modifiers ...Modifier) Tempest
	Right(value any, modifiers ...Modifier) Tempest
	Bottom(value any, modifiers ...Modifier) Tempest
	Left(value any, modifiers ...Modifier) Tempest
	Inset(value any, modifiers ...Modifier) Tempest
	InsetX(value any, modifiers ...Modifier) Tempest
	InsetY(value any, modifiers ...Modifier) Tempest
	Z(index int, modifiers ...Modifier) Tempest
	Visible(modifiers ...Modifier) Tempest
	Invisible(modifiers ...Modifier) Tempest
}

func (b *Builder) Container() Tempest {
	return b.createStyle(
		style{
			prefix: "container",
			fn: func(_, _ string) string {
				return containerClass(GlobalConfig.Breakpoint, GlobalConfig.Container)
			},
		},
	)
}

func (b *Builder) Overflow(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "overflow-",
			value:     value,
			fn:        overflowClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) OverflowX(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "overflow-x-",
			value:     value,
			fn:        overflowXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) OverflowY(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "overflow-y-",
			value:     value,
			fn:        overflowYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Position(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    value,
			value:     value,
			fn:        positionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Absolute(modifiers ...Modifier) Tempest {
	return b.Position("absolute", modifiers...)
}

func (b *Builder) Relative(modifiers ...Modifier) Tempest {
	return b.Position("relative", modifiers...)
}

func (b *Builder) Fixed(modifiers ...Modifier) Tempest {
	return b.Position("fixed", modifiers...)
}

func (b *Builder) Sticky(modifiers ...Modifier) Tempest {
	return b.Position("sticky", modifiers...)
}

func (b *Builder) Static(modifiers ...Modifier) Tempest {
	return b.Position("static", modifiers...)
}

func (b *Builder) Top(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "top-",
			value:     value,
			unit:      Rem,
			fn:        topClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Right(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "right-",
			value:     value,
			unit:      Rem,
			fn:        rightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Bottom(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "bottom-",
			value:     value,
			unit:      Rem,
			fn:        bottomClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Left(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "left-",
			value:     value,
			unit:      Rem,
			fn:        leftClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Inset(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inset-",
			value:     value,
			unit:      Rem,
			fn:        insetClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) InsetX(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inset-x-",
			value:     value,
			unit:      Rem,
			fn:        insetXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) InsetY(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "inset-y-",
			value:     value,
			unit:      Rem,
			fn:        insetYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Z(index int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "z-",
			value:     index,
			fn:        zIndexClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Visible(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "visible",
			value:     "visible",
			fn:        visibilityClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Invisible(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "invisible",
			value:     "hidden",
			fn:        visibilityClass,
			modifiers: modifiers,
		},
	)
}
