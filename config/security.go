package config

import (
	"github.com/daarlabs/arcanum/auth"
	"github.com/daarlabs/arcanum/csrf"
	"github.com/daarlabs/arcanum/firewall"
)

type Security struct {
	Auth      auth.Config
	Csrf      csrf.Csrf
	Firewalls []firewall.Firewall
}
