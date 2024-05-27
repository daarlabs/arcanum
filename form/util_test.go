package form

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"is request form", func(t *testing.T) {
			req := testCreateFormRequest()
			assert.Equal(t, true, isRequestForm(req))
		},
	)
	t.Run(
		"is request multipart form", func(t *testing.T) {
			_, req, err := testCreateMultipartRequest()
			assert.Nil(t, err)
			assert.Equal(t, true, isRequestMultipartForm(req))
		},
	)
	t.Run(
		"is form submitted", func(t *testing.T) {
			reqGet := testGetRequest()
			reqPost := testCreateFormRequest()
			assert.Equal(t, false, isFormSubmitted(reqGet))
			assert.Equal(t, true, isFormSubmitted(reqPost))
		},
	)
	t.Run(
		"convert slice", func(t *testing.T) {
			src := []string{"1", "2", "3"}
			result := convertSlice[string, int](
				src, func(v string) int {
					return convertToInt(v)
				},
			)
			assert.Equal(t, []int{1, 2, 3}, result)
		},
	)
	t.Run(
		"convert to int", func(t *testing.T) {
			assert.Equal(t, 0, convertToInt("a"))
			assert.Equal(t, 1, convertToInt("1"))
		},
	)
	t.Run(
		"convert to float", func(t *testing.T) {
			assert.Equal(t, float64(0), convertToFloat("a"))
			assert.Equal(t, 999.99, convertToFloat("999.99"))
		},
	)
	t.Run(
		"get file suffix from name", func(t *testing.T) {
			assert.Equal(t, "png", getFileSuffixFromName("image.png"))
			assert.Equal(t, "pdf", getFileSuffixFromName("abc.pdf"))
			assert.Equal(t, "txt", getFileSuffixFromName("test.txt"))
		},
	)
}
