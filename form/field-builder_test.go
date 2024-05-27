package form

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFieldBuilder(t *testing.T) {
	t.Run(
		"create text", func(t *testing.T) {
			f := Add("name").With(Text(testNameValue))
			assert.Equal(t, "name", f.name)
			assert.Equal(t, testNameValue, testNameValue)
			assert.Equal(t, fieldTypeText, f.fieldType)
			assert.Equal(t, fieldDataTypeString, f.dataType)
		},
	)
	t.Run(
		"create hidden", func(t *testing.T) {
			token := "12345"
			f := Add(CsrfToken).With(Hidden(token))
			assert.Equal(t, CsrfToken, f.name)
			assert.Equal(t, token, f.value)
			assert.Equal(t, fieldTypeHidden, f.fieldType)
			assert.Equal(t, fieldDataTypeString, f.dataType)
		},
	)
	t.Run(
		"with validate ok", func(t *testing.T) {
			f := Add("name").With(Text("abc"), Validate.Required())
			f.valid = len(validateField(f, nil)) == 0
			assert.Equal(t, true, f.valid)
			assert.Equal(t, fieldTypeText, f.fieldType)
			assert.Equal(t, fieldDataTypeString, f.dataType)
		},
	)
	t.Run(
		"with validate fail", func(t *testing.T) {
			f := Add("name").With(Text("abc"), Validate.Min(4))
			f.valid = len(validateField(f, nil)) == 0
			assert.Equal(t, false, f.valid)
			assert.Equal(t, fieldTypeText, f.fieldType)
			assert.Equal(t, fieldDataTypeString, f.dataType)
		},
	)
	t.Run(
		"with number slice", func(t *testing.T) {
			f := Add("name").With(Number[int](1, 2))
			assert.Equal(t, []int{1, 2}, f.value)
		},
	)
}
