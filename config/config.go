package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Database database
}

type database struct {
	Host     string `env:"DB_HOST" env-default:"db"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func NewConfig() (*config, error) {
	var cfg config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
