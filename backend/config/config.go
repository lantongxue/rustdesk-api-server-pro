package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"rustdesk-api-server-pro/util"
)

type ServerConfig struct {
	Port string     `yaml:"port"`
	Db   *DbConfig  `yaml:"db"`
	JWT  *JWTConfig `yaml:"jwt"`
	//RSA  *RSAConfig `yaml:"rsa"`
}

type DbConfig struct {
	Driver  string `yaml:"driver"`
	Dsn     string `yaml:"dsn"`
	ShowSql bool   `yaml:"showSql"`
}

type RSAConfig struct {
	Private string `yaml:"private"`
	Public  string `yaml:"public"`
}

type JWTConfig struct {
	SignKey string `yaml:"signKey"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	//var key, pub = util.GenerateRSAKeys()
	return &ServerConfig{
		Port: ":8080",
		Db: &DbConfig{
			Driver:  "sqlite",
			Dsn:     "./server.db",
			ShowSql: true,
		},
		JWT: &JWTConfig{
			SignKey: util.RandomString(32),
		},
		//RSA: &RSAConfig{
		//	Private: string(key),
		//	Public:  string(pub),
		//},
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
	//if cfg.RSA == nil || cfg.RSA.Private == "" || cfg.RSA.Public == "" {
	//	cfg.RSA = GetDefaultServerConfig().RSA
	//	WriteServerConfig(cfg)
	//}
	return cfg
}

func WriteServerConfig(cfg *ServerConfig) {
	bytes, _ := yaml.Marshal(cfg)
	_ = os.WriteFile(yamlFile, bytes, 0755)
}
