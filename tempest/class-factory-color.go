package tempest

import "fmt"

func colorClass(name string, selector string, hex string, opacity float64) string {
	return fmt.Sprintf(
		`%s{%s: %s};`,
		selector,
		name,
		HexToRGB(hex, opacity),
	)
}
