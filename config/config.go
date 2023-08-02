package config

import (
	"errors"
	"fmt"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Host       string `env:"SERVER_HOST,default=localhost"`
	Port       string `env:"SERVER_PORT"`
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`
	AppMode    string `env:"APP_MODE,default=dev"`
	MySQL      MySQL
}

type MySQL struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     string `env:"DB_PORT,default=3306"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		message := fmt.Sprintf("error load %s file", env)
		return nil, errors.New(message)
	}
	return &config, nil
}
