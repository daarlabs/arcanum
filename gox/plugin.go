package gox

import (
	"fmt"
	"strings"
)

// Clsx plugin
// Conditional classes rendering
type Clsx map[any]bool

func (c Clsx) Node() Node {
	return Class(c.String())
}

func (c Clsx) Merge(items ...Clsx) Clsx {
	for _, item := range items {
		for k, v := range item {
			c[k] = v
		}
	}
	return c
}

func (c Clsx) String() string {
	result := make([]string, 0)
	for class, use := range c {
		if !use {
			continue
		}
		switch v := class.(type) {
		case string:
			result = append(result, v)
		case fmt.Stringer:
			result = append(result, v.String())
		default:
			result = append(result, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(result, " ")
}
