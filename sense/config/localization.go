package config

import (
	"github.com/daarlabs/arcanum/translator"
	"github.com/daarlabs/arcanum/validator"
)

type Localization struct {
	Enabled    bool
	Languages  []Language
	Translator translator.Translator
	Validator  validator.Messages
}

type Language struct {
	Main bool
	Code string
}
