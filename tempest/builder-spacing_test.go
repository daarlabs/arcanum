package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestSpacingBuilder(t *testing.T) {
	t.Run(
		"basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P(4).String()
			assert.Equal(
				t,
				paddingClass(".p-4", "1rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"negative", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P(-4).String()
			assert.Equal(
				t,
				paddingClass(".-p-4", "-1rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"custom", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P("1px").String()
			assert.Equal(
				t,
				paddingClass(`.p-1px`, "1px"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"dark modifier", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P(4, Dark()).String()
			assert.Equal(
				t,
				paddingClass(applySelectorModifiers("p-4", Dark()), "1rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"breakpoint modifier", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P(4, Xs()).String()
			assert.Equal(
				t,
				applyBreakpointModifiers(
					DefaultBreakpoints,
					applySelectorModifiers(paddingClass("p-4", "1rem"), Xs()),
					Xs(),
				),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"multiple modifiers", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().P(4, Dark(), Lg()).String()
			assert.Equal(
				t,
				applyBreakpointModifiers(
					DefaultBreakpoints,
					paddingClass(applySelectorModifiers("p-4", Dark(), Lg()), "1rem"), Lg(),
				),
				c.Core.classes[class],
			)
		},
	)
}
