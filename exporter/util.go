package exporter

import (
	"fmt"
)

const (
	gotenbergHtmlToPdfPath = "/forms/chromium/convert/html"
)

var (
	letters = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
	}
	lettersLen             = len(letters)
	maxLettersCombinations = lettersLen * lettersLen
)

func createStringSliceFromRow(row *row) []string {
	r := make([]string, len(row.cols))
	for i, c := range row.cols {
		r[i] = fmt.Sprintf("%v", c.value)
	}
	return r
}

func getLetterWithIndex(index int) string {
	if index > maxLettersCombinations {
		return ""
	}
	if index > lettersLen-1 {
		var secondLetterIndex int
		modulo := index % (lettersLen - 1)
		if modulo > 0 {
			secondLetterIndex = modulo - 1
		}
		if modulo == 0 {
			secondLetterIndex = lettersLen - 1
		}
		firstLetterIndex := ((index - secondLetterIndex) / (lettersLen - 1)) - 1
		if firstLetterIndex > (lettersLen-1) || secondLetterIndex > (lettersLen-1) {
			return ""
		}
		return letters[firstLetterIndex] + letters[secondLetterIndex]
	}
	return letters[index]
}

func createGotenbergEndpoint(host string) string {
	return host + gotenbergHtmlToPdfPath
}
