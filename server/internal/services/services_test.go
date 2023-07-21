package services

import (
	"io"
	"reflect"
	"strings"
	"testing"

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
	return 200, nil
}

func TestRegistration(t *testing.T) {
	var keeper keeperMock
	type args struct {
		data   io.Reader
		keeper storage.Keeper
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "positive test",
			args: args{
				data:   strings.NewReader(`{"login": "login", "password": "password"}`),
				keeper: keeper,
			},
			want:    200,
			wantErr: false,
		},
		{
			name: "negative test",
			args: args{
				data:   strings.NewReader(`{"login": 1231, "password": "password"}`),
				keeper: keeper,
			},
			want:    400,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Registration(tt.args.data, tt.args.keeper)
			if (err != nil) != tt.wantErr {
				t.Errorf("Registration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Registration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	var keeper keeperMock
	type args struct {
		data   io.Reader
		secret []byte
		keeper storage.Keeper
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "positive test",
			args: args{
				data:   strings.NewReader(`{"login": "login", "password": "password"}`),
				keeper: keeper,
				secret: []byte(""),
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoibG9naW4ifQ.6geBsSGpHdriKCPGxM4EF9MjpC2gVVl4sXe-SjguT5Y",
			want1:   200,
			wantErr: false,
		},
		{
			name: "negative test",
			args: args{
				data:   strings.NewReader(`{"login": 1231, "password": "password"}`),
				keeper: keeper,
				secret: []byte(""),
			},
			want:    "",
			want1:   400,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Login(tt.args.data, tt.args.secret, tt.args.keeper)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewEntry(t *testing.T) {
	var keeper keeperMock
	type args struct {
		data   io.Reader
		keeper storage.Keeper
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "positive test",
			args: args{
				data:   strings.NewReader(`{"user": "user", "service": "service", "entry": "entry", "meta": "meta"}`),
				keeper: keeper,
			},
			want:    200,
			wantErr: false,
		},
		{
			name: "negative test",
			args: args{
				data:   strings.NewReader(`{"user": user, "service": "service", "entry": "entry", "meta": "meta"}`),
				keeper: keeper,
			},
			want:    400,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEntry(tt.args.data, tt.args.keeper)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateEntry(t *testing.T) {
	var keeper keeperMock
	type args struct {
		data   io.Reader
		keeper storage.Keeper
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		{
			name: "positive test",
			args: args{
				data:   strings.NewReader(`{"user": "user", "service": "service", "entry": "entry", "meta": "meta"}`),
				keeper: keeper,
			},
			want:  []byte("entry updatet"),
			want1: 200,
		},
		{
			name: "negative test",
			args: args{
				data:   strings.NewReader(`{"user": user, "service": "service", "entry": "entry", "meta": "meta"}`),
				keeper: keeper,
			},
			want1: 500,
			want:  []byte("invalid character 'u' looking for beginning of value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := UpdateEntry(tt.args.data, tt.args.keeper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateEntry() got = %s, want %s", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UpdateEntry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetEntry(t *testing.T) {
	var keeper keeperMock
	type args struct {
		username string
		service  string
		keeper   storage.Keeper
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "Positive test",
			args: args{
				username: "Username",
				service:  "Service",
				keeper:   keeper,
			},
			want:  []byte(`{"user":"","service":"","entry":"","metadata":""}`),
			want1: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetEntry(tt.args.username, tt.args.service, tt.args.keeper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEntry() got = %s, want %s", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetEntry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeleteEntry(t *testing.T) {
	var keeper keeperMock
	type args struct {
		username string
		service  string
		keeper   storage.Keeper
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive test",
			args: args{
				username: "Username",
				service:  "Service",
				keeper:   keeper,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteEntry(tt.args.username, tt.args.service, tt.args.keeper); got != tt.want {
				t.Errorf("DeleteEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
