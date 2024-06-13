package mirage

import (
	"compress/gzip"
	"net/http"
)

type compressedResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w compressedResponseWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func createCompressedWriter(w http.ResponseWriter) compressedResponseWriter {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Vary", "Accept-Encoding")
	return compressedResponseWriter{Writer: gzip.NewWriter(w), ResponseWriter: w}
}
