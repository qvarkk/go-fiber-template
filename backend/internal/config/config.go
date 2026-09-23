package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
)

type Config struct {
	Env      AppEnv   `envconfig:"APP_ENV" default:"development" validate:"required,validEnv"`
	LogLevel LogLevel `envconfig:"LOG_LEVEL" default:"info" validate:"required,validLogLevel"`
}

type ConfigOut struct {
	fx.Out

	Env      AppEnv
	LogLevel LogLevel
}

func Load() (ConfigOut, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return ConfigOut{}, fmt.Errorf("failed to parse configuration: %w", err)
	}

	validate, err := newValidator()
	if err != nil {
		return ConfigOut{}, fmt.Errorf("failed to spin up validation framework: %w", err)
	}

	if err := validate.Struct(&cfg); err != nil {
		return ConfigOut{}, fmt.Errorf("configuration verification engine caught critical error: %w", err)
	}

	return ConfigOut{
		Env:      cfg.Env,
		LogLevel: cfg.LogLevel,
	}, nil
}
