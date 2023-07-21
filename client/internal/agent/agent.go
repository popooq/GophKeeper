// Пакет Agent клиента
// агент распределяет запросы по сендеру
package agent

import (
	"gtihub.com/popooq/Gophkeeper/client/internal/config"
	"gtihub.com/popooq/Gophkeeper/client/internal/sender"
)

// структура Agent необходимая для работы агента
// состоит из конфига и сендера
type Agent struct {
	config *config.Config
	sender sender.Sender
}

// Функция New создает новый агент
func New(config *config.Config, sender sender.Sender) *Agent {
	return &Agent{
		config: config,
		sender: sender,
	}
}

// Функция Run запускает агент в работу
func (a *Agent) Run() {
	switch a.config.RequestType {
	case "reg":
		a.sender.Reg(a.config.Login, a.config.Password)
	case "login":
		a.sender.Login(a.config.Login, a.config.Password)
	case "add":
		a.sender.Add(a.config.Login, a.config.Service, a.config.Entry, a.config.Meta)
	case "get":
		a.sender.Get(a.config.Login, a.config.Service)
	case "delete":
		a.sender.Delete(a.config.Login, a.config.Service)
	case "update":
		a.sender.Update(a.config.Login, a.config.Service, a.config.Entry, a.config.Meta)
	}
}
