package firewall

import (
	"regexp"
	"strings"
	
	"github.com/daarlabs/arcanum/auth"
)

type Firewall struct {
	Enabled  bool
	Name     string
	Groups   []string
	Matchers []*regexp.Regexp
	Paths    []string
	Redirect string
	Roles    []auth.Role
	Secret   string
}

type Attempt struct {
	Roles  []auth.Role
	Secret string
}

type Result struct {
	Ok       bool
	Err      error
	Redirect string
}

func New(configs ...Config) *Firewall {
	f := &Firewall{
		Groups:   make([]string, 0),
		Matchers: make([]*regexp.Regexp, 0),
		Paths:    make([]string, 0),
		Roles:    make([]auth.Role, 0),
	}
	for _, item := range configs {
		c, ok := item.(*config)
		if !ok {
			continue
		}
		switch c.name {
		case configEnabled:
			f.Enabled = c.value.(bool)
		case configName:
			f.Name = c.value.(string)
		case configGroup:
			f.Groups = append(f.Groups, c.value.([]string)...)
		case configMatcher:
			f.Matchers = append(f.Matchers, c.value.([]*regexp.Regexp)...)
		case configPath:
			f.Paths = append(f.Paths, c.value.([]string)...)
		case configRedirect:
			f.Redirect = c.value.(string)
		case configRole:
			f.Roles = append(f.Roles, c.value.([]auth.Role)...)
		case configSecret:
			f.Secret = c.value.(string)
		}
	}
	return f
}

func (f *Firewall) Try(attempt Attempt) Result {
	if !f.Enabled {
		return Result{Ok: true}
	}
	if len(f.Secret) > 0 && f.Secret == attempt.Secret {
		return Result{Ok: true}
	}
	if len(f.Roles) > 0 {
		for _, fr := range f.Roles {
			for _, ar := range attempt.Roles {
				if fr.Compare(ar) {
					return Result{Ok: true}
				}
			}
		}
	}
	if len(f.Redirect) > 0 {
		return Result{Redirect: f.Redirect}
	}
	return Result{Ok: true}
}

func (f *Firewall) Match(path string) bool {
	for _, matcher := range f.Matchers {
		if matcher.MatchString(path) {
			return true
		}
	}
	return false
}

func (f *Firewall) MatchPath(path string) bool {
	for _, firewallPath := range f.Paths {
		if len(path) > 0 && strings.HasPrefix(path, firewallPath) {
			return true
		}
	}
	return false
}

func (f *Firewall) MatchGroup(group string) bool {
	for _, firewallGroup := range f.Groups {
		if len(group) > 0 && strings.HasPrefix(group, firewallGroup) {
			return true
		}
	}
	return false
}
