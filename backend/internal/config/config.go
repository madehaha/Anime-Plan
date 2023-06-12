package config

import (
	"backend/internal/logger"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	HttpHost  string `yaml:"http_host"`
	HttpPort  string `yaml:"http_port"`
	WebDomain string `yaml:"web_domain"`
	RedisUrl  string `yaml:"redis_url"`

	Mysql struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		user string `yaml:"user"`
		db   string `yaml:"db"`
	}
}

func AppConfigReader() (config AppConfig, err error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		logger.Error("Failed to read config.yaml")
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		// TODO log
		logger.Error("Failed to parse config")
		return
	}
	return
}

func (c *AppConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%s", c.HttpHost, c.HttpPort)
}
