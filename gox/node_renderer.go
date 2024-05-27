package gox

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

type nodeRenderer struct {
	node
}

const (
	openBracket  = "{"
	closeBracket = "}"
)

func (n nodeRenderer) render() string {
	switch n.nodeType {
	case nodeFragment:
		var result string
		if len(n.attributes) > 0 {
			result += n.renderAttributes()
		}
		if len(n.children) > 0 {
			result += n.renderChildren()
		}
		return result
	case nodeText:
		return n.renderText()
	default:
		return n.renderElement()
	}
}

func (n nodeRenderer) renderElement() string {
	result := "<" + n.name
	isVoid := slices.Contains(voidElements, n.name)
	if len(n.attributes) > 0 {
		result += " " + n.renderAttributes()
	}
	if isVoid {
		result += " />"
	}
	if !isVoid {
		result += ">"
		if len(n.children) > 0 {
			result += n.renderChildren()
		}
		result += "</" + n.name + ">"
	}
	return result
}

func (n nodeRenderer) renderAttribute() string {
	if n.value == nil {
		return n.name
	}
	value := fmt.Sprintf("%v", n.value)
	if strings.HasPrefix(value, openBracket) && strings.HasSuffix(value, closeBracket) {
		return fmt.Sprintf(`%s='%s'`, n.name, value)
	}
	return fmt.Sprintf(`%s="%s"`, n.name, value)
}

func (n nodeRenderer) renderAttributes() string {
	size := len(n.attributes)
	var wg sync.WaitGroup
	wg.Add(size)
	result := make([]string, len(n.attributes))
	for i, a := range n.attributes {
		go func(i int, a node) {
			defer wg.Done()
			result[i] = nodeRenderer{a}.renderAttribute()
		}(i, a)
	}
	wg.Wait()
	return strings.Join(result, " ")
}

func (n nodeRenderer) renderChildren() string {
	size := len(n.children)
	var wg sync.WaitGroup
	wg.Add(size)
	result := make([]string, size)
	for i, ch := range n.children {
		go func(i int, ch node) {
			defer wg.Done()
			switch ch.nodeType {
			case nodeRaw:
				result[i] = nodeRenderer{ch}.renderRaw()
			case nodeElement:
				result[i] = nodeRenderer{ch}.renderElement()
			case nodeAttribute:
				result[i] = nodeRenderer{ch}.renderAttribute()
			case nodeText:
				result[i] = nodeRenderer{ch}.renderText()
			}
		}(i, ch)
	}
	wg.Wait()
	return strings.Join(result, "")
}

func (n nodeRenderer) renderText() string {
	return fmt.Sprintf(`%v`, n.value)
}

func (n nodeRenderer) renderRaw() string {
	return fmt.Sprintf(`%v`, n.value)
}
