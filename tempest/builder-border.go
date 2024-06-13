package tempest

type BorderClass interface {
	Border(size int, modifiers ...Modifier) Tempest
	BorderT(size int, modifiers ...Modifier) Tempest
	BorderR(size int, modifiers ...Modifier) Tempest
	BorderB(size int, modifiers ...Modifier) Tempest
	BorderL(size int, modifiers ...Modifier) Tempest
	BorderColor(name string, code int, modifiers ...Modifier) Tempest
	BorderTColor(name string, code int, modifiers ...Modifier) Tempest
	BorderRColor(name string, code int, modifiers ...Modifier) Tempest
	BorderBColor(name string, code int, modifiers ...Modifier) Tempest
	BorderLColor(name string, code int, modifiers ...Modifier) Tempest
	BorderTransparent(modifiers ...Modifier) Tempest
	BorderWhite(modifiers ...Modifier) Tempest
	BorderSlate(code int, modifiers ...Modifier) Tempest
	BorderGray(code int, modifiers ...Modifier) Tempest
	BorderZinc(code int, modifiers ...Modifier) Tempest
	BorderNeutral(code int, modifiers ...Modifier) Tempest
	BorderStone(code int, modifiers ...Modifier) Tempest
	BorderRed(code int, modifiers ...Modifier) Tempest
	BorderOrange(code int, modifiers ...Modifier) Tempest
	BorderAmber(code int, modifiers ...Modifier) Tempest
	BorderYellow(code int, modifiers ...Modifier) Tempest
	BorderLime(code int, modifiers ...Modifier) Tempest
	BorderGreen(code int, modifiers ...Modifier) Tempest
	BorderEmerald(code int, modifiers ...Modifier) Tempest
	BorderTeal(code int, modifiers ...Modifier) Tempest
	BorderCyan(code int, modifiers ...Modifier) Tempest
	BorderSky(code int, modifiers ...Modifier) Tempest
	BorderBlue(code int, modifiers ...Modifier) Tempest
	BorderIndigo(code int, modifiers ...Modifier) Tempest
	BorderViolet(code int, modifiers ...Modifier) Tempest
	BorderPurple(code int, modifiers ...Modifier) Tempest
	BorderFuchsia(code int, modifiers ...Modifier) Tempest
	BorderPink(code int, modifiers ...Modifier) Tempest
	BorderRose(code int, modifiers ...Modifier) Tempest
	BorderRadius(size string, modifiers ...Modifier) Tempest
	Rounded(modifiers ...Modifier) Tempest
	RoundedSm(modifiers ...Modifier) Tempest
	RoundedLg(modifiers ...Modifier) Tempest
	RoundedXl(modifiers ...Modifier) Tempest
	RoundedFull(modifiers ...Modifier) Tempest
}

func (b *Builder) Border(size int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "border-",
			value:     size,
			unit:      Px,
			fn:        borderWidthClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderT(size int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "border-t-",
			value:  size,
			unit:   Px,
			fn: func(selector, value string) string {
				return borderDirectionWidthClass("top", selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderR(size int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "border-r-",
			value:  size,
			unit:   Px,
			fn: func(selector, value string) string {
				return borderDirectionWidthClass("right", selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderB(size int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "border-b-",
			value:  size,
			unit:   Px,
			fn: func(selector, value string) string {
				return borderDirectionWidthClass("bottom", selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderL(size int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "border-l-",
			value:  size,
			unit:   Px,
			fn: func(selector, value string) string {
				return borderDirectionWidthClass("left", selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: createColorPrefix("border", name, code),
			value:  GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderTColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			priority: 1,
			prefix:   createColorPrefix("border-top", name, code),
			value:    GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-top-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderRColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			priority: 1,
			prefix:   createColorPrefix("border-right", name, code),
			value:    GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-right-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderBColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			priority: 1,
			prefix:   createColorPrefix("border-bottom", name, code),
			value:    GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-bottom-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderLColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			priority: 1,
			prefix:   createColorPrefix("border-left", name, code),
			value:    GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-left-color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BorderTransparent(modifiers ...Modifier) Tempest {
	return b.BorderColor(Transparent, 0, modifiers...)
}

func (b *Builder) BorderWhite(modifiers ...Modifier) Tempest {
	return b.BorderColor(White, 0, modifiers...)
}

func (b *Builder) BorderSlate(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Slate, code, modifiers...)
}

func (b *Builder) BorderGray(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Gray, code, modifiers...)
}

func (b *Builder) BorderZinc(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Zinc, code, modifiers...)
}

func (b *Builder) BorderNeutral(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Neutral, code, modifiers...)
}

func (b *Builder) BorderStone(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Stone, code, modifiers...)
}

func (b *Builder) BorderAmber(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Amber, code, modifiers...)
}

func (b *Builder) BorderRed(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Red, code, modifiers...)
}

func (b *Builder) BorderOrange(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Orange, code, modifiers...)
}

func (b *Builder) BorderYellow(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Yellow, code, modifiers...)
}

func (b *Builder) BorderLime(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Lime, code, modifiers...)
}

func (b *Builder) BorderGreen(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Green, code, modifiers...)
}

func (b *Builder) BorderEmerald(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Emerald, code, modifiers...)
}

func (b *Builder) BorderTeal(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Teal, code, modifiers...)
}

func (b *Builder) BorderCyan(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Cyan, code, modifiers...)
}

func (b *Builder) BorderSky(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Sky, code, modifiers...)
}

func (b *Builder) BorderBlue(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Blue, code, modifiers...)
}

func (b *Builder) BorderIndigo(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Indigo, code, modifiers...)
}

func (b *Builder) BorderViolet(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Violet, code, modifiers...)
}

func (b *Builder) BorderPurple(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Purple, code, modifiers...)
}

func (b *Builder) BorderFuchsia(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Fuchsia, code, modifiers...)
}

func (b *Builder) BorderPink(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Pink, code, modifiers...)
}

func (b *Builder) BorderRose(code int, modifiers ...Modifier) Tempest {
	return b.BorderColor(Rose, code, modifiers...)
}

func (b *Builder) BorderRadius(size string, modifiers ...Modifier) Tempest {
	prefix := "rounded"
	if size != "" {
		prefix += "-" + size
	}
	if size == "" {
		size = SizeMain
	}
	return b.createStyle(
		style{
			prefix:    prefix,
			value:     size,
			fn:        borderRadiusClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Rounded(modifiers ...Modifier) Tempest {
	return b.BorderRadius("", modifiers...)
}

func (b *Builder) RoundedSm(modifiers ...Modifier) Tempest {
	return b.BorderRadius(SizeSm, modifiers...)
}

func (b *Builder) RoundedLg(modifiers ...Modifier) Tempest {
	return b.BorderRadius(SizeLg, modifiers...)
}

func (b *Builder) RoundedXl(modifiers ...Modifier) Tempest {
	return b.BorderRadius(SizeXl, modifiers...)
}

func (b *Builder) RoundedFull(modifiers ...Modifier) Tempest {
	return b.BorderRadius(Full, modifiers...)
}
