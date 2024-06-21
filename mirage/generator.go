package mirage

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/util"
	
	"github.com/daarlabs/arcanum/csrf"
	"github.com/daarlabs/arcanum/form"
)

type Generator interface {
	Assets() gox.Node
	Action(name string, args ...Map) string
	Current(qpm ...Map) string
	Csrf(name string) gox.Node
	Link(name string, args ...Map) string
	PublicUrl(path string) string
	Query(args Map) string
	SwitchLang(langCode string) string
}

type generator struct {
	*ctx
}

func (g *generator) Assets() gox.Node {
	if g.assets == nil {
		return gox.Fragment()
	}
	googleLinksExist := false
	return gox.Fragment(
		gox.Range(
			g.assets.fonts, func(font string, _ int) gox.Node {
				var preconnects []gox.Node
				isGoogle := strings.Contains(font, "googleapis.com")
				if isGoogle && !googleLinksExist {
					preconnects = append(preconnects, gox.Link(gox.Rel("preconnect"), gox.Href("https://fonts.googleapis.com")))
					preconnects = append(
						preconnects, gox.Link(gox.Rel("preconnect"), gox.Href("https://fonts.gstatic.com"), gox.CrossOrigin()),
					)
					googleLinksExist = true
				}
				return gox.Fragment(
					gox.If(len(preconnects) > 0, gox.Fragment(preconnects...)),
					gox.Link(gox.Rel("stylesheet"), gox.Href(font)),
				)
			},
		),
		gox.Range(
			g.assets.styles, func(style string, _ int) gox.Node {
				return gox.Link(gox.Rel("stylesheet"), gox.Type("text/css"), gox.Href(style))
			},
		),
		gox.Range(
			g.assets.scripts, func(style string, _ int) gox.Node {
				return gox.Script(gox.Defer(), gox.Src(style))
			},
		),
	)
}

func (g *generator) Action(name string, args ...Map) string {
	if g.component == nil {
		return ""
	}
	qpm := Map{Action: createDividedName(g.route.Name, g.component.name, name)}
	parsed := g.Request().Parsed()
	for k, v := range parsed {
		qpm[k] = v
	}
	if len(args) > 0 {
		for k, v := range args[0] {
			vv := reflect.ValueOf(v)
			if vv.IsZero() {
				continue
			}
			qpm[k] = v
		}
	}
	return g.Request().Path() + g.Generate().Query(qpm)
}

func (g *generator) Current(qpm ...Map) string {
	qp := make(Map)
	if len(qpm) > 0 {
		qp = qpm[0]
	}
	for k, v := range g.Request().Raw().URL.Query() {
		if k == Action || k == langQueryKey {
			continue
		}
		qp[k] = v
	}
	return g.proxyPathIfExists(g.ensurePathEndSlash(g.Request().Path())) + g.Query(qp)
}

func (g *generator) Csrf(name string) gox.Node {
	token := g.csrf.MustCreate(
		csrf.Token{
			Name:      name,
			Ip:        g.Request().Ip(),
			UserAgent: g.Request().UserAgent(),
		},
	)
	return form.Csrf(name, token)
}

func (g *generator) Link(name string, args ...Map) string {
	l := g.Lang().Current()
	for _, r := range *g.routes {
		if g.config.Localization.Enabled && !g.config.Localization.Path {
			if r.Name == name {
				return g.generatePath(r.Path, args...)
			}
			continue
		}
		if g.config.Localization.Enabled && r.Name == name && r.Lang == l {
			return g.generatePath(r.Path, args...)
		}
		if r.Name == name {
			return g.generatePath(r.Path, args...)
		}
	}
	return name
}

func (g *generator) Query(args Map) string {
	if len(args) == 0 {
		return ""
	}
	result := make([]string, 0)
	for k, v := range args {
		if v == nil {
			continue
		}
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Slice:
			for i := 0; i < vv.Len(); i++ {
				result = append(result, fmt.Sprintf("%s=%v", k, vv.Index(i).Interface()))
			}
		default:
			result = append(result, fmt.Sprintf("%s=%v", k, v))
		}
	}
	if len(result) == 0 {
		return ""
	}
	return "?" + strings.Join(result, "&")
}

func (g *generator) PublicUrl(path string) string {
	r, err := url.JoinPath("/", g.config.Router.Prefix.Proxy, g.config.App.Public, path)
	if err != nil {
		return path
	}
	return r
}

func (g *generator) SwitchLang(langCode string) string {
	path := g.Request().Path()
	name := g.Request().Name()
	g.cookie.Set(langCookieKey, langCode, langCookieDuration)
	if !g.config.Localization.Path {
		return path
	}
	for _, r := range *g.routes {
		if r.Name == name && r.Lang == langCode {
			return r.Path
		}
	}
	return path
}

func (g *generator) generatePath(path string, args ...Map) string {
	path = g.replacePathParamsWithArgs(path, args...)
	return g.proxyPathIfExists(g.ensurePathEndSlash(path))
}

func (g *generator) ensurePathEndSlash(path string) string {
	return strings.TrimSuffix(path, "/") + "/"
}

func (g *generator) proxyPathIfExists(path string) string {
	if len(g.config.Router.Prefix.Proxy) > 0 {
		return util.MustJoinPath(g.config.Router.Prefix.Proxy, path)
	}
	return path
}

func (g *generator) replacePathParamsWithArgs(path string, args ...Map) string {
	replaceArgs := g.Request().Parsed()
	argsExists := len(args) > 0
	replaceArgsExists := len(replaceArgs) > 0
	if !argsExists && !replaceArgsExists {
		return path
	}
	if argsExists {
		for k, v := range args[0] {
			replaceArgs[k] = v
		}
	}
	replace := make([]string, 0)
	for k, v := range replaceArgs {
		replace = append(replace, "{"+k+"}", fmt.Sprintf("%v", v))
	}
	r := strings.NewReplacer(replace...)
	return r.Replace(path)
}
