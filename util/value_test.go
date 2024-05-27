package util

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Run(
		"convert", func(t *testing.T) {
			s1 := "1"
			r1 := new(int)
			MustConvertValue(s1, r1)
			assert.Equal(t, 1, *r1)
		},
	)
}
