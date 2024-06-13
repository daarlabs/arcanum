package tempest

import (
	"fmt"
	"strings"
)

type Modifier struct {
	Name  string
	Value any
}

type modifiers struct {
	dark         bool
	placeholder  bool
	opacity      float64
	useOpacity   bool
	hover        bool
	checked      bool
	focus        bool
	breakpoint   string
	peerChecked  bool
	peerFocus    bool
	peerHover    bool
	groupChecked bool
	groupFocus   bool
	groupHover   bool
}

const (
	Group = "group"
	Peer  = "peer"
)

const (
	darkModifier        = "dark"
	placeholderModifier = "placeholder"
	opacityModifier     = "opacity"
	hoverModifier       = "hover"
	checkedModifier     = "checked"
	focusModifier       = "focus"
	breakpointModifier  = "breakpoint"
)

func Dark() Modifier {
	return Modifier{Name: darkModifier}
}

func Placeholder() Modifier {
	return Modifier{Name: placeholderModifier}
}

func Opacity(opacity float64) Modifier {
	return Modifier{Name: opacityModifier, Value: opacity}
}

func Hover(category ...string) Modifier {
	var c string
	if len(category) > 0 {
		c = category[0]
	}
	return Modifier{Name: hoverModifier, Value: c}
}

func Checked(category ...string) Modifier {
	var c string
	if len(category) > 0 {
		c = category[0]
	}
	return Modifier{Name: checkedModifier, Value: c}
}
func Focus(category ...string) Modifier {
	var c string
	if len(category) > 0 {
		c = category[0]
	}
	return Modifier{Name: focusModifier, Value: c}
}

func Xs() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeXs}
}

func Sm() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeSm}
}

func Md() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeMd}
}

func Lg() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeLg}
}

func Xl() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeXl}
}

func Xxl() Modifier {
	return Modifier{Name: breakpointModifier, Value: SizeXxl}
}

func processModifiers(items ...Modifier) modifiers {
	var r modifiers
	for _, item := range items {
		switch item.Name {
		case darkModifier:
			r.dark = true
		case hoverModifier:
			if v, ok := item.Value.(string); ok {
				if v == Group {
					r.groupHover = true
					continue
				}
				if v == Peer {
					r.peerHover = true
					continue
				}
			}
			r.hover = true
		case checkedModifier:
			if v, ok := item.Value.(string); ok {
				if v == Group {
					r.groupChecked = true
					continue
				}
				if v == Peer {
					r.peerChecked = true
					continue
				}
			}
			r.checked = true
		case focusModifier:
			if v, ok := item.Value.(string); ok {
				if v == Group {
					r.groupFocus = true
					continue
				}
				if v == Peer {
					r.peerFocus = true
					continue
				}
			}
			r.focus = true
		case placeholderModifier:
			r.placeholder = true
		case opacityModifier:
			r.opacity = item.Value.(float64)
			r.useOpacity = true
		case breakpointModifier:
			r.breakpoint = item.Value.(string)
		}
	}
	return r
}

func createModifiersParts(m modifiers) []string {
	parts := make([]string, 0)
	if m.dark {
		parts = append(parts, "dark")
	}
	if m.hover {
		parts = append(parts, "hover")
	}
	if m.checked {
		parts = append(parts, "checked")
	}
	if m.focus {
		parts = append(parts, "focus")
	}
	if m.placeholder {
		parts = append(parts, "placeholder")
	}
	if m.peerChecked {
		parts = append(parts, "peer-checked")
	}
	if m.peerFocus {
		parts = append(parts, "peer-focus")
	}
	if m.peerHover {
		parts = append(parts, "peer-hover")
	}
	if m.groupChecked {
		parts = append(parts, "group-checked")
	}
	if m.groupFocus {
		parts = append(parts, "group-focus")
	}
	if m.groupHover {
		parts = append(parts, "group-hover")
	}
	if len(m.breakpoint) > 0 {
		parts = append(parts, m.breakpoint)
	}
	return parts
}

func applyClassModifiers(class string, modifiers ...Modifier) string {
	m := processModifiers(modifiers...)
	parts := createModifiersParts(m)
	parts = append(parts, class)
	r := strings.Join(parts, ":")
	if m.useOpacity {
		r += fmt.Sprintf("/%s", createOpacitySuffix(m.opacity))
	}
	return r
}

func createOpacitySuffix(opacity float64) string {
	if opacity < 1 {
		opacity = opacity * 100
	}
	return createMostSuitableNumber(opacity)
}

func applySelectorModifiers(selector string, modifiers ...Modifier) string {
	var result string
	m := processModifiers(modifiers...)
	if m.peerHover {
		result += ".peer:hover ~ "
	}
	if m.peerChecked {
		result += ".peer:checked ~ "
	}
	if m.peerFocus {
		result += ".peer:focus ~ "
	}
	if m.groupHover {
		result += ".group:hover > "
	}
	if m.groupChecked {
		result += ".group:checked > "
	}
	if m.groupFocus {
		result += ".group:focus > "
	}
	parts := createModifiersParts(m)
	parts = append(parts, selector)
	result += "." + strings.Join(parts, `\:`)
	if m.useOpacity {
		result += fmt.Sprintf(`\/%s`, createOpacitySuffix(m.opacity))
	}
	if m.hover {
		result += ":hover"
	}
	if m.checked {
		result += ":checked"
	}
	if m.focus {
		result += ":focus"
	}
	if m.dark {
		result += ":is(.dark *)"
	}
	if m.placeholder {
		result += "::placeholder"
	}
	return result
}
func applyBreakpointModifiers(breakpoints map[string]string, class string, modifiers ...Modifier) string {
	var breakpoint string
	for _, m := range modifiers {
		switch m.Name {
		case breakpointModifier:
			switch v := m.Value.(type) {
			case string:
				breakpoint = v
			}
		}
	}
	if len(breakpoint) == 0 {
		return class
	}
	return createBreakpoint(breakpoints, breakpoint, class)
}
