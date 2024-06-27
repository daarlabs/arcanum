package mirage

import (
	"fmt"
	"net/url"
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/config"
	"github.com/daarlabs/arcanum/tempest"
)

type assets struct {
	config  config.Config
	code    string
	styles  []string
	scripts []string
	fonts   []string
}

const (
	tempestAssetsPath          = "/.tempest/assets/"
	tempestAssetsCacheDuration = 7 * 24 * time.Hour
)

func createAssets(config config.Config) *assets {
	a := &assets{
		config: config,
		code:   uniuri.New(),
	}
	return a
}

func (a *assets) process() error {
	if err := a.prepareTempestStyles(); err != nil {
		return err
	}
	if err := a.prepareTempestScripts(); err != nil {
		return err
	}
	if err := a.prepareTempestFonts(); err != nil {
		return err
	}
	return nil
}

func (a *assets) mustProcess() {
	if err := a.process(); err != nil {
		panic(err)
	}
}
func (a *assets) prepareTempestStyles() error {
	r, err := url.JoinPath("/", a.config.Router.Prefix.Proxy, tempestAssetsPath, fmt.Sprintf("%s-%s.css", Main, a.code))
	if err != nil {
		return err
	}
	a.styles = append(a.styles, r)
	return nil
}

func (a *assets) prepareTempestScripts() error {
	r, err := url.JoinPath("/", a.config.Router.Prefix.Proxy, tempestAssetsPath, fmt.Sprintf("%s-%s.js", Main, a.code))
	if err != nil {
		return err
	}
	a.scripts = append(a.scripts, r)
	return nil
}

func (a *assets) prepareTempestFonts() error {
	for _, font := range tempest.GlobalConfig.Font {
		a.fonts = append(a.fonts, font.Url)
	}
	return nil
}
