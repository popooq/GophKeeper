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

type Sender struct {
	endpoint string
	saver    *saver.Saver
}

func New(endpoint string, saver *saver.Saver) Sender {
	return Sender{
		endpoint: endpoint,
		saver:    saver,
	}
}

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

func (s *Sender) Add(user, service, entry, metadata string) {
	data := entryBuild(user, service, entry, metadata)

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
func (s *Sender) Update(user, service, entry, metadata string) {
	data := entryBuild(user, service, entry, metadata)

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

func entryBuild(user, service, entry, metadata string) []byte {
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
