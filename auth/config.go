package auth

import "time"

type Config struct {
	Roles    []Role        `json:"roles" yaml:"roles" toml:"roles"`
	Duration time.Duration `json:"duration" yaml:"duration" toml:"duration"`
}
