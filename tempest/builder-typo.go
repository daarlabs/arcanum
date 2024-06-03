package tempest

import "fmt"

type TypoClass interface {
	FontColor(name string, code int, modifiers ...Modifier) Class
	FontFamily(name string, modifiers ...Modifier) Class
	FontSize(size string, modifiers ...Modifier) Class
	FontThin(modifiers ...Modifier) Class
	FontExtralight(modifiers ...Modifier) Class
	FontLight(modifiers ...Modifier) Class
	FontNormal(modifiers ...Modifier) Class
	FontMedium(modifiers ...Modifier) Class
	FontSemibold(modifiers ...Modifier) Class
	FontBold(modifiers ...Modifier) Class
	FontExtrabold(modifiers ...Modifier) Class
	FontBlack(modifiers ...Modifier) Class
	TextAlign(position string, modifiers ...Modifier) Class
	TextDecoration(decoration string, modifiers ...Modifier) Class
	Truncate(modifiers ...Modifier) Class
}

func (b *Builder) FontColor(name string, code int, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: fmt.Sprintf("font-%s-%d", name, code),
			value:  b.Tempest.config.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("color", selector, value, b.createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontFamily(name string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-" + name,
			value:     b.Tempest.config.Font[name].Value,
			fn:        fontFamilyClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontSize(size string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-",
			value:     size,
			fn:        fontSizeClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontThin(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-thin",
			value:     "100",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontExtralight(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-extralight",
			value:     "200",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontLight(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-light",
			value:     "300",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontNormal(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-normal",
			value:     "400",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontMedium(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-normal",
			value:     "500",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontSemibold(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-semibold",
			value:     "600",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontBold(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-bold",
			value:     "700",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontExtrabold(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-bold",
			value:     "800",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) FontBlack(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "font-black",
			value:     "900",
			fn:        fontWeightClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextAlign(position string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "text-align-",
			value:     position,
			fn:        textAlignClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) TextDecoration(decoration string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "text-decoration-",
			value:     decoration,
			fn:        textDecorationClass,
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Truncate(modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "truncate",
			value:     "",
			fn:        truncateClass,
			modifiers: modifiers,
		},
	)
}
