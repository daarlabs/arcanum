package util

import (
	"net/http"
	"strings"
	
	"github.com/daarlabs/arcanum/util/constant/contentType"
	"github.com/daarlabs/arcanum/util/constant/header"
)

func IsRequestMultipart(req *http.Request) bool {
	return strings.Contains(req.Header.Get(header.ContentType), contentType.MultipartForm)
}
