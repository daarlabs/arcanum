package tempest

type AlignmentClass interface {
	AlignItems(value string, modifiers ...Modifier) Class
	JustifyContent(value string, modifiers ...Modifier) Class
	PlaceItems(value string, modifiers ...Modifier) Class
}

func (b *Builder) AlignItems(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "align-items-",
			value:  value,
			fn: func(selector, value string) string {
				return alignClass(selector, "items", value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) JustifyContent(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "justify-content-",
			value:  value,
			fn: func(selector, value string) string {
				return justifyClass(selector, "content", value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) PlaceItems(value string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "place-items-",
			value:  value,
			fn: func(selector, value string) string {
				return placeClass(selector, "items", value)
			},
			modifiers: modifiers,
		},
	)
}
