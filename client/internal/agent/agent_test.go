package agent

import (
	"reflect"
	"testing"

	"gtihub.com/popooq/Gophkeeper/client/internal/config"
	"gtihub.com/popooq/Gophkeeper/client/internal/saver"
	"gtihub.com/popooq/Gophkeeper/client/internal/sender"
)

var endpoint string = "127.0.0.1:8080"

func TestNew(t *testing.T) {
	type args struct {
		config *config.Config
		sender sender.Sender
	}
	tests := []struct {
		name string
		args args
		want *Agent
	}{
		{
			name: "Positive test",
			args: args{
				config: &config.Config{
					RequestType: "",
					Login:       "login",
					Password:    "password",
					Service:     "service",
					Entry:       "entry",
					Meta:        "meta",
				},
				sender: sender.Sender{},
			},
			want: &Agent{
				config: &config.Config{
					RequestType: "",
					Login:       "login",
					Password:    "password",
					Service:     "service",
					Entry:       "entry",
					Meta:        "meta",
				},
				sender: sender.Sender{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config, tt.args.sender); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgent_Run(t *testing.T) {
	type fields struct {
		config *config.Config
		sender sender.Sender
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "positive reg",
			fields: fields{
				config: &config.Config{
					RequestType: "reg",
					Login:       "login",
					Password:    "password",
				},
			},
		},
		{
			name: "positive login",
			fields: fields{
				config: &config.Config{
					RequestType: "login",
					Login:       "login",
					Password:    "password",
				},
			},
		},
		{
			name: "positive add",
			fields: fields{
				config: &config.Config{
					RequestType: "add",
					Login:       "login",
					Service:     "service",
					Entry:       "entry",
					Meta:        "",
				},
				sender: sender.New(endpoint, saver.New("jwt")),
			},
		},
		{
			name: "positive update",
			fields: fields{
				config: &config.Config{
					RequestType: "update",
					Login:       "login",
					Service:     "service",
					Entry:       "entry",
					Meta:        "",
				},
				sender: sender.New(endpoint, saver.New("jwt")),
			},
		},
		{
			name: "positive get",
			fields: fields{
				config: &config.Config{
					RequestType: "get",
					Login:       "login",
					Service:     "service",
				},
				sender: sender.New(endpoint, saver.New("jwt")),
			},
		},
		{
			name: "positive delete",
			fields: fields{
				config: &config.Config{
					RequestType: "delete",
					Login:       "login",
					Service:     "service",
				},
				sender: sender.New(endpoint, saver.New("jwt")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Agent{
				config: tt.fields.config,
				sender: tt.fields.sender,
			}
			a.Run()
		})
	}
}
