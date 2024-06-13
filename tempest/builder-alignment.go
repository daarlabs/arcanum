package tempest

type AlignmentClass interface {
	Items(value string, modifiers ...Modifier) Tempest
	ItemsCenter(modifiers ...Modifier) Tempest
	ItemsEnd(modifiers ...Modifier) Tempest
	ItemsStart(modifiers ...Modifier) Tempest
	Justify(value string, modifiers ...Modifier) Tempest
	JustifyCenter(modifiers ...Modifier) Tempest
	JustifyEnd(modifiers ...Modifier) Tempest
	JustifyStart(modifiers ...Modifier) Tempest
	PlaceItems(value string, modifiers ...Modifier) Tempest
	PlaceItemsCenter(modifiers ...Modifier) Tempest
	PlaceItemsEnd(modifiers ...Modifier) Tempest
	PlaceItemsStart(modifiers ...Modifier) Tempest
}

func (b *Builder) Items(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "align-",
			value:  value,
			fn: func(selector, value string) string {
				return alignClass(selector, "items", value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ItemsCenter(modifiers ...Modifier) Tempest {
	return b.Items("center", modifiers...)
}

func (b *Builder) ItemsEnd(modifiers ...Modifier) Tempest {
	return b.Items("end", modifiers...)
}

func (b *Builder) ItemsStart(modifiers ...Modifier) Tempest {
	return b.Items("start", modifiers...)
}

func (b *Builder) Justify(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "justify-",
			value:  value,
			fn: func(selector, value string) string {
				return justifyClass(selector, "content", value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) JustifyCenter(modifiers ...Modifier) Tempest {
	return b.Justify("center", modifiers...)
}

func (b *Builder) JustifyEnd(modifiers ...Modifier) Tempest {
	return b.Justify("end", modifiers...)
}

func (b *Builder) JustifyStart(modifiers ...Modifier) Tempest {
	return b.Justify("start", modifiers...)
}

func (b *Builder) PlaceItems(value string, modifiers ...Modifier) Tempest {
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

func (b *Builder) PlaceItemsCenter(modifiers ...Modifier) Tempest {
	return b.PlaceItems("center", modifiers...)
}

func (b *Builder) PlaceItemsEnd(modifiers ...Modifier) Tempest {
	return b.PlaceItems("end", modifiers...)
}

func (b *Builder) PlaceItemsStart(modifiers ...Modifier) Tempest {
	return b.PlaceItems("start", modifiers...)
}
