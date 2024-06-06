package config

import (
	"flag"
)

type Config struct {
	Host         string `env:"SERVER_ADDRESS"`
	DatabasePath string `env:"DATABASE_DSN"`
}

func NewConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Host, "a", "localhost:8080", "It's a host")
	flag.StringVar(&cfg.DatabasePath, "d", "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable", "connection path to database")

	return cfg
}
