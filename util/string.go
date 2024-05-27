package util

import "strings"

func EscapeString(value string) string {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "'", "", "\"", "", "`", "")
	value = replacer.Replace(value)
	return value
}
