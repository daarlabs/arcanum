package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestInteractivityBuilder(t *testing.T) {
	t.Run(
		"cursor", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Cursor("pointer").String()
			assert.Equal(
				t,
				cursorClass(".cursor-pointer", "pointer"),
				c.Core.classes[class],
			)
		},
	)
}
