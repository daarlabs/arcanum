package tempest

import "fmt"

func cursorClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{cursor: %s;}`,
		selector,
		value,
	)
}
