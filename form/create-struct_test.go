package form

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestCreateStruct(t *testing.T) {
	t.Run(
		"convert", func(t *testing.T) {
			form, err := Build[testForm](
				New(
					Add("name").With(Text(testNameValue)),
					Add("quantity").With(Number[int](testQuantityValue)),
					Add("amount").With(Number[float64](testAmountValue)),
					Add("checked").With(Checkbox(testCheckedValue)),
				),
			)
			assert.Nil(t, err)
			result := CreateStruct[testForm, testModel](&form)
			assert.Equal(
				t,
				testModel{
					Name:     testNameValue,
					Quantity: testQuantityValue,
					Amount:   testAmountValue,
					Checked:  testCheckedValue,
				},
				result,
			)
		},
	)
}
