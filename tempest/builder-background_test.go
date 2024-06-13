package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestBackgroundBuilder(t *testing.T) {
	t.Run(
		"color", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Bg(Slate, 500).String()
			assert.Equal(
				t,
				colorClass("background-color", `.bg-slate-500`, Pallete[Slate][500], 1),
				c.Core.classes[class],
			)
		},
	)
	t.Run(
		"color opacity", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Bg(Slate, 500, Opacity(80)).String()
			assert.Equal(
				t,
				colorClass("background-color", `.bg-slate-500\/80`, Pallete[Slate][500], 0.8),
				c.Core.classes[class],
			)
		},
	)
}
