package tempest

import (
	"strings"
	"sync"
)

var (
	namedClassmaps struct {
		sync.RWMutex
		classmaps map[string][]*sync.Map
	}
)

var (
	namedStylesCache struct {
		sync.RWMutex
		styles map[string]string
		valid  map[string]bool
	}
)

func init() {
	namedClassmaps.classmaps = make(map[string][]*sync.Map)
	namedStylesCache.styles = make(map[string]string)
	namedStylesCache.valid = make(map[string]bool)
}

func buildNamedStyles(name string) string {
	if _, ok := namedStylesCache.valid[name]; !ok {
		namedStylesCache.Lock()
		namedStylesCache.valid[name] = false
		namedStylesCache.Unlock()
	}
	if _, ok := namedStylesCache.styles[name]; !ok {
		namedStylesCache.Lock()
		namedStylesCache.styles[name] = ""
		namedStylesCache.Unlock()
	}
	namedStylesCache.RLock()
	if valid, ok := namedStylesCache.valid[name]; ok && valid {
		namedStylesCache.RUnlock()
		return namedStylesCache.styles[name]
	}
	namedStylesCache.RUnlock()
	namedStylesCache.Lock()
	if globalStylesCache.valid {
		namedStylesCache.Unlock()
		return namedStylesCache.styles[name]
	}
	r := new(strings.Builder)
	for _, m := range namedClassmaps.classmaps[name] {
		m.Range(
			func(key, value any) bool {
				r.WriteString(value.(string))
				r.WriteString(" ")
				return true
			},
		)
	}
	result := r.String()
	namedStylesCache.styles[name] = result
	namedStylesCache.valid[name] = true
	namedStylesCache.Unlock()
	return result
}

func invalidateNamedStylesCache(name string) {
	namedStylesCache.Lock()
	namedStylesCache.valid[name] = false
	namedStylesCache.Unlock()
}

func ensureNamedClassmapExists(name string, priority int) {
	if _, ok := namedClassmaps.classmaps[name]; !ok {
		namedClassmaps.Lock()
		namedClassmaps.classmaps[name] = make([]*sync.Map, 0)
		namedClassmaps.Unlock()
	}
	namedClassmaps.RLock()
	n := len(namedClassmaps.classmaps[name])
	if n > priority {
		namedClassmaps.RUnlock()
		return
	}
	namedClassmaps.RUnlock()
	namedClassmaps.Lock()
	defer namedClassmaps.Unlock()
	if n == priority {
		namedClassmaps.classmaps[name] = append(namedClassmaps.classmaps[name], new(sync.Map))
		return
	}
	if n < priority {
		for _ = range priority - n + 1 {
			namedClassmaps.classmaps[name] = append(namedClassmaps.classmaps[name], new(sync.Map))
		}
	}
}

func namedClassExists(name string, key string, priority int) bool {
	ensureNamedClassmapExists(name, priority)
	_, ok := namedClassmaps.classmaps[name][priority].Load(key)
	return ok
}
