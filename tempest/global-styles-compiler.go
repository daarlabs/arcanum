package tempest

import (
	"strings"
	"sync"
)

var (
	globalClassmaps struct {
		sync.RWMutex
		classmaps []*sync.Map
	}
)

var (
	globalStylesCache struct {
		sync.RWMutex
		styles string
		valid  bool
	}
)

func init() {
	globalClassmaps.classmaps = make([]*sync.Map, 0)
}

func buildStyles() string {
	globalStylesCache.RLock()
	if globalStylesCache.valid {
		globalStylesCache.RUnlock()
		return globalStylesCache.styles
	}
	globalStylesCache.RUnlock()
	globalStylesCache.Lock()
	if globalStylesCache.valid {
		globalStylesCache.Unlock()
		return globalStylesCache.styles
	}
	r := new(strings.Builder)
	for _, m := range globalClassmaps.classmaps {
		m.Range(
			func(key, value any) bool {
				r.WriteString(value.(string))
				r.WriteString(" ")
				return true
			},
		)
	}
	result := externalStyles + "\n" + baseStyles + "\n" + GlobalConfig.keyframes + "\n" + r.String()
	globalStylesCache.styles = result
	globalStylesCache.valid = true
	globalStylesCache.Unlock()
	return result
}

func invalidateStylesCache() {
	globalStylesCache.Lock()
	globalStylesCache.valid = false
	globalStylesCache.Unlock()
}

func initGlobalClassmaps() {
	globalClassmaps.classmaps = make([]*sync.Map, initialClassmapPriority)
	for i := range initialClassmapPriority {
		globalClassmaps.classmaps[i] = new(sync.Map)
	}
}

func ensureClassmapExists(priority int) {
	globalClassmaps.RLock()
	n := len(globalClassmaps.classmaps)
	if n > priority {
		globalClassmaps.RUnlock()
		return
	}
	globalClassmaps.RUnlock()
	globalClassmaps.Lock()
	defer globalClassmaps.Unlock()
	if n == priority {
		globalClassmaps.classmaps = append(globalClassmaps.classmaps, new(sync.Map))
		return
	}
	if n < priority {
		for _ = range priority - n + 1 {
			globalClassmaps.classmaps = append(globalClassmaps.classmaps, new(sync.Map))
		}
	}
}

func classExists(key string, priority int) bool {
	ensureClassmapExists(priority)
	_, ok := globalClassmaps.classmaps[priority].Load(key)
	return ok
}
