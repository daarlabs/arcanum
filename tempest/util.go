package tempest

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	sizeRemCoeficient = float64(4)
)

var (
	classEscaper = strings.NewReplacer(
		` `, `_`,
		`.`, `\.`,
		`:`, `\:`,
		`[`, `\[`,
		`]`, `\]`,
		`/`, `\/`,
	)
)

func HexToRGB(hex string, opacity float64) string {
	rgb := convertHexToRGB(hex)
	return createRGBString(rgb, opacity/float64(100))
}

func escape(value string) string {
	return classEscaper.Replace(value)
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

func createRGBString(rgb RGB, opacity float64) string {
	return fmt.Sprintf("rgb(%d %d %d / %.2f)", rgb.R, rgb.G, rgb.B, opacity)
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
		return fmt.Sprintf("%.2f", v)
	case float32:
		if float64(v) == math.Floor(float64(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%.2f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
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
