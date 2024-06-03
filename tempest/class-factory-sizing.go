package tempest

import "fmt"

func widthClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{width: %s;}`,
		selector,
		transformKeywordToValue(Width, value),
	)
}

func minWidthClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{min-width: %s;}`,
		selector,
		transformKeywordToValue(Width, value),
	)
}

func maxWidthClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{max-width: %s;}`,
		selector,
		transformKeywordToValue(Width, value),
	)
}

func heightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{height: %s;}`,
		selector,
		transformKeywordToValue(Height, value),
	)
}

func minHeightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{min-height: %s;}`,
		selector,
		transformKeywordToValue(Height, value),
	)
}

func maxHeightClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{max-height: %s;}`,
		selector,
		transformKeywordToValue(Height, value),
	)
}

func sizeClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{width: %s; height: %s;}`,
		selector,
		transformKeywordToValue(Width, value),
		transformKeywordToValue(Height, value),
	)
}
