package tempest

import "fmt"

func gridTemplateColumnsClass(selector string, value string, custom bool) string {
	if custom {
		return fmt.Sprintf(
			`%s{grid-template-columns: %s;}`,
			selector,
			value,
		)
	}
	return fmt.Sprintf(
		`%s{grid-template-columns: repeat(%s, minmax(0, 1fr));}`,
		selector,
		value,
	)
}

func gridTemplateRowsClass(selector string, value string, custom bool) string {
	if custom {
		return fmt.Sprintf(
			`%s{grid-template-rows: %s;}`,
			selector,
			value,
		)
	}
	return fmt.Sprintf(
		`%s{grid-template-rows: repeat(%s, minmax(0, 1fr));}`,
		selector,
		value,
	)
}

func gapClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{gap: %s;}`,
		selector,
		value,
	)
}

func orderClass(selector string, value string) string {
	return fmt.Sprintf(
		`%s{order: %s;}`,
		selector,
		value,
	)
}
