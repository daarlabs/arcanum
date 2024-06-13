package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestGridBuilder(t *testing.T) {
	t.Run(
		"cols", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().GridCols(3).String()
			assert.Equal(
				t,
				gridTemplateColumnsClass(".grid-cols-3", "3", false),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"custom cols", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().GridCols("1fr 1fr 1fr").String()
			assert.Equal(
				t,
				gridTemplateColumnsClass(".grid-cols-1fr_1fr_1fr", "1fr 1fr 1fr", true),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"order", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Order(1).String()
			assert.Equal(
				t,
				orderClass(".order-1", "1"),
				c.Core.classes[class],
			)
		},
	)
}
