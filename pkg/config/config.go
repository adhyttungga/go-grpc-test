package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type DB struct {
	Name     string `env:"DB_NAME" envDefault:"db_grpc"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type Redis struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:8081"`
	Password string `env:"REDIS_PASSWORD"`
}

type ServerConfig struct {
	Port       string `env:"PORT" envDefault:"8080"`
	DB         DB
	Redis      Redis
	PrivateKey string `env:"PRIVATE_KEY"`
	PublicKey  string `env:"PUBLIC_KEY"`
}

var Config ServerConfig

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	if err := env.Parse(&Config); err != nil {
		log.Fatalf("error parse env: %v", err)
	}
}
