package tempest

import "fmt"

func colorClass(name string, selector string, hex string, opacity float64) string {
	var color string
	if hex != Transparent {
		color = HexToRGB(hex, opacity)
	}
	if hex == Transparent {
		color = Transparent
	}
	return fmt.Sprintf(
		`%s{%s: %s;}`,
		selector,
		name,
		color,
	)
}
