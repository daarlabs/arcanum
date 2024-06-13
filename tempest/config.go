package tempest

import (
	"fmt"
	"strings"
)

type Config struct {
	FontSize   float64
	FontFamily string
	Animation  map[string]Animation
	Color      map[string]Color
	Font       map[string]Font
	Shadow     map[string][]Shadow
	Container  map[string]string
	Breakpoint map[string]string
	Scripts    []string
	Styles     []string
	
	composition      map[string]Composition
	keyframes        string
	processedShadows map[string]Shadow
}

type Color map[int]string

type Font struct {
	Value string
	Url   string
}

func Compose(name string, composition Composition) {
	if GlobalConfig.composition == nil {
		GlobalConfig.composition = make(map[string]Composition)
	}
	GlobalConfig.composition[name] = composition
	props := make([]string, 0)
	for _, classname := range composition.Class.ClassList() {
		for _, sm := range globalClassmaps.classmaps {
			class, ok := sm.Load(classname)
			if !ok {
				continue
			}
			props = append(props, extractCssProps(class.(string))...)
		}
	}
	if len(props) > 0 {
		globalClassmaps.classmaps[len(globalClassmaps.classmaps)-1].Store(
			name, fmt.Sprintf(".%s{%s}", name, strings.Join(props, "")),
		)
	}
}

func (c *Config) processAnimations() {
	for name, a := range c.Animation {
		keyframes := make([]string, len(a.Keyframes))
		for i, k := range a.Keyframes {
			keyframes[i] = fmt.Sprintf("%s{%s}", k.Offset, createStylesFromMap(k.Styles))
		}
		c.keyframes += fmt.Sprintf(
			`@keyframes %s{%s}`,
			name,
			strings.Join(keyframes, ""),
		)
		nDelay := len(a.Delay)
		if nDelay == 0 {
			a.Value = fmt.Sprintf("%s %s %s %s", name, a.Duration, a.Timing, a.Repeat)
		}
		if nDelay > 0 {
			a.Value = fmt.Sprintf("%s %s %s %s %s", name, a.Duration, a.Timing, a.Delay, a.Repeat)
		}
		c.Animation[name] = a
	}
	c.keyframes += "\n"
}

func (c *Config) processShadows() {
	shadows := make(map[string]Shadow)
	for name, shadow := range c.Shadow {
		var color string
		parts := make([]string, len(shadow))
		for i, s := range shadow {
			if len(color) == 0 && len(s.Hex) == 0 {
				color = HexToRGB(DefaultShadowColor, s.Opacity)
			}
			if len(color) == 0 && len(s.Hex) > 0 {
				color = HexToRGB(s.Hex, s.Opacity)
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
}
