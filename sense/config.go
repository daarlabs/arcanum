package sense

import (
	"github.com/daarlabs/arcanum/filesystem"
	
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/sense/config"
	
	"github.com/daarlabs/arcanum/mailer"
)

type Config struct {
	App          config.App
	Cache        config.Cache
	Database     map[string]*quirk.DB
	Export       config.Export
	Filesystem   filesystem.Config
	Localization config.Localization
	Parser       config.Parser
	Router       config.Router
	Security     config.Security
	Smtp         mailer.Config
}

const (
	Main = "main"
)
