package tempest

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"convert size to rem", func(t *testing.T) {
			assert.Equal(t, 0.25, convertSizeToRem(16, 1))
			assert.Equal(t, 1, convertSizeToRem(16, 4))
			assert.Equal(t, 4, convertSizeToRem(16, 16))
		},
	)
}
