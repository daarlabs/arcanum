package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	type testStruct struct {
		Title    string  `json:"title"`
		Active   bool    `json:"active"`
		Amount   float64 `json:"amount"`
		Quantity int     `json:"quantity"`
	}
	t.Run(
		"struct required", func(t *testing.T) {
			v := New()
			v.Add("title").Required()
			v.Add("active").Required()
			v.Add("amount").Required()
			v.Add("quantity").Required()
			v.Json(testStruct{})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 4)
			assert.Equal(t, defaultRequiredMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["active"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"struct min", func(t *testing.T) {
			v := New()
			v.Add("title").Min(8)
			v.Add("amount").Min(8)
			v.Add("quantity").Min(8)
			v.Json(testStruct{})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 3)
			assert.Equal(t, defaultMinTextMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultMinNumberMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultMinNumberMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"struct max", func(t *testing.T) {
			v := New()
			v.Add("title").Max(2)
			v.Add("amount").Max(2)
			v.Add("quantity").Max(2)
			v.Json(testStruct{Title: "abc", Amount: 3, Quantity: 3})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 3)
			assert.Equal(t, defaultMaxTextMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultMaxNumberMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultMaxNumberMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"map required", func(t *testing.T) {
			v := New()
			v.Add("title").Required()
			v.Add("active").Required()
			v.Add("amount").Required()
			v.Add("quantity").Required()
			v.Json(map[string]any{})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 4)
			assert.Equal(t, defaultRequiredMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["active"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"map min", func(t *testing.T) {
			v := New()
			v.Add("title").Min(8)
			v.Add("amount").Min(8)
			v.Add("quantity").Min(8)
			v.Json(map[string]any{"title": "", "amount": 0, "quantity": 0})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 3)
			assert.Equal(t, defaultMinTextMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultMinNumberMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultMinNumberMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"map missing fields", func(t *testing.T) {
			v := New()
			v.Add("title").Min(8)
			v.Add("amount").Min(8)
			v.Add("quantity").Min(8)
			v.Json(map[string]any{})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 3)
			assert.Equal(t, defaultRequiredMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultRequiredMessage, v.Errors()["quantity"][0])
		},
	)
	t.Run(
		"map max", func(t *testing.T) {
			v := New()
			v.Add("title").Max(2)
			v.Add("amount").Max(2)
			v.Add("quantity").Max(2)
			v.Json(map[string]any{"title": "abc", "amount": 3, "quantity": 3})
			assert.False(t, v.Ok())
			assert.True(t, len(v.Errors()) == 3)
			assert.Equal(t, defaultMaxTextMessage, v.Errors()["title"][0])
			assert.Equal(t, defaultMaxNumberMessage, v.Errors()["amount"][0])
			assert.Equal(t, defaultMaxNumberMessage, v.Errors()["quantity"][0])
		},
	)
}
