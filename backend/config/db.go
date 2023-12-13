package config

type DbConfig struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}
