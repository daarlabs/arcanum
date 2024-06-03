package tempest

type TransformClass interface {
	Transform(modifiers ...Modifier) Class
	Rotate(size any, modifiers ...Modifier) Class
	TranslateX(size any, modifiers ...Modifier) Class
	TranslateY(size any, modifiers ...Modifier) Class
	ScaleX(size any, modifiers ...Modifier) Class
	ScaleY(size any, modifiers ...Modifier) Class
	SkewX(size any, modifiers ...Modifier) Class
	SkewY(size any, modifiers ...Modifier) Class
}

func (b *Builder) Transform(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "transform",
			fn:        transformRotateClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Rotate(size any, modifiers ...Modifier) Class {
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

func (b *Builder) TranslateX(size any, modifiers ...Modifier) Class {
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

func (b *Builder) TranslateY(size any, modifiers ...Modifier) Class {
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

func (b *Builder) Scale(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "scale-",
			value:     size,
			fn:        transformScaleClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ScaleX(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "scale-x-",
			value:     size,
			fn:        transformScaleXAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ScaleY(size any, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "scale-y-",
			value:     size,
			fn:        transformScaleYAxisClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) SkewX(size any, modifiers ...Modifier) Class {
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

func (b *Builder) SkewY(size any, modifiers ...Modifier) Class {
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
