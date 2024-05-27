package form

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestFormBuilder(t *testing.T) {
	email, quantity, amount, checked := "test@test.cz", 2, 5.55, true
	t.Run(
		"build", func(t *testing.T) {
			testForm, err := Build[testForm](
				New(
					Add("email").With(Text(email)),
				),
			)
			assert.Nil(t, err)
			assert.Equal(t, email, testForm.Email.Value)
		},
	)
	t.Run(
		"with request", func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				"/test",
				strings.NewReader(fmt.Sprintf("email=%s&quantity=%d&amount=%.2f&checked=on", email, quantity, amount)),
			)
			req.Header.Set(contentType, contentTypeForm)
			test, err := Build[testForm](
				New(
					Add("email").With(Text(), Validate.Required()),
					Add("amount").With(Number[float64]()),
					Add("checked").With(Checkbox()),
				).Request(req),
			)
			assert.Nil(t, err)
			assert.Equal(t, true, test.Valid)
			assert.Equal(t, true, test.Submitted)
			assert.Equal(t, email, test.Email.Value)
			assert.Equal(t, amount, test.Amount.Value)
			assert.Equal(t, checked, test.Checked.Value)
		},
	)
	t.Run(
		"with multipart request", func(t *testing.T) {
			fileBytes := bytes.Repeat([]byte("test"), 1<<8)
			bodyBuf := new(bytes.Buffer)
			bodyWriter := multipart.NewWriter(bodyBuf)
			err := bodyWriter.WriteField("email", email)
			assert.Nil(t, err)
			testFile, err := bodyWriter.CreateFormFile("test", "test.txt")
			assert.Nil(t, err)
			_, err = testFile.Write(fileBytes)
			assert.Nil(t, err)
			assert.Nil(t, bodyWriter.Close())
			req := httptest.NewRequest(
				http.MethodPost,
				"/test",
				bodyBuf,
			)
			req.Header.Set(contentType, bodyWriter.FormDataContentType())
			testForm, err := Build[testForm](
				New(
					Add("email").With(Text()),
					Add("test").With(File()),
				).Request(req),
			)
			assert.Nil(t, err)
			assert.Equal(t, true, testForm.Valid)
			assert.Equal(t, true, testForm.Submitted)
			assert.Equal(t, email, testForm.Email.Value)
			assert.Equal(t, fileBytes, testForm.Test.Value.Data)
		},
	)
}
