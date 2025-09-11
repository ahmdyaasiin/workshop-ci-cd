package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type App struct {
	AppName string `env:"APP_NAME" envDefault:"Go App"`
	AppPort int    `env:"APP_PORT" envDefault:"8001"`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
}

type DB struct {
	DatabaseUsername string `env:"DATABASE_USERNAME,required"`
	DatabasePassword string `env:"DATABASE_PASSWORD,required"`
	DatabaseHost     string `env:"DATABASE_HOST,required"`
	DatabasePort     int    `env:"DATABASE_PORT,required"`
	DatabaseDB       string `env:"DATABASE_DB,required"`
	DatabaseSSL      string `env:"DATABASE_SSL,required"`
}

type Env struct {
	App
	DB
}

func Load() (*Env, error) {
	_env := new(Env)

	if err := env.Parse(_env); err != nil {
		return nil, err
	}

	return _env, nil
}
