package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
	
	"github.com/daarlabs/arcanum/quirk"
)

func TestAuth(t *testing.T) {
	var sessionCookie, tfaCookie *http.Cookie
	db, redis := createTestDatabaseConnection(t), createTestRedisConnection(t)
	assert.NoError(t, DropTable(quirk.New(db)))
	assert.NoError(t, CreateTable(quirk.New(db)))
	user := User{
		Active:   true,
		Roles:    []string{"owner"},
		Email:    "dominik@linduska.dev",
		Password: "123456789",
	}
	createTestManager := func(req *http.Request, res http.ResponseWriter) *manager {
		return &manager{
			db:     db,
			req:    req,
			res:    res,
			cookie: cookie.New(req, res, "/"),
			cache:  cache.New(context.Background(), nil, redis),
			config: Config{},
		}
	}
	t.Cleanup(
		func() {
			assert.NoError(t, DropTable(quirk.New(db)))
		},
	)
	t.Run(
		"create user", func(t *testing.T) {
			id, err := CreateUserManager(db, nil, 0, "").Create(user)
			assert.NoError(t, err)
			user.Id = id
			assert.Equal(t, true, id != 0)
		},
	)
	t.Run(
		"auth in", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()
			m := createTestManager(req, res)
			auth, err := m.In("dominik@linduska.dev", "123456789")
			assert.NoError(t, err)
			cookies := res.Result().Cookies()
			sessionCookie = cookies[0]
			assert.Equal(t, 1, len(cookies))
			assert.Equal(t, true, len(cookies[0].Value) > 0)
			assert.Equal(t, true, auth.Ok)
		},
	)
	t.Run(
		"enable tfa", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/enable/tfa", nil)
			res := httptest.NewRecorder()
			req.AddCookie(sessionCookie)
			m := createTestManager(req, res)
			assert.NoError(t, m.Tfa().Enable())
			u, err := m.User().Get()
			assert.NoError(t, err)
			assert.Equal(t, true, u.Tfa)
			assert.Equal(t, true, len(u.TfaSecret.V) > 0)
			assert.Equal(t, true, len(u.TfaUrl.V) > 0)
			assert.Equal(t, true, len(u.TfaCodes.V) > 0)
		},
	)
	t.Run(
		"auth in tfa", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/auth/in/tfa", nil)
			res := httptest.NewRecorder()
			m := createTestManager(req, res)
			a, err := m.In("dominik@linduska.dev", "123456789")
			assert.NoError(t, err)
			cookies := res.Result().Cookies()
			tfaCookie = cookies[0]
			req.AddCookie(tfaCookie)
			assert.Equal(t, true, a.Tfa)
			assert.Equal(t, true, len(tfaCookie.Value) > 0)
			var u User
			assert.NoError(t, m.cache.Get(createTfaCacheKey(tfaCookie.Value), &u))
			assert.Equal(t, true, u.Id > 0)
			quirk.New(m.db).Q(`SELECT tfa_secret FROM users WHERE id = @id`, quirk.Map{"id": u.Id}).MustExec(&u)
			otp, err := totp.GenerateCode(u.TfaSecret.V, time.Now())
			assert.NoError(t, err)
			_, err = m.Tfa().Verify(otp)
			assert.NoError(t, err)
		},
	)
	t.Run(
		"auth out", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/out", nil)
			res := httptest.NewRecorder()
			req.AddCookie(sessionCookie)
			m := createTestManager(req, res)
			session, err := m.Session().Get()
			assert.NoError(t, err)
			assert.Equal(t, true, session.Id > 0)
			assert.NoError(t, m.Out())
			time.Sleep(time.Millisecond * 10)
			session, err = m.Session().Get()
			assert.NoError(t, err)
			assert.Equal(t, true, session.Id < 1)
			sessionCookie = nil
		},
	)
	t.Run(
		"update user", func(t *testing.T) {
			u := User{Active: false}
			assert.NoError(t, CreateUserManager(db, nil, user.Id, "dominik@linduska.dev").Update(u, "active"))
			u, err := CreateUserManager(db, nil, user.Id, "dominik@linduska.dev").Get()
			assert.NoError(t, err)
			assert.Equal(t, false, u.Active)
		},
	)
}
