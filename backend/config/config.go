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

type DbConfig struct {
	Driver  string `yaml:"driver"`
	Dsn     string `yaml:"dsn"`
	ShowSql bool   `yaml:"showSql"`
}

var DefaultConfig = &ServerConfig{
	Port: ":8080",
	Db: &DbConfig{
		Driver:  "sqlite",
		Dsn:     "./server.db",
		ShowSql: true,
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
