package auth

import (
	"net/http"
	"slices"
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type SessionManager interface {
	Token() string
	Exists() (bool, error)
	Get(token ...string) (Session, error)
	New(user User) (string, error)
	Renew() error
	Destroy() error
	
	MustExists() bool
	MustGet(token ...string) Session
	MustNew(user User) string
	MustRenew()
	MustDestroy()
}

type Session struct {
	Id        int      `json:"id"`
	Email     string   `json:"email"`
	Roles     []string `json:"role"`
	Super     bool     `json:"super"`
	Ip        string   `json:"ip"`
	UserAgent string   `json:"userAgent"`
}

type sessionManager struct {
	req    *http.Request
	res    http.ResponseWriter
	cookie cookie.Cookie
	cache  cache.Client
	config Config
}

const (
	SessionCookieKey = "X-Session"
	DefaultDuration  = 24 * time.Hour
)

func createSessionManager(
	req *http.Request,
	res http.ResponseWriter,
	cookie cookie.Cookie,
	cache cache.Client,
	config Config,
) SessionManager {
	return &sessionManager{
		req:    req,
		res:    res,
		cookie: cookie,
		cache:  cache,
		config: config,
	}
}

func (s sessionManager) Token() string {
	return s.cookie.Get(SessionCookieKey)
}

func (s sessionManager) Exists() (bool, error) {
	token := s.Token()
	if len(token) == 0 {
		return false, ErrorMissingSessionCookie
	}
	return len(token) > 0 && s.cache.Exists(createSessionCacheKey(token)), nil
}

func (s sessionManager) MustExists() bool {
	exists, err := s.Exists()
	if err != nil {
		panic(err)
	}
	return exists
}

func (s sessionManager) Get(token ...string) (Session, error) {
	var t string
	var r Session
	if len(token) == 0 {
		t = s.Token()
	}
	if len(token) > 0 {
		t = token[0]
	}
	err := s.cache.Get(createSessionCacheKey(t), &r)
	return r, err
}

func (s sessionManager) MustGet(token ...string) Session {
	r, err := s.Get(token...)
	if err != nil {
		panic(err)
	}
	return r
}

func (s sessionManager) New(user User) (string, error) {
	token := uniuri.New()
	if s.config.Duration.Hours() == 0 {
		s.config.Duration = DefaultDuration
	}
	s.cookie.Set(SessionCookieKey, token, s.config.Duration)
	return token, s.cache.Set(createSessionCacheKey(token), s.createSession(user), s.config.Duration)
}

func (s sessionManager) MustNew(user User) string {
	token, err := s.New(user)
	if err != nil {
		panic(err)
	}
	return token
}

func (s sessionManager) Renew() error {
	token := s.Token()
	if len(token) == 0 {
		return ErrorMissingSessionCookie
	}
	session, err := s.Get()
	if err != nil {
		return err
	}
	if session.Id == 0 || session.Ip != s.getIp() || session.UserAgent != s.getUserAgent() {
		return ErrorCredentialsMismatch
	}
	if s.config.Duration.Hours() == 0 {
		s.config.Duration = DefaultDuration
	}
	s.cookie.Set(SessionCookieKey, token, s.config.Duration)
	return s.cache.Set(createSessionCacheKey(token), session, s.config.Duration)
}

func (s sessionManager) MustRenew() {
	err := s.Renew()
	if err != nil {
		panic(err)
	}
}

func (s sessionManager) Destroy() error {
	token := s.Token()
	s.cookie.Set(SessionCookieKey, "", time.Millisecond)
	return s.cache.Set(createSessionCacheKey(token), "", time.Millisecond)
}

func (s sessionManager) MustDestroy() {
	if err := s.Destroy(); err != nil {
		panic(err)
	}
}

func (s sessionManager) createSession(user User) Session {
	return Session{
		Id:        user.Id,
		Email:     user.Email,
		Ip:        s.getIp(),
		UserAgent: s.getUserAgent(),
		Roles:     user.Roles,
		Super:     s.containsSuperRole(user.Roles...),
	}
}

func (s sessionManager) getIp() string {
	return s.req.Header.Get("X-Forwarded-For")
}

func (s sessionManager) getUserAgent() string {
	return s.req.Header.Get("User-Agent")
}

func (s sessionManager) containsSuperRole(roles ...string) bool {
	for _, r := range s.config.Roles {
		if slices.Contains(roles, r.Name) && r.Super {
			return true
		}
	}
	return false
}
