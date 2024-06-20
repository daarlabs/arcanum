package tempest

import "fmt"

type FilterClass interface {
	Blur(blur float64, modifiers ...Modifier) Tempest
	Brightness(brightness float64, modifiers ...Modifier) Tempest
}

func (b *Builder) Blur(blur float64, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "blur-",
			value:  blur,
			unit:   Px,
			fn: func(selector, value string) string {
				return filterClass(selector, fmt.Sprintf("blur(%s)", value))
			},
			modifiers: modifiers,
		},
	)
}

func (b *Builder) Brightness(brightness float64, modifiers ...Modifier) Tempest {
	return b.createStyle(
		style{
			prefix: "brightness-",
			value:  createMostSuitableNumber(brightness),
			fn: func(selector, value string) string {
				return filterClass(selector, fmt.Sprintf("brightness(%s)", value))
			},
			modifiers: modifiers,
		},
	)
}
