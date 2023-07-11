package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Address     string `env:"ADDRESS"`
	RequestType string `env:"REQUEST TYPE"`
	Entry       string `env:"ENTRY"`
	Meta        string `env:"META"`
	Service     string `env:"SERVICE"`
	Login       string `env:"LOGIN"`
	Password    string `env:"PASSWORD"`
	Key         string `env:"KEY"`
	JWT         string `env:"JWTFILE"`
}

func New() *Config {
	var cfg Config
	flag.StringVar(&cfg.Address, "a", "127.0.0.1:8080", "set server address")
	flag.StringVar(&cfg.RequestType, "r", "", "set request type reg/login/add/get/delete/update")
	flag.StringVar(&cfg.Login, "l", "", "your login to system")
	flag.StringVar(&cfg.Password, "p", "", "your password to system")
	flag.StringVar(&cfg.Entry, "d", "", "set data u want to save")
	flag.StringVar(&cfg.Meta, "m", "", "metadate for your entry")
	flag.StringVar(&cfg.Service, "s", "", "service of your entry")
	flag.StringVar(&cfg.Key, "k", "", "encryption key")
	flag.StringVar(&cfg.JWT, "j", "", "path to the file to save jwt to")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Printf("env parse failed :%s", err)
	}

	return &cfg
}
