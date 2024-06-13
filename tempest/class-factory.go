package tempest

import (
	"fmt"
	"strings"
)

func createSuffix(value any) string {
	var r string
	switch v := value.(type) {
	case int:
		r = fmt.Sprintf("%d", v)
	case float32:
		r = createMostSuitableNumber(float64(v))
	case float64:
		r = createMostSuitableNumber(v)
	case string:
		r = v
	default:
		r = fmt.Sprintf("%v", v)
	}
	r = strings.TrimPrefix(r, "-")
	return r
}

func createValue(value any, unit string) string {
	var r string
	switch v := value.(type) {
	case int:
		if v == 0 {
			r = fmt.Sprintf("%d", v)
		}
		if v != 0 {
			r = fmt.Sprintf("%d%s", v, unit)
		}
	case float32:
		if v == 0 {
			r = createMostSuitableNumber(float64(v))
		}
		if v != 0 {
			r = createMostSuitableNumber(float64(v)) + unit
		}
	case float64:
		if v == 0 {
			r = createMostSuitableNumber(v)
		}
		if v != 0 {
			r = createMostSuitableNumber(v) + unit
		}
	case string:
		r = valueEscaper.Replace(v)
	default:
		r = fmt.Sprintf("%v", v)
	}
	return r
}

// Breakpoint
func createBreakpoint(breakpoints map[string]string, key string, class string) string {
	if _, ok := breakpoints[key]; !ok {
		return class
	}
	return fmt.Sprintf(
		`@media (min-width: %s) {%s}`,
		breakpoints[key],
		class,
	)
}

// Container
func createContainer(value string) string {
	if value == "100"+Pct {
		return fmt.Sprintf(
			`.container{width:%s;}`,
			value,
		)
	}
	return fmt.Sprintf(
		`.container{max-width:%s;}`,
		value,
	)
}

func createContainerCenter() string {
	return `.container{margin:0 auto;}`
}
