package config

type DbConfig struct {
	Driver string `yaml:"type"`
	Dsn    string `yaml:"dsn"`
}
