package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTransformBuilder(t *testing.T) {
	t.Run(
		"rotate", func(t *testing.T) {
			c := New(Config{}).Context()
			class1 := c.Class().Rotate(45).String()
			class2 := c.Class().Rotate(-180).String()
			assert.Equal(
				t,
				transformRotateClass(".rotate-45", "45deg"),
				c.Core.classes[class1],
			)
			assert.Equal(
				t,
				transformRotateClass(".-rotate-180", "-180deg"),
				c.Core.classes[class2],
			)
		},
	)
	t.Run(
		"multiple", func(t *testing.T) {
			c := New(Config{}).Context()
			class1 := c.Class().Rotate(45).String()
			class2 := c.Class().TranslateX(4).String()
			assert.Equal(
				t,
				transformRotateClass(".rotate-45", "45deg"),
				c.Core.classes[class1],
			)
			assert.Equal(
				t,
				transformTranslateXAxisClass(".translate-x-4", "1rem"),
				c.Core.classes[class2],
			)
		},
	)
	t.Run(
		"skew x", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().SkewX(45).String()
			assert.Equal(
				t,
				transformSkewXAxisClass(".skew-x-45", "45deg"),
				c.Core.classes[class],
			)
		},
	)
}
