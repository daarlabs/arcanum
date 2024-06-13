package tempest

import "fmt"

func fillClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{fill: %s;}`,
		selector,
		value,
	)
}

func strokeClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{stroke: %s;}`,
		selector,
		value,
	)
}
