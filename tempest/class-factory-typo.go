package tempest

import "fmt"

func fontFamilyClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{font-family: %s;}`,
		selector,
		value,
	)
}

func fontSizeClass(selector string, value string) string {
	if lh, ok := standardizedLineHeight[value]; ok {
		return fmt.Sprintf(
			`%s{font-size: %s; line-height: %s;}`,
			selector,
			transformKeywordWithMap(value, standardizedSize),
			lh,
		)
	}
	return fmt.Sprintf(
		`%s{font-size: %s;}`,
		selector,
		transformKeywordWithMap(value, standardizedSize),
	)
}

func fontWeightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{font-weight: %s;}`,
		selector,
		value,
	)
}

func textAlignClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{text-align: %s;}`,
		selector,
		value,
	)
}

func textDecorationClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{text-decoration: %s;}`,
		selector,
		value,
	)
}

func truncateClass(selector string, _ string) string {
	return fmt.Sprintf(
		`%s{overflow: hidden;text-overflow: ellipsis;white-space: nowrap;}`,
		selector,
	)
}
