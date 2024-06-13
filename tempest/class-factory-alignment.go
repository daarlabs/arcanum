package tempest

import "fmt"

func justifyClass(selector, name, value string) string {
	return fmt.Sprintf(
		`%s{justify-%s: %s;}`,
		selector,
		name,
		transformAlignmentValue(value),
	)
}

func alignClass(selector, name, value string) string {
	return fmt.Sprintf(
		`%s{align-%s: %s;}`,
		selector,
		name,
		transformAlignmentValue(value),
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

func transformAlignmentValue(value string) string {
	if value == "start" {
		return "flex-start"
	}
	if value == "end" {
		return "flex-end"
	}
	if value == "between" {
		return "space-between"
	}
	if value == "around" {
		return "space-around"
	}
	if value == "evenly" {
		return "space-evenly"
	}
	return value
}
