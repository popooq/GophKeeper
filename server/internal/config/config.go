// Пакет config
// Содержит конфиг сервера
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

// Структура Config
// Содержит поля с конфигурацией сервера
type Config struct {
	Address         string `env:"ADDRESS"` // адрес сервера
	DatabaseAddress string `env:"DBA"`     // адрес базы данных
	Secret          string `env:"SECRET"`  // ключ для создания jwt
}

// Функция New создает новый конфиг
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
