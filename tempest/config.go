package tempest

import (
	"fmt"
	"strings"
)

type Config struct {
	FontSize         float64
	FontFamily       string
	Color            map[string]Color
	Font             map[string]Font
	Shadow           map[string][]Shadow
	Container        map[string]string
	Breakpoint       map[string]string
	Scripts          []string
	Styles           []string
	processedShadows map[string]Shadow
}

type Color map[int]string

type Font struct {
	Value string
	Url   string
}

func (c Config) processShadows() Config {
	shadows := make(map[string]Shadow)
	for name, shadow := range c.Shadow {
		var color string
		parts := make([]string, len(shadow))
		for i, s := range shadow {
			if len(color) == 0 && len(s.Hex) == 0 {
				color = HexToRGB(DefaultShadowColor, s.Opacity)
			}
			if len(color) == 0 && len(s.Hex) > 0 {
				color = HexToRGB(s.Color, s.Opacity)
			}
			if len(color) == 0 && len(s.Color) > 0 {
				color = s.Color
			}
			parts[i] = fmt.Sprintf("%s var(%s)", s.Value, shadowColorVar)
		}
		shadows[name] = Shadow{
			Value: strings.Join(parts, ","),
			Color: color,
		}
	}
	c.processedShadows = shadows
	return c
}
