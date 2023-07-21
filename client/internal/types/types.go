// пакет types
// содержит в себе структуры которые используются в нескольких пакетах сервиса
// но не относятся ни к одному из них
package types

// данная структура содержит в себе поля с информацией отправляемой на сервер
type Entry struct {
	User     string `json:"user"`     // пользователь
	Service  string `json:"service"`  // сервис
	Entry    string `json:"entry"`    // чувствительная информация
	Metadata string `json:"metadata"` // метаданные
}

// структура содержит в себе поля с информацией о пользователе
type User struct {
	Login    string `json:"login"`    // логин пользователся
	Password string `json:"password"` // хеш пароля пользователя
	ID       int    `json:"id"`       // айди пользователя
}
