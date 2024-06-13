package tempest

type BackgroundClass interface {
	Bg(name string, code int, modifiers ...Modifier) Tempest
	BgTransparent(modifiers ...Modifier) Tempest
	BgWhite(modifiers ...Modifier) Tempest
	BgSlate(code int, modifiers ...Modifier) Tempest
	BgGray(code int, modifiers ...Modifier) Tempest
	BgZinc(code int, modifiers ...Modifier) Tempest
	BgNeutral(code int, modifiers ...Modifier) Tempest
	BgStone(code int, modifiers ...Modifier) Tempest
	BgRed(code int, modifiers ...Modifier) Tempest
	BgOrange(code int, modifiers ...Modifier) Tempest
	BgAmber(code int, modifiers ...Modifier) Tempest
	BgYellow(code int, modifiers ...Modifier) Tempest
	BgLime(code int, modifiers ...Modifier) Tempest
	BgGreen(code int, modifiers ...Modifier) Tempest
	BgEmerald(code int, modifiers ...Modifier) Tempest
	BgTeal(code int, modifiers ...Modifier) Tempest
	BgCyan(code int, modifiers ...Modifier) Tempest
	BgSky(code int, modifiers ...Modifier) Tempest
	BgBlue(code int, modifiers ...Modifier) Tempest
	BgIndigo(code int, modifiers ...Modifier) Tempest
	BgViolet(code int, modifiers ...Modifier) Tempest
	BgPurple(code int, modifiers ...Modifier) Tempest
	BgFuchsia(code int, modifiers ...Modifier) Tempest
	BgPink(code int, modifiers ...Modifier) Tempest
}

func (b *Builder) Bg(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: createColorPrefix("bg", name, code),
			value:  GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("background-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BgTransparent(modifiers ...Modifier) Tempest {
	return b.Bg(Transparent, 0, modifiers...)
}

func (b *Builder) BgWhite(modifiers ...Modifier) Tempest {
	return b.Bg(White, 0, modifiers...)
}

func (b *Builder) BgSlate(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Slate, code, modifiers...)
}

func (b *Builder) BgGray(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Gray, code, modifiers...)
}

func (b *Builder) BgZinc(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Zinc, code, modifiers...)
}

func (b *Builder) BgNeutral(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Neutral, code, modifiers...)
}

func (b *Builder) BgStone(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Stone, code, modifiers...)
}

func (b *Builder) BgAmber(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Amber, code, modifiers...)
}

func (b *Builder) BgRed(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Red, code, modifiers...)
}

func (b *Builder) BgOrange(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Orange, code, modifiers...)
}

func (b *Builder) BgYellow(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Yellow, code, modifiers...)
}

func (b *Builder) BgLime(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Lime, code, modifiers...)
}

func (b *Builder) BgGreen(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Green, code, modifiers...)
}

func (b *Builder) BgEmerald(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Emerald, code, modifiers...)
}

func (b *Builder) BgTeal(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Teal, code, modifiers...)
}

func (b *Builder) BgCyan(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Cyan, code, modifiers...)
}

func (b *Builder) BgSky(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Sky, code, modifiers...)
}

func (b *Builder) BgBlue(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Blue, code, modifiers...)
}

func (b *Builder) BgIndigo(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Indigo, code, modifiers...)
}

func (b *Builder) BgViolet(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Violet, code, modifiers...)
}

func (b *Builder) BgPurple(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Purple, code, modifiers...)
}

func (b *Builder) BgFuchsia(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Fuchsia, code, modifiers...)
}

func (b *Builder) BgPink(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Pink, code, modifiers...)
}

func (b *Builder) BgRose(code int, modifiers ...Modifier) Tempest {
	return b.Bg(Rose, code, modifiers...)
}
