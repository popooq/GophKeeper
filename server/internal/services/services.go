package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gtihub.com/popooq/Gophkeeper/server/internal/storage"
	"gtihub.com/popooq/Gophkeeper/server/types"
)

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

func GetEntry(username string, service string, keeper storage.Keeper) ([]byte, int) {

	entry, err := keeper.GetEntry(username, service)
	if err != nil {
		log.Println("error during getting entry. pckg services", err)
		return nil, http.StatusInternalServerError
	}

	i, err := json.Marshal(entry)
	if err != nil {
		log.Println("error during marshalling. pckg setvices", err)
		return nil, http.StatusInternalServerError
	}

	return i, http.StatusOK
}
