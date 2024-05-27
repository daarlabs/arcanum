package gox

import (
	"strings"
)

// Clsx plugin
// Conditional classes rendering
type Clsx map[string]bool

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
	for classes, use := range c {
		if !use {
			continue
		}
		result = append(result, classes)
	}
	return strings.Join(result, " ")
}
