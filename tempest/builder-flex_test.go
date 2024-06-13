package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFlexBuilder(t *testing.T) {
	t.Run(
		"flex none", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().FlexNone().String()
			assert.Equal(
				t,
				flexClass(".flex-none", "none"),
				c.Core.classes[class],
			)
		},
	)
}
