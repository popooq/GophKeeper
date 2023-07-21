// пакет Sender отправляет данные на сервер
package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/url"

	"github.com/go-resty/resty/v2"
	"gtihub.com/popooq/Gophkeeper/client/internal/saver"
	"gtihub.com/popooq/Gophkeeper/client/internal/types"
)

// структура Sender
// состоит из эндпоинта - адреса сервера
// и из saver - структуры для сохранения jwt-ключа
type Sender struct {
	endpoint string
	saver    *saver.Saver
}

// функция New создает новый Sender
func New(endpoint string, saver *saver.Saver) Sender {
	return Sender{
		endpoint: endpoint,
		saver:    saver,
	}
}

// Регистрация нового пользователя
func (s *Sender) Reg(login, password string) {
	user := types.User{
		Login:    login,
		Password: password,
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := bytes.NewBuffer(body)

	endpoint, err := url.JoinPath("http://", s.endpoint, "server/registration")
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	req := client.R().
		SetHeader("Content-Type", "application/json")

	resp, err := req.SetBody(reqBody).Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		defer resp.RawBody().Close()
	}
}

// Авторизация пользователя
func (s *Sender) Login(login, password string) {
	user := types.User{
		Login:    login,
		Password: password,
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	reqBody := bytes.NewBuffer(body)

	endpoint, err := url.JoinPath("http://", s.endpoint, "server/login")
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	req := client.R().
		SetHeader("Content-Type", "application/json")

	resp, err := req.SetBody(reqBody).Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		err := s.saver.SaveJWT(resp.Body())
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.RawBody().Close()
	}
}

// добавление новой информации
func (s *Sender) Add(user, service, entry, metadata string) {
	data := s.entryBuild(user, service, entry, metadata)

	body := bytes.NewBuffer(data)

	endpoint, err := url.JoinPath("http://", s.endpoint, "entry/new")
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	jwt, err := s.saver.LoadJWT()

	req := client.R().
		SetAuthToken(string(jwt))

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := req.SetBody(body).Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		defer resp.RawBody().Close()
	}
}

// изменение существующей информации
func (s *Sender) Update(user, service, entry, metadata string) {
	data := s.entryBuild(user, service, entry, metadata)

	body := bytes.NewBuffer(data)

	endpoint, err := url.JoinPath("http://", s.endpoint, "entry/update")
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	jwt, err := s.saver.LoadJWT()

	req := client.R().
		SetAuthToken(string(jwt))

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := req.SetBody(body).Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		defer resp.RawBody().Close()
	}
}

// получение информации по паре логин/сервис
func (s *Sender) Get(user, service string) {
	endpoint, err := url.JoinPath("http://", s.endpoint, "entry/get", user, service)
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	jwt, err := s.saver.LoadJWT()

	req := client.R().
		SetAuthToken(string(jwt))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := req.Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		defer resp.RawBody().Close()
	}
}

// удаление информации по паре логин/сервис
func (s *Sender) Delete(user, service string) {
	endpoint, err := url.JoinPath("http://", s.endpoint, "entry/delete", user, service)
	if err != nil {
		log.Printf("url joining failed, error: %s", err)
	}

	client := resty.New().SetBaseURL(endpoint)

	jwt, err := s.saver.LoadJWT()

	req := client.R().
		SetAuthToken(string(jwt))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := req.Post(endpoint)
	if err != nil {
		log.Printf("Server unreachible, error: %s", err)
	} else {
		defer resp.RawBody().Close()
	}
}

func (s *Sender) entryBuild(user, service, entry, metadata string) []byte {
	data := types.Entry{
		User:     user,
		Service:  service,
		Entry:    entry,
		Metadata: metadata,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}
	return body
}
