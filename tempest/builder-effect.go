package tempest

import "fmt"

type EffectClass interface {
	Shadow(size string, modifiers ...Modifier) Class
	ShadowColor(name string, code int, modifiers ...Modifier) Class
}

func (b *Builder) Shadow(size string, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: "shadow-",
			value:  size,
			fn: func(selector, value string) string {
				return shadowClass(b.Tempest.config.processedShadows, selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ShadowColor(name string, code int, modifiers ...Modifier) Class {
	return b.createStyle(
		style{
			prefix: fmt.Sprintf("shadow-%s-%d", name, code),
			value:  b.Tempest.config.Color[name][code],
			fn: func(selector, value string) string {
				return shadowColorClass(selector, value, b.createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}
