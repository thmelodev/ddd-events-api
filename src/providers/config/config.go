package config

import "time"

type HttpConfig struct {
	HttpPort int `env:"PORT"`
}

type AppConfig struct {
	AppLogLevel string `env:"LOG_LEVEL"`
}

type Config struct {
	Http HttpConfig `envPrefix:"HTTP_"`
	App  AppConfig  `envPrefix:"APP_"`
	Db   DbConfig   `envPrefix:"DB_"`
}

type DbConfig struct {
	Driver                string        `env:"DRIVER"`
	Host                  string        `env:"HOST"`
	Port                  string        `env:"PORT"`
	Name                  string        `env:"NAME"`
	User                  string        `env:"USER"`
	Pass                  string        `env:"PASSWORD"`
	ConnectionRetries     int           `env:"CONNECTION_RETRIES"`
	MaxIdleConnections    int           `env:"MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections    int           `env:"MAX_OPEN_CONNECTIONS"`
	ConnectionMaxLifetime time.Duration `env:"CONNECTION_MAX_LIFETIME"`
}
