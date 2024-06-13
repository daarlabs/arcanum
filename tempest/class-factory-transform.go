package tempest

import (
	"fmt"
	"strings"
)

var (
	transformOriginEscaper = strings.NewReplacer(
		`-`, ` `,
	)
)

func transformClass(selector string, _ string) string {
	return fmt.Sprintf(
		`%s{transform: translate(var(%s),var(%s)) rotate(var(%s)) skewX(var(%s)) skewY(var(%s)) scaleX(var(%s)) scaleY(var(%s));}`,
		selector,
		transformTranslateXVar,
		transformTranslateYVar,
		transformRotateVar,
		transformSkewXVar,
		transformSkewYVar,
		transformScaleXVar,
		transformScaleYVar,
	)
}

func transformRotateClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformRotateVar,
		value,
		transformClass(selector, value),
	)
}

func transformTranslateXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformTranslateXVar,
		value,
		transformClass(selector, value),
	)
}

func transformTranslateYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformTranslateYVar,
		value,
		transformClass(selector, value),
	)
}

func transformSkewXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformSkewXVar,
		value,
		transformClass(selector, value),
	)
}

func transformSkewYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformSkewYVar,
		value,
		transformClass(selector, value),
	)
}

func transformScaleClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;%s: %s;}%s`,
		selector,
		transformScaleXVar,
		value,
		transformSkewYVar,
		value,
		transformClass(selector, value),
	)
}

func transformScaleXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformScaleXVar,
		value,
		transformClass(selector, value),
	)
}

func transformScaleYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{%s: %s;}%s`,
		selector,
		transformScaleYVar,
		value,
		transformClass(selector, value),
	)
}

func transformOriginClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{transform-origin: %s;}`,
		selector,
		transformOriginEscaper.Replace(value),
	)
}
