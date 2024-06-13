package tempest

import (
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
)

type ClassComposition interface {
	AlignmentClass
	AnimationClass
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
	SvgClass
	TransformClass
	TransitionClass
	TypoClass
}

type Tempest interface {
	ClassComposition
	If(condition bool, class Tempest) *Builder
	Extend(class Tempest) *Builder
	Name(name string) *Builder
	ClassList() []string
	Node() gox.Node
	String() string
}

type Builder struct {
	classes []string
	name    string
}

func createBuilder(classes ...string) *Builder {
	b := &Builder{
		classes: make([]string, 0),
	}
	if len(classes) > 0 {
		b.classes = append(b.classes, classes...)
	}
	return b
}

type style struct {
	prefix    string
	value     any
	unit      string
	fn        func(selector, value string) string
	modifiers []Modifier
	priority  int
}

func (b *Builder) If(condition bool, class Tempest) *Builder {
	if !condition {
		return b
	}
	b.Extend(class)
	return b
}

func (b *Builder) Extend(class Tempest) *Builder {
	b.classes = append(b.classes, class.ClassList()...)
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.name = name
	return b
}

func (b *Builder) ClassList() []string {
	return b.classes
}

func (b *Builder) Node() gox.Node {
	return gox.Class(b.String())
}

func (b *Builder) String() string {
	return strings.Join(b.classes, " ")
}

func (b *Builder) createStyle(s style) *Builder {
	var suffix string
	shouldHaveSuffix := strings.HasSuffix(s.prefix, "-")
	validatedValue := s.value
	if s.unit == Rem {
		validatedValue = convertSizeToRem(GlobalConfig.FontSize, s.value)
	}
	value := createValue(validatedValue, s.unit)
	if shouldHaveSuffix {
		suffix = createSuffix(s.value)
	}
	if strings.HasPrefix(value, "-") {
		s.prefix = "-" + s.prefix
	}
	class := applyClassModifiers(s.prefix+classEscaper.Replace(suffix), s.modifiers...)
	b.classes = append(b.classes, class)
	if classExists(class, s.priority) {
		return b
	}
	if shouldHaveSuffix {
		suffix = selectorEscaper.Replace(suffix)
	}
	selector := applySelectorModifiers(s.prefix+suffix, s.modifiers...)
	b.add(class, applyBreakpointModifiers(GlobalConfig.Breakpoint, s.fn(selector, value), s.modifiers...), s.priority)
	return b
}

func (b *Builder) add(k, v string, priority int) {
	if len(b.name) > 0 {
		b.addNamed(k, v, priority)
	}
	_, loaded := globalClassmaps.classmaps[priority].LoadOrStore(k, v)
	if !loaded {
		invalidateStylesCache()
	}
}

func (b *Builder) addNamed(k, v string, priority int) {
	if namedClassExists(b.name, k, priority) {
		return
	}
	_, loaded := namedClassmaps.classmaps[b.name][priority].LoadOrStore(k, v)
	if !loaded {
		invalidateNamedStylesCache(b.name)
	}
}
