// Пакет main клиента
// Используется для запуска клиента
package main

import (
	"gtihub.com/popooq/Gophkeeper/client/internal/agent"
	"gtihub.com/popooq/Gophkeeper/client/internal/config"
	"gtihub.com/popooq/Gophkeeper/client/internal/saver"
	"gtihub.com/popooq/Gophkeeper/client/internal/sender"
)

func main() {
	config := config.New()
	saver := saver.New(config.JWT)
	sender := sender.New(config.Address, saver)
	agent := agent.New(config, sender)

	agent.Run()
}
