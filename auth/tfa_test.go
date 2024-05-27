package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/dchest/uniuri"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
	
	"github.com/daarlabs/arcanum/quirk"
)

func TestTfa(t *testing.T) {
	db, redis := createTestDatabaseConnection(t), createTestRedisConnection(t)
	assert.NotNil(t, db)
	assert.NotNil(t, redis)
	assert.NoError(t, DropTable(quirk.New(db)))
	assert.NoError(t, CreateTable(quirk.New(db)))
	um := CreateUserManager(db, nil, 0, "")
	_, err := um.Create(
		User{
			Active:   true,
			Roles:    []string{"owner"},
			Email:    "dominik@linduska.dev",
			Password: "123456789",
		},
	)
	assert.NoError(t, err)
	r, err := um.Get()
	assert.NoError(t, err)
	assert.True(t, r.Id > 0)
	t.Cleanup(
		func() {
			assert.NoError(t, DropTable(quirk.New(db)))
		},
	)
	t.Run(
		"pending verification", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			res := httptest.NewRecorder()
			m := &manager{
				db:     db,
				req:    req,
				res:    res,
				cookie: cookie.New(req, res, "/"),
				cache:  cache.New(context.Background(), nil, redis),
				config: Config{},
			}
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", TfaCookieKey, uniuri.New()))
			tm := createTfaManager(m)
			assert.True(t, tm.MustGetPendingVerification())
		},
	)
	t.Run(
		"enable", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			res := httptest.NewRecorder()
			m := &manager{
				db:     db,
				req:    req,
				res:    res,
				cookie: cookie.New(req, res, "/"),
				cache:  cache.New(context.Background(), nil, redis),
				config: Config{},
			}
			auth, err := m.In("dominik@linduska.dev", "123456789")
			assert.NoError(t, err)
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", TfaCookieKey, uniuri.New()))
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", SessionCookieKey, auth.Token))
			assert.NoError(t, m.Tfa().Enable())
			auth, err = m.In("dominik@linduska.dev", "123456789")
			assert.NoError(t, err)
			assert.True(t, auth.Tfa)
		},
	)
	t.Run(
		"verify", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			res := httptest.NewRecorder()
			m := &manager{
				db:     db,
				req:    req,
				res:    res,
				cookie: cookie.New(req, res, "/"),
				cache:  cache.New(context.Background(), nil, redis),
				config: Config{},
			}
			auth, err := m.In("dominik@linduska.dev", "123456789")
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", TfaCookieKey, auth.Token))
			var u User
			assert.NoError(
				t, quirk.New(db).Q(
					`SELECT tfa_secret FROM users WHERE email = @email`, quirk.Map{"email": "dominik@linduska.dev"},
				).Exec(&u),
			)
			otp, err := totp.GenerateCode(u.TfaSecret.V, time.Now())
			assert.NoError(t, err)
			token, err := m.Tfa().Verify(otp)
			assert.NoError(t, err)
			assert.True(t, len(token) > 0)
		},
	)
	t.Run(
		"disable", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			res := httptest.NewRecorder()
			m := &manager{
				db:     db,
				req:    req,
				res:    res,
				cookie: cookie.New(req, res, "/"),
				cache:  cache.New(context.Background(), nil, redis),
				config: Config{},
			}
			auth, err := m.In("dominik@linduska.dev", "123456789")
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", TfaCookieKey, auth.Token))
			var u User
			assert.NoError(
				t, quirk.New(db).Q(
					`SELECT tfa_secret FROM users WHERE email = @email`, quirk.Map{"email": "dominik@linduska.dev"},
				).Exec(&u),
			)
			otp, err := totp.GenerateCode(u.TfaSecret.V, time.Now())
			assert.NoError(t, err)
			token, err := m.Tfa().Verify(otp)
			assert.NoError(t, err)
			m.req.Header.Add("Cookie", fmt.Sprintf("%s=%s", SessionCookieKey, token))
			assert.NoError(t, m.Tfa().Disable())
			auth, err = m.In("dominik@linduska.dev", "123456789")
			assert.NoError(t, err)
			assert.False(t, auth.Tfa)
		},
	)
}
