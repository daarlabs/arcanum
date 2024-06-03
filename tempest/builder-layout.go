package tempest

type LayoutClass interface {
	Container() Class
	Overflow(value string, modifiers ...Modifier) Class
	OverflowX(value string, modifiers ...Modifier) Class
	OverflowY(value string, modifiers ...Modifier) Class
	Position(value string, modifiers ...Modifier) Class
	Top(value any, modifiers ...Modifier) Class
	Right(value any, modifiers ...Modifier) Class
	Bottom(value any, modifiers ...Modifier) Class
	Left(value any, modifiers ...Modifier) Class
	Inset(value any, modifiers ...Modifier) Class
	InsetX(value any, modifiers ...Modifier) Class
	InsetY(value any, modifiers ...Modifier) Class
}

func (b *Builder) Container() Class {
	return b.createStyle(
		style{
			prefix: "container",
			fn: func(_, _ string) string {
				return containerClass(b.Tempest.config.Breakpoint, b.Tempest.config.Container)
			},
		},
	)
}

func (b *Builder) Overflow(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "overflow-",
			value:     value,
			fn:        overflowClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) OverflowX(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "overflow-x-",
			value:     value,
			fn:        overflowXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) OverflowY(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "overflow-y-",
			value:     value,
			fn:        overflowYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Position(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    value,
			value:     value,
			fn:        positionClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Top(value any, modifiers ...Modifier) Class {
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

func (b *Builder) Right(value any, modifiers ...Modifier) Class {
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

func (b *Builder) Bottom(value any, modifiers ...Modifier) Class {
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

func (b *Builder) Left(value any, modifiers ...Modifier) Class {
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

func (b *Builder) Inset(value any, modifiers ...Modifier) Class {
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

func (b *Builder) InsetX(value any, modifiers ...Modifier) Class {
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

func (b *Builder) InsetY(value any, modifiers ...Modifier) Class {
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
