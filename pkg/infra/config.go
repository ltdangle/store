package infra

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	MYSQL_DSN    string `env:"MYSQL_DSN"`
	POSTGRES_URL string `env:"POSTGRES_URL"`
}

func ReadConfig(envFile string) Config {

	// Load .env config.
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}
