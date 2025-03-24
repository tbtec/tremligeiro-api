package env

import (
	"context"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Env              string `env:"ENV" envDefault:"local"`
	Port             int    `env:"PORT" envDefault:"8080"`
	DbHost           string `env:"POSTGRES_HOST"`
	DbUser           string `env:"POSTGRES_USER"`
	DbPassword       string `env:"POSTGRES_PASS"`
	DbName           string `env:"POSTGRES_DB"`
	DbPort           int    `env:"POSTGRES_PORT"`
	PaymentUrl       string `env:"PAYMENT_URL"`
	PaymentAuthToken string `env:"PAYMENT_AUTH_TOKEN"`
}

func LoadEnvConfig() (Config, error) {
	cfg := Config{}
	var err error

	if fileExists("../../.env") {
		err = godotenv.Load("../../.env")
	}
	if fileExists("../.env") {
		err = godotenv.Load("../.env")
	}
	if fileExists(".env") {
		err = godotenv.Load(".env")
	}
	if err != nil {
		slog.ErrorContext(context.Background(), "❌ Could not load .env file")
		panic(err)
	}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}
	slog.InfoContext(context.Background(), "✅ Success on read .env configuration...")

	return cfg, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return info != nil
}
