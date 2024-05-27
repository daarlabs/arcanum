package dd

import (
	"fmt"
	"strings"
)

func Print(values ...interface{}) string {
	debug := make([]string, len(values))
	for i, v := range values {
		debug[i] = fmt.Sprintf("%+v", v)
	}
	fmt.Println(strings.Join(debug, " "))
	return ""
}
