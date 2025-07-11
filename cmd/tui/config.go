package main

import "github.com/ilyakaznacheev/cleanenv"

type config struct {
	LogLevel string `env:"LOG_LEVEL"`

}

func mustNewConfig() *config {
	config := new(config)

	if err := cleanenv.ReadEnv(config); err != nil {
		panic("cannot read environment variables for config.")
	}

	return config
}
