package form

import (
	"mime/multipart"
	"net/url"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestProcessor(t *testing.T) {
	t.Run(
		"process request", func(t *testing.T) {
			req := testCreateFormRequest()
			formData, _, err := processRequest(req, defaultBodyLimit)
			assert.Nil(t, err)
			assert.Equal(t, testNameValue, formData.Get("name"))
			assert.Equal(t, testQuantityValue, convertToInt(formData.Get("quantity")))
			assert.Equal(t, testAmountValue, convertToFloat(formData.Get("amount")))
			assert.Equal(t, testCheckedValue, formData.Get("checked") == "on")
		},
	)
	t.Run(
		"process form data", func(t *testing.T) {
			req := testCreateFormRequest()
			form, err := Build[testForm](
				New(
					Add("roles").Multiple().With(Text()),
					Add("name").With(Text()),
					Add("quantity").With(Number[int]()),
					Add("amount").With(Number[float64]()),
					Add("checked").With(Checkbox()),
				).Request(req),
			)
			assert.Nil(t, err)
			assert.Equal(t, []string{testRole}, form.Roles.Value)
			assert.Equal(t, testNameValue, form.Name.Value)
			assert.Equal(t, testQuantityValue, form.Quantity.Value)
			assert.Equal(t, testAmountValue, form.Amount.Value)
			assert.Equal(t, testCheckedValue, form.Checked.Value)
		},
	)
	t.Run(
		"process form files", func(t *testing.T) {
			fileBytes, req, err := testCreateMultipartRequest()
			assert.Nil(t, err)
			form, err := Build[testForm](
				New(Add("test").With(File())).
					Request(req),
			)
			assert.Nil(t, err)
			assert.Equal(t, fileBytes, form.Test.Value.Data)
		},
	)
	t.Run(
		"parse form", func(t *testing.T) {
			req := testCreateFormRequest()
			rt, err := parseForm(req, defaultBodyLimit)
			assert.Nil(t, err)
			assert.Equal(t, requestTypeForm, rt)
		},
	)
	t.Run(
		"parse multipart form", func(t *testing.T) {
			_, req, err := testCreateMultipartRequest()
			assert.Nil(t, err)
			rt, err := parseForm(req, defaultBodyLimit)
			assert.Nil(t, err)
			assert.Equal(t, requestTypeMultipartForm, rt)
		},
	)
	t.Run(
		"create empty process result", func(t *testing.T) {
			formData, files, err := createEmptyProcessRequestResult(nil)
			assert.Nil(t, err)
			assert.Equal(t, make(url.Values), formData)
			assert.Equal(t, make(map[string][]*multipart.FileHeader), files)
		},
	)
}
