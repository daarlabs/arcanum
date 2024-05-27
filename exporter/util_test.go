package exporter

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"get letter with index", func(t *testing.T) {
			assert.Equal(t, "A", getLetterWithIndex(0))
			assert.Equal(t, "Z", getLetterWithIndex(25))
			assert.Equal(t, "BA", getLetterWithIndex(51))
			assert.Equal(t, "AA", getLetterWithIndex(26))
			assert.Equal(t, "ZZ", getLetterWithIndex(675))
		},
	)
}
