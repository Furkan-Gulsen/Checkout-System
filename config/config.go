package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	ServerConfig struct {
		Host string
		Port string
	}

	MongoDBConfig struct {
		Host     string
		Port     string
		Database string
	}

	Config struct {
		Server ServerConfig
		Mongo  MongoDBConfig
	}
)

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "docker" {
		viper.SetConfigName("docker-config")
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}

	return &cfg, nil
}
