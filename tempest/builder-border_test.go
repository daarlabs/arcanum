package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestBorderBuilder(t *testing.T) {
	t.Run(
		"width basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Border(4).String()
			assert.Equal(
				t,
				borderWidthClass(`.border-4`, "4px"),
				c.Tempest.classes[class],
			)
		},
	)
	t.Run(
		"rounded basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Rounded(SizeBase).String()
			assert.Equal(
				t,
				borderRadiusClass(`.rounded-base`, SizeBase),
				c.Tempest.classes[class],
			)
		},
	)
	t.Run(
		"rounded custom", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Rounded("4px").String()
			assert.Equal(
				t,
				borderRadiusClass(`.rounded-4px`, "4px"),
				c.Tempest.classes[class],
			)
		},
	)
}
