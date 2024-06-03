package tempest

type TransitionClass interface {
	Transition(modifiers ...Modifier) Class
}

func (b *Builder) Transition(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "transition",
			fn:        transitionClass,
			modifiers: modifiers,
		},
	)
}
