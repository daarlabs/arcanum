package tempest

import "fmt"

func borderWidthClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{border-width: %s;}`,
		selector,
		value,
	)
}

func borderDirectionWidthClass(direction string, selector string, value string) string {
	return fmt.Sprintf(
		`%s{border-%s-width: %s;}`,
		selector,
		direction,
		value,
	)
}

func borderRadiusClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{border-radius: %s;}`,
		selector,
		transformKeywordWithMap(value, standardizedRadius),
	)
}
