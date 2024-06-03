package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTypeBuilder(t *testing.T) {
	t.Run(
		"standardized", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().FontSize(SizeLg).String()
			assert.Equal(
				t,
				fontSizeClass(".font-lg", SizeLg),
				c.Tempest.classes[class],
			)
		},
	)
	t.Run(
		"custom", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().FontSize("14px").String()
			assert.Equal(
				t,
				fontSizeClass(`.font-14px`, "14px"),
				c.Tempest.classes[class],
			)
		},
	)
	t.Run(
		"bold", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().FontBold().String()
			assert.Equal(
				t,
				fontWeightClass(`.font-bold`, "700"),
				c.Tempest.classes[class],
			)
		},
	)
}
