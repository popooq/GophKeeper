// пакет содержит в себе интерфейс структуры обращения к БД
package storage

import (
	"gtihub.com/popooq/Gophkeeper/server/internal/types"
)

// интерфейс структуры обращения к БД
type Keeper interface {
	Registration(username, password string) (types.User, error)
	Login(username, password string) bool
	NewEntry(entry types.Entry) error
	UpdateEntry(entry types.Entry) error
	GetEntry(username, services string) (types.Entry, error)
	DeleteEntry(username, service string) (int, error)
}
