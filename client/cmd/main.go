package main

import (
	"gtihub.com/popooq/Gophkeeper/client/internal/agent"
	"gtihub.com/popooq/Gophkeeper/client/internal/config"
)

func main() {
	config := config.New()
	agent := agent.New(config.Address)
	agent.Run()
}
