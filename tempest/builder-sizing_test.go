package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestSizingBuilder(t *testing.T) {
	t.Run(
		"basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().W(4).String()
			assert.Equal(
				t,
				widthClass(".w-4", "1rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"custom", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().W("1px").String()
			assert.Equal(
				t,
				widthClass(`.w-1px`, "1px"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"size", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Size(8).String()
			assert.Equal(
				t,
				sizeClass(`.size-8`, "2rem"),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"reserved", func(t *testing.T) {
			c := New(Config{}).Context()
			class1 := c.Class().W("auto").String()
			class2 := c.Class().W("full").String()
			class3 := c.Class().W(0).String()
			class4 := c.Class().W("screen").String()
			assert.Equal(
				t,
				widthClass(`.w-auto`, "auto"),
				c.Core.classes[class1],
			)
			assert.Equal(
				t,
				widthClass(`.w-full`, "full"),
				c.Core.classes[class2],
			)
			assert.Equal(
				t,
				widthClass(`.w-0`, "0"),
				c.Core.classes[class3],
			)
			assert.Equal(
				t,
				widthClass(`.w-screen`, "screen"),
				c.Core.classes[class4],
			)
		},
	)
}
