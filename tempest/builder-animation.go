package tempest

type AnimationClass interface {
	Animate(name string, modifiers ...Modifier) Tempest
	Spin(modifiers ...Modifier) Tempest
}

func (b *Builder) Animate(name string, modifiers ...Modifier) Tempest {
	animation, ok := GlobalConfig.Animation[name]
	if !ok {
		return b
	}
	return b.createStyle(
		style{
			prefix:    "animate-" + name,
			value:     animation.Value,
			fn:        animationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Spin(modifiers ...Modifier) Tempest {
	return b.Animate("spin", modifiers...)
}
