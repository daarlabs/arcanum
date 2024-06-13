package config

import (
	"github.com/daarlabs/arcanum/filesystem"
	"github.com/daarlabs/arcanum/form"
	
	"github.com/daarlabs/arcanum/mailer"
	"github.com/daarlabs/arcanum/quirk"
)

type Config struct {
	App          App
	Cache        Cache
	Database     map[string]*quirk.DB
	Export       Export
	Form         form.Config
	Filesystem   filesystem.Config
	Localization Localization
	Parser       Parser
	Router       Router
	Security     Security
	Smtp         mailer.Config
}

func (c Config) Init() Config {
	if c.Form.Limit == 0 {
		c.Form.Limit = form.DefaultBodyLimit
	}
	return c
}
