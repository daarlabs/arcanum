package tempest

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	sizeRemCoeficient = float64(4)
)

var (
	classEscaper = strings.NewReplacer(
		` `, `_`,
	)
)

var (
	selectorEscaper = strings.NewReplacer(
		` `, `_`,
		`.`, `\.`,
		`:`, `\:`,
		`[`, `\[`,
		`]`, `\]`,
		`/`, `\/`,
		`%`, `\%`,
	)
)

var (
	valueEscaper = strings.NewReplacer(
		`_`, ` `,
	)
)

var (
	cssExtractMatcher = regexp.MustCompile(`(?s)\.([\w-]+)\s*\{(.*?)}`)
)

func HexToRGB(hex string, opacity float64) string {
	rgb := convertHexToRGB(hex)
	return createRGBString(rgb, opacity)
}

func mergeConfigMap[T any](c1, c2 map[string]T) map[string]T {
	result := make(map[string]T)
	for k, v := range c1 {
		result[k] = v
	}
	for k, v := range c2 {
		result[k] = v
	}
	return result
}

func createStylesFromMap(m map[string]string) string {
	n := len(m)
	if n == 0 {
		return ""
	}
	result := make([]string, n)
	i := 0
	for k, v := range m {
		result[i] = fmt.Sprintf("%s: %s", k, v)
		i++
	}
	return strings.Join(result, ";") + ";"
}

func createRGBString(rgb RGB, opacity float64) string {
	if opacity > 1 {
		opacity = opacity / 100
	}
	return fmt.Sprintf("rgb(%d %d %d / %s)", rgb.R, rgb.G, rgb.B, createMostSuitableNumber(opacity))
}

func convertHexToRGB(hex string) RGB {
	var rgb RGB
	values, err := strconv.ParseUint(strings.TrimPrefix(hex, "#"), 16, 32)
	if err != nil {
		return RGB{}
	}
	rgb = RGB{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
	}
	return rgb
}

func transformKeywordToValue(name string, keyword string) string {
	if keyword == Full {
		return "100" + Pct
	}
	if keyword == Screen {
		if name == Width {
			return "100" + Vw
		}
		if name == Height {
			return "100" + Vh
		}
	}
	return keyword
}

func transformKeywordWithMap(keyword string, transforms map[string]string) string {
	if _, ok := transforms[keyword]; ok {
		return transforms[keyword]
	}
	return keyword
}

func stringifyMostSuitableNumericType(value any) string {
	switch v := value.(type) {
	case float64:
		if v == math.Floor(v) {
			return fmt.Sprintf("%d", int(v))
		}
		return createMostSuitableNumber(v)
	case float32:
		if float64(v) == math.Floor(float64(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return createMostSuitableNumber(float64(v))
	default:
		return fmt.Sprintf("%v", v)
	}
}

func createMostSuitableNumber(value float64) string {
	r := fmt.Sprintf("%.6f", value)
	r = strings.TrimRight(r, "0")
	r = strings.TrimRight(r, ".")
	return r
}

func convertSizeToRem(fontSize float64, value any) any {
	switch v := value.(type) {
	case int:
		r := float64(v) / (fontSize / sizeRemCoeficient)
		if r == math.Floor(r) {
			return int(r)
		}
		return r
	case float32:
		return float64(v) / (fontSize / sizeRemCoeficient)
	case float64:
		return v / (fontSize / sizeRemCoeficient)
	default:
		return v
	}
}

func createColorPrefix(category, name string, code int) string {
	if code == 0 {
		return fmt.Sprintf("%s-%s", category, name)
	}
	return fmt.Sprintf("%s-%s-%d", category, name, code)
}

func createOpacity(modifiers []Modifier) float64 {
	opacity := float64(1)
	for _, modifier := range modifiers {
		if modifier.Name != opacityModifier {
			continue
		}
		opacity = modifier.Value.(float64)
		break
	}
	if opacity > 1 {
		opacity = opacity / 100
	}
	return opacity
}

func extractCssProps(css string) []string {
	matches := cssExtractMatcher.FindAllStringSubmatch(css, -1)
	var properties []string
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}
		cleaned := strings.TrimSpace(match[2])
		properties = append(properties, cleaned)
	}
	return properties
}
