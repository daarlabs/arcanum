package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestModifier(t *testing.T) {
	t.Run(
		"basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Bg(Slate, 500, Hover()).String()
			assert.Equal(
				t,
				colorClass("background-color", applySelectorModifiers(`bg-slate-500`, Hover()), Pallete[Slate][500], 1),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"multiple", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Bg(Slate, 500, Dark(), Hover()).String()
			assert.Equal(
				t,
				colorClass("background-color", applySelectorModifiers(`bg-slate-500`, Dark(), Hover()), Pallete[Slate][500], 1),
				c.Core.classes[class],
			)
		},
	)
}
