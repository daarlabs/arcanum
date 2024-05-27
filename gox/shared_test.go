package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharedNodes(t *testing.T) {
	t.Run(
		"render shared element", func(t *testing.T) {
			styleEl := Style(Element())
			assert.Equal(t, "<style></style>", Render(styleEl))
		},
	)
	t.Run(
		"render shared attribute", func(t *testing.T) {
			styleValue := `background:blue;`
			styleAttr := Style(Text(styleValue))
			assert.Equal(t, fmt.Sprintf(`style="%s"`, styleValue), Render(styleAttr))
		},
	)
}
