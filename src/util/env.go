package util

import (
	e "github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload" // autoload environment variables from .env
)

type Environment struct {
	HOST string `env:"HOST"`
	PORT string `env:"PORT"`
}

func LoadEnv() (Environment, error) {
	return e.ParseAs[Environment]()
}
