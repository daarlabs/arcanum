package auth

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/dchest/uniuri"
	"github.com/matthewhartstonge/argon2"
	
	"github.com/daarlabs/arcanum/quirk"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type Manager interface {
	Session() SessionManager
	Tfa() TfaManager
	User() UserManager
	CustomUser(id int, email string) UserManager
	
	In(email, password string) (In, error)
	Out() error
	
	MustIn(email, password string) In
	MustOut()
}

type In struct {
	Token string
	Ok    bool
	Tfa   bool `json:"tfa"`
}

type manager struct {
	db     *quirk.DB
	req    *http.Request
	res    http.ResponseWriter
	cookie cookie.Cookie
	cache  cache.Client
	config Config
}

func New(
	db *quirk.DB,
	req *http.Request,
	res http.ResponseWriter,
	cookie cookie.Cookie,
	cache cache.Client,
	config Config,
) Manager {
	return &manager{
		db:     db,
		req:    req,
		res:    res,
		cookie: cookie,
		cache:  cache,
		config: config,
	}
}

func (m *manager) In(email, password string) (In, error) {
	var r User
	err := quirk.New(m.db).
		Q(fmt.Sprintf(`SELECT id, email, roles, password, tfa FROM %s`, usersTable)).
		Q("WHERE email = @email", quirk.Map{"email": email}).
		Q("AND active = true").
		Exec(&r)
	if err != nil {
		return In{
			Ok:  false,
			Tfa: false,
		}, err
	}
	if r.Id == 0 {
		return In{
			Ok:  false,
			Tfa: false,
		}, ErrorInvalidCredentials
	}
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(r.Password))
	if !ok || err != nil {
		return In{
			Ok:  false,
			Tfa: false,
		}, err
	}
	if r.Tfa {
		token := uniuri.New()
		if err := m.cache.Set(createTfaCacheKey(token), User{Id: r.Id}, time.Minute*5); err != nil {
			return In{}, err
		}
		m.cookie.Set(TfaCookieKey, token, time.Minute*5)
		return In{
			Token: token,
			Ok:    true,
			Tfa:   true,
		}, nil
	}
	token, err := m.Session().New(r)
	if err != nil {
		return In{}, err
	}
	return In{
		Token: token,
		Ok:    true,
		Tfa:   false,
	}, nil
}

func (m *manager) MustIn(email, password string) In {
	r, err := m.In(email, password)
	if err != nil {
		panic(err)
	}
	return r
}

func (m *manager) Out() error {
	return m.Session().Destroy()
}

func (m *manager) MustOut() {
	err := m.Out()
	if err != nil {
		panic(err)
	}
}

func (m *manager) Session() SessionManager {
	return createSessionManager(
		m.req,
		m.res,
		m.cookie,
		m.cache,
		m.config,
	)
}

func (m *manager) Tfa() TfaManager {
	return createTfaManager(m)
}

func (m *manager) User() UserManager {
	session := m.Session().MustGet()
	return CreateUserManager(m.db, m.cache, session.Id, session.Email)
}

func (m *manager) CustomUser(id int, email string) UserManager {
	return CreateUserManager(m.db, m.cache, id, email)
}
