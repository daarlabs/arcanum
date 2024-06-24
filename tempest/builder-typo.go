package tempest

type TypoClass interface {
	TextTransparent(modifiers ...Modifier) Tempest
	Text(name string, code int, modifiers ...Modifier) Tempest
	TextWhite(modifiers ...Modifier) Tempest
	TextSlate(code int, modifiers ...Modifier) Tempest
	TextGray(code int, modifiers ...Modifier) Tempest
	TextZinc(code int, modifiers ...Modifier) Tempest
	TextNeutral(code int, modifiers ...Modifier) Tempest
	TextStone(code int, modifiers ...Modifier) Tempest
	TextRed(code int, modifiers ...Modifier) Tempest
	TextOrange(code int, modifiers ...Modifier) Tempest
	TextAmber(code int, modifiers ...Modifier) Tempest
	TextYellow(code int, modifiers ...Modifier) Tempest
	TextLime(code int, modifiers ...Modifier) Tempest
	TextGreen(code int, modifiers ...Modifier) Tempest
	TextEmerald(code int, modifiers ...Modifier) Tempest
	TextTeal(code int, modifiers ...Modifier) Tempest
	TextCyan(code int, modifiers ...Modifier) Tempest
	TextSky(code int, modifiers ...Modifier) Tempest
	TextBlue(code int, modifiers ...Modifier) Tempest
	TextIndigo(code int, modifiers ...Modifier) Tempest
	TextViolet(code int, modifiers ...Modifier) Tempest
	TextPurple(code int, modifiers ...Modifier) Tempest
	TextFuchsia(code int, modifiers ...Modifier) Tempest
	TextPink(code int, modifiers ...Modifier) Tempest
	TextRose(code int, modifiers ...Modifier) Tempest
	FontFamily(name string, modifiers ...Modifier) Tempest
	TextSize(size string, modifiers ...Modifier) Tempest
	TextXs(modifiers ...Modifier) Tempest
	TextSm(modifiers ...Modifier) Tempest
	TextMain(modifiers ...Modifier) Tempest
	TextLg(modifiers ...Modifier) Tempest
	TextXl(modifiers ...Modifier) Tempest
	TextXxl(modifiers ...Modifier) Tempest
	FontThin(modifiers ...Modifier) Tempest
	FontExtralight(modifiers ...Modifier) Tempest
	FontLight(modifiers ...Modifier) Tempest
	FontNormal(modifiers ...Modifier) Tempest
	FontMedium(modifiers ...Modifier) Tempest
	FontSemibold(modifiers ...Modifier) Tempest
	FontBold(modifiers ...Modifier) Tempest
	FontExtrabold(modifiers ...Modifier) Tempest
	FontBlack(modifiers ...Modifier) Tempest
	TextLeft(modifiers ...Modifier) Tempest
	TextCenter(modifiers ...Modifier) Tempest
	TextRight(modifiers ...Modifier) Tempest
	Underline(modifiers ...Modifier) Tempest
	Overline(modifiers ...Modifier) Tempest
	NoUnderline(modifiers ...Modifier) Tempest
	LineThrough(modifiers ...Modifier) Tempest
	Truncate(modifiers ...Modifier) Tempest
	Lh(value any, modifiers ...Modifier) Tempest
	LhNone(modifiers ...Modifier) Tempest
	LhRelax(modifiers ...Modifier) Tempest
	LhLoose(modifiers ...Modifier) Tempest
	BreakAll(modifiers ...Modifier) Tempest
}

