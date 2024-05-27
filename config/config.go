package config

import (
	"github.com/daarlabs/arcanum/filesystem"
	
	"github.com/daarlabs/arcanum/mailer"
	"github.com/daarlabs/arcanum/quirk"
)

type Config struct {
	App          App
	Cache        Cache
	Database     map[string]*quirk.DB
	Export       Export
	Filesystem   filesystem.Config
	Localization Localization
	Parser       Parser
	Router       Router
	Security     Security
	Smtp         mailer.Config
}
