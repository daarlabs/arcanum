package mirage

import "strings"

func createDividedName(items ...string) string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = strings.ReplaceAll(item, "_", "-")
	}
	return strings.Join(result, namePrefixDivider)
}
