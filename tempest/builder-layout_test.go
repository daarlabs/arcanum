package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestLayoutBuilder(t *testing.T) {
	t.Run(
		"container", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Container().String()
			assert.Equal(
				t,
				containerClass(DefaultBreakpoints, DefaultContainer),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"top", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Top(4).String()
			assert.Equal(
				t,
				topClass(".top-4", "1rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"custom top", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Top("1px").String()
			assert.Equal(
				t,
				topClass(".top-1px", "1px"),
				c.Core.classes[class],
			)
		},
	)
}
