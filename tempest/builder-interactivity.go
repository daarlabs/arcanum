package tempest

type InteractivityClass interface {
	Cursor(cursor string, modifiers ...Modifier) Tempest
	CursorPointer(modifiers ...Modifier) Tempest
	CursorDefault(modifiers ...Modifier) Tempest
	CursorMove(modifiers ...Modifier) Tempest
	PointerEvents(value string, modifiers ...Modifier) Tempest
	PointerEventsNone(modifiers ...Modifier) Tempest
	UserSelect(value string, modifiers ...Modifier) Tempest
	UserSelectNone(modifiers ...Modifier) Tempest
}

func (b *Builder) Cursor(cursor string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "cursor-",
			value:     cursor,
			fn:        cursorClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) CursorPointer(modifiers ...Modifier) Tempest {
	return b.Cursor("pointer", modifiers...)
}

func (b *Builder) CursorDefault(modifiers ...Modifier) Tempest {
	return b.Cursor("default", modifiers...)
}

func (b *Builder) CursorMove(modifiers ...Modifier) Tempest {
	return b.Cursor("move", modifiers...)
}

func (b *Builder) PointerEvents(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "pointer-events-",
			value:     value,
			fn:        pointerEventsClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) PointerEventsNone(modifiers ...Modifier) Tempest {
	return b.PointerEvents("none", modifiers...)
}

func (b *Builder) UserSelect(value string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "user-select-",
			value:     value,
			fn:        userSelectClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) UserSelectNone(modifiers ...Modifier) Tempest {
	return b.UserSelect("none", modifiers...)
}
