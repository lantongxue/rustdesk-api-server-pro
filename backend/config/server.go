package config

type ServerConfig struct {
	Port string   `yaml:"port"`
	Db   DbConfig `yaml:"db"`
}
