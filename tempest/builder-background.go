package tempest

import "fmt"

type BackgroundClass interface {
	BgColor(name string, code int, modifiers ...Modifier) Class
}

func (b *Builder) BgColor(name string, code int, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: fmt.Sprintf("bg-%s-%d", name, code),
			value:  b.Tempest.config.Color[name][code],
			fn: func(selector, value string) string {
				return colorClass("background-color", selector, value, b.createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}
