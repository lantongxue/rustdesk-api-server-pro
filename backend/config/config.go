package config

import (
	"os"
	"path"
	"rustdesk-api-server-pro/util"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Db         *DbConfig   `yaml:"db"`
	SignKey    string      `yaml:"signKey"`
	HttpConfig *HttpConfig `yaml:"httpConfig"`
	JobsConfig *JobsConfig `yaml:"jobsConfig"`
}

type DbConfig struct {
	Driver   string `yaml:"driver"`
	Dsn      string `yaml:"dsn"`
	TimeZone string `yaml:"timeZone"`
	ShowSql  bool   `yaml:"showSql"`
}

type HttpConfig struct {
	PrintRequestLog bool   `yaml:"printRequestLog"`
	Port            string `yaml:"port"`
}

type DeviceCheckJob struct {
	Duration int `yaml:"duration"`
}

type JobsConfig struct {
	DeviceCheckJob *DeviceCheckJob `yaml:"deviceCheckJob"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Db: &DbConfig{
			Driver:   "sqlite",
			Dsn:      "./server.db",
			ShowSql:  true,
			TimeZone: "Asia/Shanghai",
		},
		HttpConfig: &HttpConfig{
			Port: ":8080",
		},
		SignKey: util.RandomString(32),
		JobsConfig: &JobsConfig{
			DeviceCheckJob: &DeviceCheckJob{
				Duration: 30,
			},
		},
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
