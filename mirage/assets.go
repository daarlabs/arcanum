package mirage

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/config"
)

type assets struct {
	config  config.Config
	code    string
	styles  []string
	scripts []string
	fonts   []string
}

type viteManifest struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	Css     []string `json:"css"`
}

const (
	manifestFilaname  = "manifest.json"
	distDir           = "dist"
	tempestAssetsPath = "/.tempest/assets/"
)

func createAssets(config config.Config) *assets {
	a := &assets{
		config: config,
		code:   uniuri.New(),
	}
	return a
}

func (a *assets) process() error {
	if err := a.read(); err != nil {
		return err
	}
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

func (a *assets) read() error {
	filePath, err := url.JoinPath(a.config.App.Assets, distDir, manifestFilaname)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	assetsMap := make(map[string]viteManifest)
	if err := json.Unmarshal(bytes, &assetsMap); err != nil {
		return err
	}
	for _, item := range assetsMap {
		var err error
		item.File, err = url.JoinPath(a.config.Router.Prefix.Proxy, a.config.App.Public, distDir, item.File)
		if err != nil {
			continue
		}
		a.scripts = append(a.scripts, item.File)
		for _, css := range item.Css {
			var err error
			css, err = url.JoinPath(a.config.Router.Prefix.Proxy, a.config.App.Public, distDir, css)
			if err != nil {
				continue
			}
			a.styles = append(a.styles, css)
		}
	}
	return nil
}

func (a *assets) prepareTempestStyles() error {
	r, err := url.JoinPath(a.config.Router.Prefix.Proxy, tempestAssetsPath, fmt.Sprintf("%s-%s.css", Main, a.code))
	if err != nil {
		return err
	}
	a.styles = append(a.styles, r)
	return nil
}

func (a *assets) prepareTempestScripts() error {
	r, err := url.JoinPath(a.config.Router.Prefix.Proxy, tempestAssetsPath, fmt.Sprintf("%s-%s.js", Main, a.code))
	if err != nil {
		return err
	}
	a.scripts = append(a.scripts, r)
	return nil
}

func (a *assets) prepareTempestFonts() error {
	for _, font := range a.config.Tempest.Fonts() {
		a.fonts = append(a.fonts, font.Url)
	}
	return nil
}
