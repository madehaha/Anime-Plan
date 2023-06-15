package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"backend/internal/logger"
)

type AppConfig struct {
	HttpHost  string `yaml:"http_host"`
	HttpPort  string `yaml:"http_port"`
	WebDomain string `yaml:"web_domain"`
	RedisUrl  string `yaml:"redis_url"`

	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	}

	JwtSecret       string `yaml:"jwt"`
	StaticDirectory string `yaml:"static_directory"`
}

func AppConfigReader() AppConfig {
	var config AppConfig
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		logger.Fatal("Failed to read config.yaml")
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Fatal("Failed to parse config")
	}
	return config
}

func (c *AppConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%s", c.HttpHost, c.HttpPort)
}
