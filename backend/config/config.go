package config

import (
	"os"
	"path"
	"rustdesk-api-server-pro/util"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	TimeZone   string      `yaml:"timeZone"`
	Db         *DbConfig   `yaml:"db"`
	SignKey    string      `yaml:"signKey"`
	HttpConfig *HttpConfig `yaml:"httpConfig"`
}

type DbConfig struct {
	Driver  string `yaml:"driver"`
	Dsn     string `yaml:"dsn"`
	ShowSql bool   `yaml:"showSql"`
}

type HttpConfig struct {
	PrintRequestLog bool   `yaml:"printRequestLog"`
	Port            string `yaml:"port"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Db: &DbConfig{
			Driver:  "sqlite",
			Dsn:     "./server.db",
			ShowSql: true,
		},
		HttpConfig: &HttpConfig{
			Port: ":8080",
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
