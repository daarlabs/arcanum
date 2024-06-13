package tempest

import (
	"fmt"
)

func shadowClass(shadows map[string]Shadow, selector string, value string) string {
	shadow, ok := shadows[value]
	if !ok {
		shadow = Shadow{}
	}
	return fmt.Sprintf(
		`%[1]s{%[2]s: %[3]s;%[4]s: %[5]s;box-shadow: var(%[2]s);}`,
		selector,
		shadowVar,
		shadow.Value,
		shadowColorVar,
		shadow.Color,
	)
}

func shadowColorClass(selector string, hex string, opacity float64) string {
	return fmt.Sprintf(
		`%s{%s: %s;}`,
		selector,
		shadowColorVar,
		createRGBString(convertHexToRGB(hex), opacity),
	)
}

func opacityClass(selector string, opacity float64) string {
	return fmt.Sprintf(
		`%s{opacity: %s;}`,
		selector,
		createMostSuitableNumber(opacity),
	)
}
