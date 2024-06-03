package tempest

import "fmt"

func displayClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{display: %s;}`,
		selector,
		value,
	)
}
