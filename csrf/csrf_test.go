package csrf

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cache/memory"
	"github.com/daarlabs/arcanum/cookie"
)

func TestCsrf(t *testing.T) {
	t.Run(
		"create and get", func(t *testing.T) {
			path := "/test"
			name := "test-name"
			req := httptest.NewRequest(http.MethodGet, path, nil)
			res := httptest.NewRecorder()
			c := New(
				Cache(cache.New(context.Background(), memory.New(t.TempDir()), nil)),
				Cookie(cookie.New(req, res, path)),
			)
			token := c.MustCreate(Token{Name: name})
			r := c.MustGet(name, token)
			assert.Equal(t, name, r.Name)
		},
	)
	t.Run(
		"clean", func(t *testing.T) {
			path := "/test"
			name1 := "test-name-1"
			name2 := "test-name-2"
			req := httptest.NewRequest(http.MethodGet, path, nil)
			res := httptest.NewRecorder()
			c := New(
				Cache(cache.New(context.Background(), memory.New(t.TempDir()), nil)),
				Cookie(cookie.New(req, res, path)),
				Request(req),
			)
			token1 := c.MustCreate(Token{Name: name1})
			token2 := c.MustCreate(Token{Name: name2})
			req.Header.Set("Cookie", fmt.Sprintf("X-Csrf-%s=%s; X-Csrf-%s=%s", name1, token1, name2, token2))
			c.MustClean(name2)
			ok1, _ := c.Exists(name1, token1)
			ok2, _ := c.Exists(name2, token2)
			assert.False(t, ok1)
			assert.False(t, ok2)
			assert.True(t, strings.HasPrefix(res.Header().Get("Set-Cookie"), "X-Csrf-"+name1+"="+token1))
		},
	)
}
