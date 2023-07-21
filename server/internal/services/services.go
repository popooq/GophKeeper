// пакет services
// содержит в себе сервисы обработчиков системы
package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"gtihub.com/popooq/Gophkeeper/server/internal/storage"
	"gtihub.com/popooq/Gophkeeper/server/internal/types"
)

func generateJWT(user string, secret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user

	log.Println(secret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("error in signing %s", err)
		return "", err
	}

	return tokenString, nil
}

// реализация регистрации пользователя
// принимает на вход data io.Reader - тело запроса
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func Registration(data io.Reader, keeper storage.Keeper) (int, error) {
	var user types.User

	body, err := io.ReadAll(data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, err = keeper.Registration(user.Login, user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// реализация авторизации пользователя
// принимает на вход data io.Reader - тело запроса,
// secret []byte - jwt-ключ
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func Login(data io.Reader, secret []byte, keeper storage.Keeper) (string, int, error) {
	var user types.User

	body, err := io.ReadAll(data)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	ok := keeper.Login(user.Login, user.Password)
	if !ok {
		return "", http.StatusUnauthorized, fmt.Errorf("")
	}

	jwt, err := generateJWT(user.Login, secret)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return jwt, http.StatusOK, nil
}

// реализация добавления инфромации в сервис
// принимает на вход data io.Reader - тело запроса
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func NewEntry(data io.Reader, keeper storage.Keeper) (int, error) {
	var entry types.Entry

	body, err := io.ReadAll(data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = json.Unmarshal(body, &entry)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = keeper.NewEntry(entry)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// реализация обновления инфромации в сервисе
// принимает на вход data io.Reader - тело запроса
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func UpdateEntry(data io.Reader, keeper storage.Keeper) ([]byte, int) {
	var entry types.Entry

	body, err := io.ReadAll(data)
	if err != nil {
		return []byte(fmt.Sprint(err)), http.StatusInternalServerError
	}

	err = json.Unmarshal(body, &entry)
	if err != nil {
		return []byte(fmt.Sprint(err)), http.StatusInternalServerError
	}

	err = keeper.UpdateEntry(entry)
	if err != nil {
		return []byte(fmt.Sprint(err)), http.StatusInternalServerError
	}

	return []byte("entry updatet"), http.StatusOK
}

// реализация получения инфромации из сервиса
// принимает на вход username string - логин пользователя,
// service string - имя сервиса
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func GetEntry(username string, service string, keeper storage.Keeper) ([]byte, int) {

	entry, err := keeper.GetEntry(username, service)
	if err != nil {
		log.Println("error during getting entry. pckg services", err)
		return nil, http.StatusInternalServerError
	}

	i, err := json.Marshal(entry)
	if err != nil {
		log.Println("error during marshalling. pckg services", err)
		return nil, http.StatusInternalServerError
	}

	return i, http.StatusOK
}

// реализация удаления инфромации из сервиса
// принимает на вход username string - логин пользователя,
// service string - имя сервиса
// и keeper storage.Keeper - интерфейс реализующий запрос к БД
func DeleteEntry(username string, service string, keeper storage.Keeper) int {

	status, err := keeper.DeleteEntry(username, service)
	if err != nil {
		return http.StatusInternalServerError
	}

	return status
}
