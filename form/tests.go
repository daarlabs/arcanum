package form

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
)

type testForm struct {
	Form
	Roles    Field[[]string]
	Email    Field[string]
	Name     Field[string]
	Quantity Field[int]
	Amount   Field[float64]
	Checked  Field[bool]
	Test     Field[Multipart]
}

type testModel struct {
	Roles    []string
	Name     string
	Quantity int
	Amount   float64
	Checked  bool
}

const (
	testAction        = "/test"
	testName          = "test"
	testNameValue     = "Test"
	testQuantityValue = 5
	testAmountValue   = 999.99
	testCheckedValue  = true
	testRole          = "owner"
)

func testGetRequest() *http.Request {
	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		strings.NewReader(""),
	)
	req.Header.Set(contentType, contentTypeHtml)
	return req
}

func testCreateFormRequest() *http.Request {
	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		strings.NewReader(
			fmt.Sprintf(
				"name=%s&quantity=%d&amount=%.2f&checked=on&roles=%s",
				testNameValue, testQuantityValue, testAmountValue, testRole,
			),
		),
	)
	req.Header.Set(contentType, contentTypeForm)
	return req
}

func testCreateMultipartRequest() ([]byte, *http.Request, error) {
	fileBytes := bytes.Repeat([]byte("test"), 1<<8)
	bodyBuf := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuf)
	if err := bodyWriter.WriteField("name", testNameValue); err != nil {
		return fileBytes, nil, err
	}
	testFile, err := bodyWriter.CreateFormFile("test", "test.txt")
	if err != nil {
		return fileBytes, nil, err
	}
	_, err = testFile.Write(fileBytes)
	if err != nil {
		return fileBytes, nil, err
	}
	if err := bodyWriter.Close(); err != nil {
		return fileBytes, nil, err
	}
	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		bodyBuf,
	)
	req.Header.Set(contentType, bodyWriter.FormDataContentType())
	return fileBytes, req, nil
}

func testCreateEmptyBuildRequest() *http.Request {
	req := httptest.NewRequest(
		http.MethodPost,
		"/test",
		strings.NewReader(""),
	)
	req.Header.Set(contentType, contentTypeForm)
	return req
}
