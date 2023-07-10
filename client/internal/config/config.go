package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address  string `env:"ADDRESS"`
	Data     string `env:"DATA"`
	Login    string `env:"LOGIN"`
	Password string `env:"PASSWORD"`
	Key      string `env:"KEY"`
}

func New() *Config {
	var cfg Config
	flag.StringVar(&cfg.Address, "a", "127.0.0.1:8080", "set server address")
	flag.StringVar(&cfg.Login, "l", "", "your login to system")
	flag.StringVar(&cfg.Password, "p", "", "your password to system")
	flag.StringVar(&cfg.Data, "d", "", "set data u want to save")
	flag.StringVar(&cfg.Key, "k", "", "encryption key")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Printf("env parse failed :%s", err)
	}

	return &cfg
}
