package util

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Run(
		"convert", func(t *testing.T) {
			s1 := []string{"1", "2", "3"}
			r1 := make([]int, 0)
			MustConvertSlice(s1, &r1)
			assert.Equal(t, []int{1, 2, 3}, r1)
		},
	)
}
