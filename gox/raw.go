package gox

func Raw(html string) Node {
	return node{
		nodeType: nodeRaw,
		value:    html,
	}
}
