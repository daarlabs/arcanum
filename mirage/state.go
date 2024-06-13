package mirage

import (
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type state struct {
	token                string
	exists               bool
	cache                cache.Client
	cookie               cookie.Cookie
	Components           map[string]any       `json:"components"`
	ComponentsExpiration map[string]time.Time `json:"componentsExpiration"`
	Messages             []Message            `json:"messages"`
}

const (
	stateCookieKey = "X-State"
	stateCacheKey  = "state"
)

var (
	stateDuration = 7 * 24 * time.Hour
)

func createState(cache cache.Client, cookie cookie.Cookie) *state {
	s := &state{
		cache:                cache,
		cookie:               cookie,
		Components:           make(map[string]any),
		ComponentsExpiration: make(map[string]time.Time),
		Messages:             make([]Message, 0),
	}
	s.token = cookie.Get(stateCookieKey)
	s.exists = len(s.token) > 0
	if !s.exists {
		s.token = uniuri.New()
	}
	if s.exists {
		cache.MustGet(stateCacheKey+":"+s.token, s)
		s.cleanComponents()
	}
	return s
}

func (s *state) save() error {
	s.cookie.Set(stateCookieKey, s.token, stateDuration)
	return s.cache.Set(stateCacheKey+":"+s.token, s, stateDuration)
}

func (s *state) mustSave() {
	err := s.save()
	if err != nil {
		panic(err)
	}
	s.cache.MustGet(stateCacheKey+":"+s.token, s)
}

func (s *state) cleanComponents() {
	t := time.Now()
	cleaned := false
	for name := range s.Components {
		expiration, ok := s.ComponentsExpiration[name]
		if !ok {
			continue
		}
		if t.After(expiration) {
			cleaned = true
			delete(s.Components, name)
			delete(s.ComponentsExpiration, name)
		}
	}
	if cleaned {
		s.mustSave()
	}
}
