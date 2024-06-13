package tempest

type TransitionClass interface {
	Transition(modifiers ...Modifier) Tempest
}

func (b *Builder) Transition(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "transition",
			fn:        transitionClass,
			modifiers: modifiers,
		},
	)
}
