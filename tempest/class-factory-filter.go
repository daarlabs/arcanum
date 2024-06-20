package tempest

import "fmt"

func filterClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{filter: %s;}`,
		selector,
		value,
	)
}
