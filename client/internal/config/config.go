// Пакет config
// Содержит конфиг агента
package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

// Структура Config
// Содержит поля с конфигурацией агента
type Config struct {
	Address     string `env:"ADDRESS"`      // адрес сервера на который будут отправлять данные
	RequestType string `env:"REQUEST TYPE"` // тип реквеста котроый будет отправляться на сервер может быть reg, login, update, delete, add, get
	Entry       string `env:"ENTRY"`        // данные которые отправляются на сервер
	Meta        string `env:"META"`         // метаданные данных которые отправляются на сервер
	Service     string `env:"SERVICE"`      // сервис данные которого отправляются на сервер
	Login       string `env:"LOGIN"`        // логин от системы
	Password    string `env:"PASSWORD"`     // пароль от системы
	Key         string `env:"KEY"`          // ключ шифрования
	JWT         string `env:"JWTFILE"`      // путь до директории где будет храниться JWT-ключ
}

// Функция New создает новый конфиг
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
