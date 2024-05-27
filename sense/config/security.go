package config

import (
	"time"
	
	"github.com/daarlabs/arcanum/auth"
)

type Security struct {
	Auth      auth.Config
	Csrf      Csrf
	Firewalls []Firewall
}

type Csrf struct {
	Expiration time.Duration
}

type Firewall struct {
	Enabled  bool
	Patterns []string
	Roles    []string
}
