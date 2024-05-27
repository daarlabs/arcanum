package form

import (
	"net/http"
	"reflect"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFormBuild(t *testing.T) {
	t.Run(
		"build", func(t *testing.T) {
			req := testCreateFormRequest()
			f, err := Build[testForm](
				New(
					Add("name").With(Text(), Validate.Required()),
					Add("quantity").With(Number[int](), Validate.Required()),
					Add("amount").With(Number[float64](), Validate.Required()),
					Add("checked").With(Checkbox(), Validate.Required()),
				).Request(req),
			)
			assert.Nil(t, err)
			assert.Equal(t, true, f.Valid)
			assert.Equal(t, true, f.Submitted)
			assert.Equal(t, testNameValue, f.Name.Value)
			assert.Equal(t, testQuantityValue, f.Quantity.Value)
			assert.Equal(t, testAmountValue, f.Amount.Value)
			assert.Equal(t, testCheckedValue, f.Checked.Value)
		},
	)
	t.Run(
		"build base form", func(t *testing.T) {
			form := &testForm{}
			formRef := reflect.ValueOf(form)
			req := testCreateFormRequest()
			req.Header.Set(contentType, contentTypeForm)
			fb := New().
				Method(http.MethodPost).
				Action(testAction).
				Name(testName).
				Limit(defaultBodyLimit).
				Request(req)
			fb.submitted = isFormSubmitted(req)
			buildBaseForm(formRef, fb)
			assert.Equal(t, http.MethodPost, form.Method)
			assert.Equal(t, testAction, form.Action)
			assert.Equal(t, true, form.Valid)
			assert.Equal(t, true, form.Submitted)
		},
	)
	t.Run(
		"create base form", func(t *testing.T) {
			req := testCreateFormRequest()
			fb := New().
				Method(http.MethodPost).
				Action(testAction).
				Name(testName).
				Limit(defaultBodyLimit).
				Request(req)
			fb.submitted = isFormSubmitted(req)
			baseForm := createBaseForm(fb)
			assert.Equal(t, true, baseForm.Submitted)
			assert.Equal(t, http.MethodPost, baseForm.Method)
			assert.Equal(t, testAction, baseForm.Action)
		},
	)
}

func TestFormFieldBuild(t *testing.T) {
	t.Run(
		"build form field", func(t *testing.T) {
			form := &testForm{}
			formRef := reflect.ValueOf(form)
			name := Add("name").With(Text(testNameValue))
			quantity := Add("quantity").With(Number[int](testQuantityValue))
			amount := Add("amount").With(Number[float64](testAmountValue))
			checked := Add("checked").With(Checkbox(testCheckedValue))
			buildFormField(formRef, name, nil)
			buildFormField(formRef, quantity, nil)
			buildFormField(formRef, amount, nil)
			buildFormField(formRef, checked, nil)
			assert.Equal(t, name.value, form.Name.Value)
			assert.Equal(t, quantity.value, form.Quantity.Value)
			assert.Equal(t, amount.value, form.Amount.Value)
			assert.Equal(t, checked.value, form.Checked.Value)
		},
	)
	t.Run(
		"create form field", func(t *testing.T) {
			fb := Add("name").With(Text(testNameValue))
			f := createFormField[string](fb, nil)
			f.Messages = validateField(fb, nil)
			fb.valid = len(f.Messages) == 0
			assert.Equal(t, f.Type, fb.fieldType)
			assert.Equal(t, f.Value, fb.value)
			assert.Equal(t, f.Id, fb.id)
			assert.Equal(t, f.Name, fb.name)
			assert.Equal(t, f.Multiple, fb.multiple)
			assert.Equal(t, true, fb.valid)
			assert.Equal(t, 0, len(f.Messages))
		},
	)
}
