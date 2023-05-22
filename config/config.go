package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	AppPort  string `env:"APP_PORT" envDefault:"4000"`
	AppMode  string `env:"APP_MODE" envDefault:"Dev"`
	DBDriver string `env:"DB_DRIVER" envDefault:"pgx"`
	DBConURL string `env:"DB_CON_URL"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("couldn't parse config: %w", err)
	}

	return &cfg, nil
}

func PrepareENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}
