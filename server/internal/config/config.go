package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address         string `env:"ADDRESS"`
	DatabaseAddress string `env:"DBA"`
	Secret          string `env:"SECRET"`
}

func New() *Config {
	var cfg Config
	flag.StringVar(&cfg.Address, "a", "", "Server Address")
	flag.StringVar(&cfg.DatabaseAddress, "db", "postgresql://leonidagupov@localhost:5432/gophkeeper", "database Address")
	flag.StringVar(&cfg.Secret, "s", "secretesse", "secret for jwt")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Printf("env parse failed :%s", err)
	}

	return &cfg
}
