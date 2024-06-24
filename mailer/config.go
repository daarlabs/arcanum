package mailer

type Config struct {
	Host     string `json:"host" toml:"host" yaml:"host"`
	Port     int    `json:"port" toml:"port" yaml:"port"`
	User     string `json:"user" toml:"user" yaml:"user"`
	Password string `json:"password" toml:"password" yaml:"password"`
	From     string `json:"from" toml:"from" yaml:"from"`
}
