package tempest

import "fmt"

type EffectClass interface {
	Opacity(opacity float64, modifiers ...Modifier) Tempest
	Shadow(size string, modifiers ...Modifier) Tempest
	ShadowSm(modifiers ...Modifier) Tempest
	ShadowMain(modifiers ...Modifier) Tempest
	ShadowLg(modifiers ...Modifier) Tempest
	ShadowXl(modifiers ...Modifier) Tempest
	ShadowXxl(modifiers ...Modifier) Tempest
	ShadowColor(name string, code int, modifiers ...Modifier) Tempest
}

func (b *Builder) Opacity(opacity float64, modifiers ...Modifier) Tempest {
	if opacity > 1 {
		opacity = opacity / 100
	}
	return b.createStyle(
		style{
			prefix: "opacity-",
			value:  opacity,
			fn: func(selector, _ string) string {
				return opacityClass(selector, opacity)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Shadow(variant string, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "shadow-",
			value:  variant,
			fn: func(selector, value string) string {
				return shadowClass(GlobalConfig.processedShadows, selector, value)
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) ShadowSm(modifiers ...Modifier) Tempest {
	return b.Shadow(SizeSm, modifiers...)
}

func (b *Builder) ShadowMain(modifiers ...Modifier) Tempest {
	return b.Shadow(SizeMain, modifiers...)
}

func (b *Builder) ShadowLg(modifiers ...Modifier) Tempest {
	return b.Shadow(SizeLg, modifiers...)
}

func (b *Builder) ShadowXl(modifiers ...Modifier) Tempest {
	return b.Shadow(SizeXl, modifiers...)
}

func (b *Builder) ShadowXxl(modifiers ...Modifier) Tempest {
	return b.Shadow(SizeXxl, modifiers...)
}

func (b *Builder) ShadowColor(name string, code int, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			priority: 1,
			prefix:   fmt.Sprintf("shadow-%s-%d", name, code),
			value:    GlobalConfig.Color[name][code],
			fn: func(selector, value string) string {
				return shadowColorClass(selector, value, createOpacity(modifiers))
			},
			modifiers: modifiers,
		},
	)
}
