package tempest

import (
	"embed"
	"fmt"
	"strings"
	"sync"
)

type Tempest struct {
	mu                      *sync.RWMutex
	config                  Config
	externalResourceManager *externalResourceManager
	classes                 map[string]string
	baseStyles              string
	stylesBundle            string
	externalStylesBundle    string
	scriptsBundle           string
}

const (
	baseCssFilename = "base.css"
)

//go:embed base.css
var baseCss embed.FS

func New(config Config) *Tempest {
	t := &Tempest{
		mu:                      new(sync.RWMutex),
		config:                  config,
		externalResourceManager: createExternalResourceManager(config),
		classes:                 make(map[string]string),
	}
	if config.FontSize == 0 {
		t.config.FontSize = DefaultFontSize
	}
	if config.Breakpoint == nil {
		t.config.Breakpoint = DefaultBreakpoints
	}
	if config.Container == nil {
		t.config.Container = DefaultContainer
	}
	t.config.Color = mergeConfigMap[Color](Pallete, config.Color)
	t.config.Shadow = mergeConfigMap[[]Shadow](BoxShadow, config.Shadow)
	t.config = t.config.processShadows()
	t.onInit()
	return t
}

func (t *Tempest) Context() *Context {
	return &Context{
		Tempest:  t,
		builders: make([]*Builder, 0),
		classes:  make(map[string]string),
	}
}

func (t *Tempest) Fonts() map[string]Font {
	return t.config.Font
}

func (t *Tempest) Styles() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.stylesBundle
}

func (t *Tempest) Scripts() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.scriptsBundle
}

func (t *Tempest) Build() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.buildStyles()
}

func (t *Tempest) onInit() {
	t.buildBaseStyles()
	t.externalResourceManager.mustRun()
	t.buildExternalStyles()
	t.buildScripts()
}

func (t *Tempest) buildBaseStyles() {
	t.mustReadBaseStylesFile()
	t.processBaseStyles()
}

func (t *Tempest) readBaseStylesFile() error {
	baseStylesBytes, err := baseCss.ReadFile(baseCssFilename)
	if err != nil {
		return err
	}
	t.baseStyles = string(baseStylesBytes)
	return nil
}

func (t *Tempest) mustReadBaseStylesFile() {
	if err := t.readBaseStylesFile(); err != nil {
		panic(err)
	}
}

func (t *Tempest) processBaseStyles() {
	replacer := strings.NewReplacer(
		" ", "",
		"\n", " ",
		"\t", "",
		"\r", "",
		baseFontSize, stringifyMostSuitableNumericType(t.config.FontSize)+Px,
		baseFontFamily, fmt.Sprintf("%s", t.config.FontFamily),
	)
	t.baseStyles = replacer.Replace(t.baseStyles)
	
}

func (t *Tempest) buildStyles() {
	result := make([]string, len(t.classes))
	i := 0
	for _, v := range t.classes {
		result[i] = v
		i++
	}
	r := strings.Join(result, " ")
	t.stylesBundle = t.externalStylesBundle + "\n" + t.baseStyles + "\n" + r
}

func (t *Tempest) buildExternalStyles() {
	w := new(strings.Builder)
	for _, item := range t.externalResourceManager.styles {
		w.Write(item)
		w.WriteString("\n")
	}
	t.stylesBundle = w.String()
}

func (t *Tempest) buildScripts() {
	w := new(strings.Builder)
	for _, item := range t.externalResourceManager.scripts {
		w.Write(item)
		w.WriteString("\n")
	}
	t.scriptsBundle = w.String()
}

func (t *Tempest) classExists(name string) bool {
	_, ok := t.classes[name]
	return ok
}
