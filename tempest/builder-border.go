package tempest

import "fmt"

type BorderClass interface {
	Border(size int, modifiers ...Modifier) Class
	BorderColor(name string, code int, modifiers ...Modifier) Class
	Rounded(size string, modifiers ...Modifier) Class
}

func (b *Builder) Border(size int, modifiers ...Modifier) Class {
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

func (b *Builder) BorderColor(name string, code int, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: fmt.Sprintf("border-%s-%d", name, code),
			value:  b.Tempest.config.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("border-color", selector, value, b.createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Rounded(size string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix:    "rounded-",
			value:     size,
			fn:        borderRadiusClass,
			modifiers: modifiers,
		},
	)
}
