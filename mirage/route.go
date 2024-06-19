package mirage

import (
	"regexp"
	
	"github.com/daarlabs/arcanum/firewall"
)

type RouteConfig struct {
	Type  int
	Value any
}

type Route struct {
	Lang       string
	Path       string
	Name       string
	Layout     layoutFactory
	Matcher    *regexp.Regexp
	Methods    []string
	Firewall   []firewall.Firewall
	PathValues []string
}

const (
	routeMethod = iota
	routeName
	routeLayout
)

func Method(method ...string) RouteConfig {
	return RouteConfig{
		Type:  routeMethod,
		Value: method,
	}
}

func Name(name string) RouteConfig {
	return RouteConfig{
		Type:  routeName,
		Value: name,
	}
}

func Layout(name string) RouteConfig {
	return RouteConfig{
		Type:  routeLayout,
		Value: name,
	}
}
