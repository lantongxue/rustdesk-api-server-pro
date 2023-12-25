package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"rustdesk-api-server-pro/util"
)

type ServerConfig struct {
	Port     string    `yaml:"port"`
	TimeZone string    `yaml:"timeZone"`
	Db       *DbConfig `yaml:"db"`
	SignKey  string    `yaml:"signKey"`
}

type DbConfig struct {
	Driver  string `yaml:"driver"`
	Dsn     string `yaml:"dsn"`
	ShowSql bool   `yaml:"showSql"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: ":8080",
		Db: &DbConfig{
			Driver:  "sqlite",
			Dsn:     "./server.db",
			ShowSql: true,
		},
		SignKey: util.RandomString(32),
	}
}

func GetServerConfig() *ServerConfig {
	cfg := GetDefaultServerConfig()
	bytes, err := os.ReadFile(yamlFile)
	if err != nil {
		WriteServerConfig(cfg)
		return cfg
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		WriteServerConfig(cfg)
		return cfg
	}
	return cfg
}

func WriteServerConfig(cfg *ServerConfig) {
	bytes, _ := yaml.Marshal(cfg)
	_ = os.WriteFile(yamlFile, bytes, 0755)
}
