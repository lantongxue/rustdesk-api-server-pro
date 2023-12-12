package config

type DbConfig struct {
	Type string `yaml:"type"`
	Dsn  string `yaml:"dsn"`
}
