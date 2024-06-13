package tempest

import (
	"fmt"
)

// Padding
func paddingClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding: %s;}`,
		selector,
		value,
	)
}

func paddingXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-left: %[2]s;padding-right: %[2]s;}`,
		selector,
		value,
	)
}

func paddingYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-top: %[2]s;padding-bottom: %[2]s;}`,
		selector,
		value,
	)
}

func paddingTopClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-top: %s;}`,
		selector,
		value,
	)
}

func paddingRightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-right: %s;}`,
		selector,
		value,
	)
}

func paddingBottomClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-bottom: %s;}`,
		selector,
		value,
	)
}

func paddingLeftClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{padding-left: %s;}`,
		selector,
		value,
	)
}

// Margin
func marginClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin: %s;}`,
		selector,
		value,
	)
}

func marginXAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-left: %[2]s;margin-right: %[2]s;}`,
		selector,
		value,
	)
}

func marginYAxisClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-top: %[2]s;margin-bottom: %[2]s;}`,
		selector,
		value,
	)
}

func marginTopClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-top: %s;}`,
		selector,
		value,
	)
}

func marginRightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-right: %s;}`,
		selector,
		value,
	)
}

func marginBottomClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-bottom: %s;}`,
		selector,
		value,
	)
}

func marginLeftClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{margin-left: %s;}`,
		selector,
		value,
	)
}
