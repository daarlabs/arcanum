package tempest

type SpecialClass interface {
	Group(modifiers ...Modifier) Tempest
	Peer(modifiers ...Modifier) Tempest
	Dark(modifiers ...Modifier) Tempest
}

func (b *Builder) Group(modifiers ...Modifier) Tempest {
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

func (b *Builder) Peer(modifiers ...Modifier) Tempest {
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

func (b *Builder) Dark(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "dark",
			fn: func(selector, value string) string {
				return ""
			},
			modifiers: modifiers,
		},
	)
}
