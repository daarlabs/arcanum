package mirage

import (
	"reflect"
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type State interface {
	Get(key string, target any) error
	Save(key string, value any) error
	
	MustGet(key string, target any)
	MustSave(key string, value any)
}

type state struct {
	token                string
	exists               bool
	cache                cache.Client
	cookie               cookie.Cookie
	Components           map[string]any                 `json:"components"`
	ComponentsExpiration map[string]time.Time           `json:"componentsExpiration"`
	Forms                map[string]map[string][]string `json:"forms"`
	Messages             []Message                      `json:"messages"`
	Customs              map[string]any                 `json:"customs"`
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
		Forms:                make(map[string]map[string][]string),
		Messages:             make([]Message, 0),
		Customs:              make(map[string]any),
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

func (s *state) Get(key string, target any) error {
	v, ok := s.Customs[key]
	if !ok {
		return nil
	}
	tt := reflect.TypeOf(target)
	if tt.Kind() != reflect.Ptr {
		return ErrorNoPtr
	}
	vt := reflect.TypeOf(v)
	if tt.Elem() != vt {
		return ErrorMismatchType
	}
	reflect.ValueOf(target).Elem().Set(reflect.ValueOf(v))
	delete(s.Customs, key)
	return s.save()
}

func (s *state) MustGet(key string, target any) {
	err := s.Get(key, target)
	if err != nil {
		panic(err)
	}
}

func (s *state) Save(key string, value any) error {
	s.Customs[key] = value
	return s.save()
}

func (s *state) MustSave(key string, value any) {
	err := s.Save(key, value)
	if err != nil {
		panic(err)
	}
}

func (s *state) GetForm(key string) (map[string][]string, error) {
	result, ok := s.Forms[key]
	if !ok {
		return make(map[string][]string), nil
	}
	delete(s.Forms, key)
	return result, s.save()
}

func (s *state) MustGetForm(key string) map[string][]string {
	form, err := s.GetForm(key)
	if err != nil {
		panic(err)
	}
	return form
}

func (s *state) SaveForm(key string, form map[string][]string) error {
	s.Forms[key] = form
	return s.save()
}

func (s *state) MustSaveForm(key string, form map[string][]string) {
	err := s.SaveForm(key, form)
	if err != nil {
		panic(err)
	}
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
