package tempest

import (
	"fmt"
)

func animationClass(selector, value string) string {
	return fmt.Sprintf(
		`%s{animation: %s;}`,
		selector,
		value,
	)
}
