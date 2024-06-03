package tempest

type SpecialClass interface {
	Group(modifiers ...Modifier) Class
	Peer(modifiers ...Modifier) Class
}

func (b *Builder) Group(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "group",
			fn: func(selector, value string) string {
				return ""
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Peer(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "peer",
			fn: func(selector, value string) string {
				return ""
			},
			modifiers: modifiers,
		},
	)
}
