package tempest

import "fmt"

func flexClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{flex: %s;}`,
		selector,
		value,
	)
}

func flexDirectionClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{flex-direction: %s;}`,
		selector,
		value,
	)
}

func flexWrapClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{flex-wrap: %s;}`,
		selector,
		value,
	)
}

func flexGrowClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{flex-grow: %s;}`,
		selector,
		value,
	)
}

func flexShrinkClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{flex-shrink: %s;}`,
		selector,
		value,
	)
}
