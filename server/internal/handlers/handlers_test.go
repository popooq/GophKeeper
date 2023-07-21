package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi/v5"
	"gtihub.com/popooq/Gophkeeper/server/internal/storage"
	"gtihub.com/popooq/Gophkeeper/server/internal/types"
)

type keeperMock struct {
}

func (k keeperMock) Registration(username, password string) (types.User, error) {
	return types.User{}, nil
}

func (k keeperMock) Login(username, password string) bool {
	return true
}

func (k keeperMock) NewEntry(entry types.Entry) error {
	return nil
}

func (k keeperMock) UpdateEntry(entry types.Entry) error {
	return nil
}

func (k keeperMock) GetEntry(username, services string) (types.Entry, error) {
	return types.Entry{}, nil
}

func (k keeperMock) DeleteEntry(username, service string) (int, error) {
	return 1, nil
}

func NewRouter() *chi.Mux {
	var keeper keeperMock
	handler := New(keeper, "")

	r := chi.NewRouter()
	r.Mount("/", handler.Route())

	return r
}

func TestNew(t *testing.T) {
	type args struct {
		keeper storage.Keeper
		secret string
	}
	tests := []struct {
		name string
		args args
		want *Handlers
	}{
		// TODO: Add test cases.
		{
			name: "positive create",
			args: args{
				keeper: keeperMock{},
				secret: "",
			},
			want: &Handlers{
				keeper: keeperMock{},
				secret: []byte(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.keeper, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_Login(t *testing.T) {
	url := "/server/login"
	tests := []struct {
		name string
		body string
		code int
	}{
		// TODO: Add test cases.
		{name: "Positive Reg test",
			body: `{"login": "login", "password": "password"}`,
			code: 200,
		},
		{
			name: "Negative Reg test",
			body: `{"login": 2341, "password": "password"}`,
			code: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody := bytes.NewBuffer([]byte(tt.body))
			r := httptest.NewRequest(http.MethodPost, url, requestBody)
			w := httptest.NewRecorder()

			h := NewRouter()
			h.ServeHTTP(w, r)
			result := w.Result()
			if result.StatusCode != tt.code {
				t.Errorf("Expected code %d, got %d", tt.code, result.StatusCode)
			}
			defer result.Body.Close()
		})
	}
}

func TestHandlers_Registration(t *testing.T) {
	url := "/server/registration"
	tests := []struct {
		name string
		body string
		code int
	}{
		// TODO: Add test cases.
		{name: "Positive Reg test",
			body: `{"login": "login", "password": "password"}`,
			code: 200,
		},
		{
			name: "Negative Reg test",
			body: `{"login": 2341, "password": "password"}`,
			code: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody := bytes.NewBuffer([]byte(tt.body))
			r := httptest.NewRequest(http.MethodPost, url, requestBody)
			w := httptest.NewRecorder()

			h := NewRouter()
			h.ServeHTTP(w, r)
			result := w.Result()
			if result.StatusCode != tt.code {
				t.Errorf("Expected code %d, got %d", tt.code, result.StatusCode)
			}
			defer result.Body.Close()
		})
	}
}
