package tempest

import (
	"strings"
)

type Class interface {
	AlignmentClass
	BackgroundClass
	BorderClass
	DisplayClass
	EffectClass
	FlexClass
	GridClass
	InteractivityClass
	LayoutClass
	SizingClass
	SpacingClass
	SpecialClass
	TransformClass
	TransitionClass
	TypoClass
	String() string
}

type Builder struct {
	*Context
	classes []string
}

type style struct {
	prefix    string
	value     any
	unit      string
	fn        func(selector, value string) string
	modifiers []Modifier
}

func (b *Builder) String() string {
	return strings.Join(b.classes, " ")
}

func (b *Builder) createStyle(s style) *Builder {
	var suffix string
	shouldHaveSuffix := strings.HasSuffix(s.prefix, "-")
	validatedValue := s.value
	if s.unit == Rem {
		validatedValue = convertSizeToRem(b.Tempest.config.FontSize, s.value)
	}
	value := createValue(validatedValue, s.unit)
	if shouldHaveSuffix {
		suffix = createSuffix(s.value)
	}
	if strings.HasPrefix(value, "-") {
		s.prefix = "-" + s.prefix
	}
	class := applyClassModifiers(s.prefix+suffix, s.modifiers...)
	b.classes = append(b.classes, class)
	if b.Tempest.classExists(class) {
		return b
	}
	if shouldHaveSuffix {
		suffix = escape(suffix)
	}
	selector := applySelectorModifiers(s.prefix+suffix, s.modifiers...)
	b.Add(class, applyBreakpointModifiers(b.Tempest.config.Breakpoint, s.fn(selector, value), s.modifiers...))
	return b
}

func (b *Builder) createOpacity(modifiers []Modifier) float64 {
	opacity := 100
	for _, modifier := range modifiers {
		if modifier.Name != opacityModifier {
			continue
		}
		opacity = modifier.Value.(int)
		break
	}
	return float64(opacity) / float64(100)
}
