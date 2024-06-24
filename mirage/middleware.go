package mirage

import (
	"errors"
	"net/http"
	"slices"
	"strings"
	
	"github.com/daarlabs/arcanum/auth"
	"github.com/daarlabs/arcanum/firewall"
	"github.com/daarlabs/arcanum/form"
)

func createLangMiddleware() Handler {
	return func(c Ctx) error {
		cfg := c.Config().Localization
		if !cfg.Enabled || len(cfg.Languages) == 0 {
			return c.Continue()
		}
		var validLang bool
		current := c.Lang().Current()
		for _, l := range cfg.Languages {
			if l.Code == current {
				validLang = true
			}
		}
		if validLang {
			c.Cookie().Set(langCookieKey, c.Lang().Current(), langCookieDuration)
		}
		if !validLang {
			mainLang := c.Lang().Main()
			c.Cookie().Set(langCookieKey, mainLang, langCookieDuration)
			if cfg.Path {
				return c.Response().Status(http.StatusMovedPermanently).Redirect("/" + mainLang + "/")
			}
		}
		if len(c.Request().QueryParam(langQueryKey)) > 0 {
			return c.Response().Refresh()
		}
		return c.Continue()
	}
}

func createFormMiddleware() Handler {
	return func(c Ctx) error {
		if c.Request().Is().Get() {
			return c.Continue()
		}
		if _, err := form.ParseForm(c.Request().Raw(), c.Config().Form.Limit); err != nil {
			return c.Response().Error(err)
		}
		return c.Continue()
	}
}

func createCsrfMiddleware() Handler {
	return func(c Ctx) error {
		if c.Request().Is().Get() {
			if c.Request().Is().Action() {
				return c.Continue()
			}
			path := c.Request().Path()
			if strings.HasPrefix(path, tempestAssetsPath) || strings.HasPrefix(path, c.Config().App.Public) {
				return c.Continue()
			}
			if err := c.Csrf().Clean(c.Request().Name()); err != nil {
				return c.Response().Refresh()
			}
			return c.Continue()
		}
		if c.Auth().Session().MustExists() {
			return c.Continue()
		}
		token := c.Request().Form().Get(form.CsrfToken)
		name := c.Request().Form().Get(form.CsrfName)
		if len(token) == 0 {
			return c.Response().Refresh()
		}
		t, err := c.Csrf().Get(name, token)
		if err != nil {
			return err
		}
		if !t.Exists {
			return c.Response().Refresh()
		}
		if t.Name != name || t.UserAgent != c.Request().UserAgent() || t.Ip != c.Request().Ip() {
			return c.Response().Refresh()
		}
		if err := c.Csrf().Destroy(t); err != nil {
			return c.Response().Refresh()
		}
		return c.Continue()
	}
}

func createFirewallMiddleware(firewalls []firewall.Firewall) Handler {
	return func(c Ctx) error {
		if len(firewalls) == 0 {
			return c.Continue()
		}
		var err error
		session, err := c.Auth().Session().Get()
		if err != nil {
			return c.Response().Error(err)
		}
		if session.Super {
			if err := c.Auth().Session().Renew(); err != nil {
				c.Auth().MustOut()
				return c.Response().Redirect(c.Generate().Current())
			}
		}
		results := make([]firewall.Result, len(firewalls))
		secret := c.Request().Header().Get("secret")
		if !session.Super {
			for i, f := range firewalls {
				sessionRoles := make([]auth.Role, 0)
				for _, sr := range session.Roles {
					for _, fr := range f.Roles {
						if slices.ContainsFunc(
							sessionRoles, func(r auth.Role) bool {
								return r.Name == fr.Name
							},
						) {
							continue
						}
						if fr.Name != sr {
							continue
						}
						sessionRoles = append(sessionRoles, fr)
					}
				}
				results[i] = f.Try(
					firewall.Attempt{
						Secret: secret,
						Roles:  sessionRoles,
					},
				)
			}
		}
		allowed := session.Super
		var redirect string
		for _, r := range results {
			if r.Ok {
				allowed = true
				continue
			}
			if len(r.Redirect) > 0 {
				redirect = c.Generate().Link(r.Redirect)
			}
			if r.Err != nil {
				err = r.Err
			}
		}
		redirectExists := len(redirect) > 0
		errorExists := err != nil
		if !allowed && !redirectExists && !errorExists {
			return c.Response().Status(http.StatusForbidden).Error(errors.New(http.StatusText(http.StatusForbidden)))
		}
		if !allowed && redirectExists {
			return c.Response().Redirect(redirect)
		}
		if !allowed && errorExists {
			return c.Response().Error(err)
		}
		if allowed {
			if err := c.Auth().Session().Renew(); err != nil {
				c.Auth().MustOut()
				return c.Response().Redirect(c.Generate().Current())
			}
		}
		return c.Continue()
	}
}
