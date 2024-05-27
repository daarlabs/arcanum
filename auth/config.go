package auth

import "time"

type Config struct {
	Roles    map[string]Role `json:"roles" yaml:"roles" toml:"roles"`
	Duration time.Duration   `json:"duration" yaml:"duration" toml:"duration"`
}

type Role struct {
	Super bool `json:"super" yaml:"super" toml:"super"`
}
