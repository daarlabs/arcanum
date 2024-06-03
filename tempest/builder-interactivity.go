package tempest

type InteractivityClass interface {
}

func (b *Builder) Cursor(cursor string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "cursor-",
			value:     cursor,
			fn:        cursorClass,
			modifiers: modifiers,
		},
	)
}
