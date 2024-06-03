package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTransitionBuilder(t *testing.T) {
	t.Run(
		"basic", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Transition().String()
			assert.Equal(
				t,
				transitionClass(".transition", ""),
				c.Tempest.classes[class],
			)
		},
	)
	t.Run(
		"with modifiers", func(t *testing.T) {
			c := New(Config{}).Context()
			class := c.Class().Transition(Dark(), Hover()).String()
			assert.Equal(
				t,
				transitionClass(applySelectorModifiers("transition", Dark(), Hover()), ""),
				c.Tempest.classes[class],
			)
		},
	)
}
