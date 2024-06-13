package tempest

import "fmt"

func cursorClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{cursor: %s;}`,
		selector,
		value,
	)
}

func pointerEventsClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{pointer-events: %s;}`,
		selector,
		value,
	)
}

func userSelectClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{user-select: %s;}`,
		selector,
		value,
	)
}
