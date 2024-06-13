package tempest

type TransformClass interface {
	Transform(modifiers ...Modifier) Tempest
	Rotate(size any, modifiers ...Modifier) Tempest
	TranslateX(size any, modifiers ...Modifier) Tempest
	TranslateY(size any, modifiers ...Modifier) Tempest
	Scale(size any, modifiers ...Modifier) Tempest
	ScaleX(size any, modifiers ...Modifier) Tempest
	ScaleY(size any, modifiers ...Modifier) Tempest
	SkewX(size any, modifiers ...Modifier) Tempest
	SkewY(size any, modifiers ...Modifier) Tempest
	Origin(position string, modifiers ...Modifier) Tempest
}

func (b *Builder) Transform(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "transform",
			fn:        transformRotateClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Rotate(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "rotate-",
			value:     size,
			unit:      Deg,
			fn:        transformRotateClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TranslateX(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "translate-x-",
			value:     size,
			unit:      Rem,
			fn:        transformTranslateXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TranslateY(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "translate-y-",
			value:     size,
			unit:      Rem,
			fn:        transformTranslateYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Scale(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "scale-",
			value:     size,
			fn:        transformScaleClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ScaleX(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "scale-x-",
			value:     size,
			fn:        transformScaleXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ScaleY(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "scale-y-",
			value:     size,
			fn:        transformScaleYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) SkewX(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "skew-x-",
			value:     size,
			unit:      Deg,
			fn:        transformSkewXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) SkewY(size any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "skew-y-",
			value:     size,
			unit:      Deg,
			fn:        transformSkewYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Origin(position string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "origin-",
			value:     position,
			fn:        transformOriginClass,
			modifiers: modifiers,
		},
	)
}
