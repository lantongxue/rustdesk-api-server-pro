package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type ServerConfig struct {
	Port string    `yaml:"port"`
	Db   *DbConfig `yaml:"db"`
}

var DefaultConfig = &ServerConfig{
	Port: ":8080",
	Db: &DbConfig{
		Driver: "sqlite3",
		Dsn:    "./server.db",
	},
}

func GetServerConfig() *ServerConfig {
	cfg := DefaultConfig

	wd, _ := os.Getwd()
	yamlFile := path.Join(wd, "server.yaml")

	bytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return cfg
	}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return cfg
	}
	return cfg
}
