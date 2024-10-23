package config

import (
	"os"
	"sync/atomic"

	"github.com/caarlos0/env/v6"
	"github.com/thmelodev/ddd-events-api/src/utils/logger"
)

type Option func(*env.Options)

const (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
	Test        = "test"
)

var (
	_global atomic.Value
)

func Get() Config {
	return _global.Load().(Config)
}

// Init will start all the environments variable into the Config struct

func Init() *Config {
	cfg := Config{}
	load("ddd-events-api", &cfg)
	_global.Store(cfg)
	return &cfg
}

// Load will map all the envs to the reference struct
func load(appName string, reference *Config) {
	sLog := logger.Get()

	if len(appName) > 0 {
		os.Setenv("APP_NAME", appName)
	}

	if err := env.Parse(reference); err != nil {
		sLog.Errorf("error parsing configs - %+v", err)
	}

}
