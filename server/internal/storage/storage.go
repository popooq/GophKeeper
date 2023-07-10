package storage

import "gtihub.com/popooq/Gophkeeper/server/types"

type Keeper interface {
	NewEntry(entry types.Entry) error
	UpdateEntry(entry types.Entry) error
	GetEntry(username, services string) (types.Entry, error)
}
