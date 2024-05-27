package form

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	t.Run(
		"required", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text(), Validate.Required()),
					Add("quantity").With(Number[int](), Validate.Required()),
					Add("amount").With(Number[float64](), Validate.Required()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, defaultRequiredMessage, form.Email.Messages[0])
			assert.Equal(t, defaultRequiredMessage, form.Quantity.Messages[0])
			assert.Equal(t, defaultRequiredMessage, form.Amount.Messages[0])
		},
	)
	t.Run(
		"email valid", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Email("test@test.cz"), Validate.Email()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, 0, len(form.Email.Messages))
		},
	)
	t.Run(
		"email invalid", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Email()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, defaultEmailMessage, form.Email.Messages[0])
		},
	)
	t.Run(
		"string min", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Min(5)),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, defaultMinTextMessage, form.Email.Messages[0])
		},
	)
	t.Run(
		"string max", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("email").With(Text("test"), Validate.Max(3)),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, defaultMaxTextMessage, form.Email.Messages[0])
		},
	)
	t.Run(
		"string slice", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("roles").Multiple().With(Text(), Validate.Required()),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, 1, len(form.Roles.Messages))
			assert.Equal(t, defaultRequiredMessage, form.Roles.Messages[0])
		},
	)
}
