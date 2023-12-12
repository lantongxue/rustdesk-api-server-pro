package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"rustdesk-api-server-pro/config"
)

var DefaultConfig = &config.ServerConfig{
	Port: ":8080",
	Db: config.DbConfig{
		Type: "sqlite3",
		Dsn:  "./server.db",
	},
}

func GetServerConfig() *config.ServerConfig {
	cfg := DefaultConfig
	wd, _ := os.Getwd()
	cfgFile := path.Join(wd, "server.yaml")

	bytes, err := os.ReadFile(cfgFile)
	if err != nil {
		return cfg
	}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return cfg
	}
	return cfg
}
