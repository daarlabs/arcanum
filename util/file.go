package util

import "strings"

func GetFilenameSuffix(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}
