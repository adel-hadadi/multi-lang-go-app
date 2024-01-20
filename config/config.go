package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App
}

type App struct {
	Port  int    `env:"APP_PORT"`
	Mode  string `env:"APP_MODE"`
	Local string `env:"APP_LOCAL"`
}

var cfg *Config

func NewConfig() *Config {
	if cfg == nil {
		readConfig := Config{}
		err := cleanenv.ReadConfig(".env", &readConfig)
		if err != nil {
			panic(err)
		}

		cfg = &readConfig
	}

	return cfg
}
