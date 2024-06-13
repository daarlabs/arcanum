package util

import (
	"net/url"
	"strings"
)

func JoinPath(basePath string, path string) string {
	if strings.Contains(path, "...") {
		return strings.TrimSuffix(basePath, "/") + "/" + strings.TrimPrefix(path, "/")
	}
	p, err := url.JoinPath(basePath, path)
	if err != nil {
		panic(err)
	}
	return p
}

func MustJoinPath(basePath string, path string) string {
	p, err := url.JoinPath(basePath, path)
	if err != nil {
		panic(err)
	}
	return p
}
