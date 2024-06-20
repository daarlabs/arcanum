package form

import (
	"net/http"
	"strconv"
	"strings"
)

func isRequestForm(req *http.Request) bool {
	formType := req.Header.Get(contentType)
	return strings.Contains(formType, contentTypeForm)
}

func isRequestMultipartForm(req *http.Request) bool {
	formType := req.Header.Get(contentType)
	return strings.Contains(formType, contentTypeMultipartForm)
}

func isFormSubmitted(req *http.Request) bool {
	return isRequestForm(req) || isRequestMultipartForm(req)
}

func convertSlice[S, R any](ts []S, f func(S) R) []R {
	us := make([]R, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func convertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return r
}

func convertToInt64(v string) int64 {
	r, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return r
}

func convertToFloat(v string) float64 {
	r, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}
	return r
}

func getFileSuffixFromName(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}
