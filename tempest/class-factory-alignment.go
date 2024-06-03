package tempest

import "fmt"

func justifyClass(selector, name, value string) string {
	return fmt.Sprintf(
		`%s{justify-%s: %s;}`,
		selector,
		name,
		value,
	)
}

func alignClass(selector, name, value string) string {
	return fmt.Sprintf(
		`%s{align-%s: %s;}`,
		selector,
		name,
		value,
	)
}

func placeClass(selector, name, value string) string {
	return fmt.Sprintf(
		`%s{place-%s: %s;}`,
		selector,
		name,
		value,
	)
}