func (b *Builder) Text(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: createColorPrefix("text", name, code),
			value:  GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("color", selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextTransparent(modifiers ...Modifier) Tempest {
	return b.Text(Transparent, 0, modifiers...)
}

func (b *Builder) TextWhite(modifiers ...Modifier) Tempest {
	return b.Text(White, 0, modifiers...)
}

func (b *Builder) TextSlate(code int, modifiers ...Modifier) Tempest {
	return b.Text(Slate, code, modifiers...)
}

func (b *Builder) TextGray(code int, modifiers ...Modifier) Tempest {
	return b.Text(Gray, code, modifiers...)
}

func (b *Builder) TextZinc(code int, modifiers ...Modifier) Tempest {
	return b.Text(Zinc, code, modifiers...)
}

func (b *Builder) TextNeutral(code int, modifiers ...Modifier) Tempest {
	return b.Text(Neutral, code, modifiers...)
}

func (b *Builder) TextStone(code int, modifiers ...Modifier) Tempest {
	return b.Text(Stone, code, modifiers...)
}

func (b *Builder) TextAmber(code int, modifiers ...Modifier) Tempest {
	return b.Text(Amber, code, modifiers...)
}

func (b *Builder) TextRed(code int, modifiers ...Modifier) Tempest {
	return b.Text(Red, code, modifiers...)
}

func (b *Builder) TextOrange(code int, modifiers ...Modifier) Tempest {
	return b.Text(Orange, code, modifiers...)
}

func (b *Builder) TextYellow(code int, modifiers ...Modifier) Tempest {
	return b.Text(Yellow, code, modifiers...)
}

func (b *Builder) TextLime(code int, modifiers ...Modifier) Tempest {
	return b.Text(Lime, code, modifiers...)
}

func (b *Builder) TextGreen(code int, modifiers ...Modifier) Tempest {
	return b.Text(Green, code, modifiers...)
}

func (b *Builder) TextEmerald(code int, modifiers ...Modifier) Tempest {
	return b.Text(Emerald, code, modifiers...)
}

func (b *Builder) TextTeal(code int, modifiers ...Modifier) Tempest {
	return b.Text(Teal, code, modifiers...)
}

func (b *Builder) TextCyan(code int, modifiers ...Modifier) Tempest {
	return b.Text(Cyan, code, modifiers...)
}

func (b *Builder) TextSky(code int, modifiers ...Modifier) Tempest {
	return b.Text(Sky, code, modifiers...)
}

func (b *Builder) TextBlue(code int, modifiers ...Modifier) Tempest {
	return b.Text(Blue, code, modifiers...)
}

func (b *Builder) TextIndigo(code int, modifiers ...Modifier) Tempest {
	return b.Text(Indigo, code, modifiers...)
}

func (b *Builder) TextViolet(code int, modifiers ...Modifier) Tempest {
	return b.Text(Violet, code, modifiers...)
}

func (b *Builder) TextPurple(code int, modifiers ...Modifier) Tempest {
	return b.Text(Purple, code, modifiers...)
}

func (b *Builder) TextFuchsia(code int, modifiers ...Modifier) Tempest {
	return b.Text(Fuchsia, code, modifiers...)
}

func (b *Builder) TextPink(code int, modifiers ...Modifier) Tempest {
	return b.Text(Pink, code, modifiers...)
}

func (b *Builder) TextRose(code int, modifiers ...Modifier) Tempest {
	return b.Text(Rose, code, modifiers...)
}

func (b *Builder) FontFamily(name string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-" + name,
			value:     GlobalConfig.Font[name].Value,
			fn:        fontFamilyClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextSize(size string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-",
			value:     size,
			fn:        fontSizeClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextXs(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeXs, modifiers...)
}

func (b *Builder) TextSm(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeSm, modifiers...)
}

func (b *Builder) TextMain(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeMain, modifiers...)
}

func (b *Builder) TextLg(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeLg, modifiers...)
}

func (b *Builder) TextXl(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeXl, modifiers...)
}

func (b *Builder) TextXxl(modifiers ...Modifier) Tempest {
	return b.TextSize(SizeXxl, modifiers...)
}

func (b *Builder) FontThin(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-thin",
			value:     "100",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontExtralight(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-extralight",
			value:     "200",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontLight(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-light",
			value:     "300",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontNormal(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-normal",
			value:     "400",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontMedium(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-normal",
			value:     "500",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontSemibold(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-semibold",
			value:     "600",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontBold(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-bold",
			value:     "700",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontExtrabold(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-bold",
			value:     "800",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontBlack(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "font-black",
			value:     "900",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextLeft(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "text-left",
			value:     "left",
			fn:        textAlignClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextCenter(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "text-center",
			value:     "center",
			fn:        textAlignClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextRight(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "text-right",
			value:     "right",
			fn:        textAlignClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Underline(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "underline",
			value:     "underline",
			fn:        textDecorationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) NoUnderline(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "no-underline",
			value:     "none",
			fn:        textDecorationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Overline(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "overline",
			value:     "overline",
			fn:        textDecorationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) LineThrough(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "line-through",
			value:     "line-through",
			fn:        textDecorationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Truncate(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "truncate",
			value:     "",
			fn:        truncateClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) BreakAll(modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "break-all",
			value:     "break-all",
			fn:        wordBreakClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Lh(value any, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix:    "lh-",
			value:     value,
			unit:      Rem,
			fn:        lineHeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) LhNone(modifiers ...Modifier) Tempest {
	return b.Lh("none", modifiers...)
}

func (b *Builder) LhRelax(modifiers ...Modifier) Tempest {
	return b.Lh("relax", modifiers...)
}

func (b *Builder) LhLoose(modifiers ...Modifier) Tempest {
	return b.Lh("loose", modifiers...)
}
