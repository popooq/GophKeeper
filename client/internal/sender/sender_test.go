package sender

import (
	"reflect"
	"testing"

	"gtihub.com/popooq/Gophkeeper/client/internal/saver"
)

var endpoint string = "127.0.0.1:8080"

func TestNew(t *testing.T) {
	type args struct {
		endpoint string
		saver    *saver.Saver
	}
	tests := []struct {
		name string
		args args
		want Sender
	}{
		{
			name: "testNew",
			args: args{
				saver:    nil,
				endpoint: endpoint,
			},
			want: Sender{
				endpoint: endpoint,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.endpoint, tt.args.saver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_Reg(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    nil,
			},
			args: args{
				login:    "login",
				password: "password",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Reg(tt.args.login, tt.args.password)
		})
	}
}

func TestSender_Login(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    nil,
			},
			args: args{
				login:    "login",
				password: "password",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Login(tt.args.login, tt.args.password)
		})
	}
}

func TestSender_Add(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		user     string
		service  string
		entry    string
		metadata string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    saver.New("jwt"),
			},
			args: args{
				user:     "test_user",
				service:  "test_service",
				entry:    "test_entry",
				metadata: "test_metadata",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Add(tt.args.user, tt.args.service, tt.args.entry, tt.args.metadata)
		})
	}
}

func TestSender_Update(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		user     string
		service  string
		entry    string
		metadata string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    saver.New("jwt"),
			},
			args: args{
				user:     "test_user",
				service:  "test_service",
				entry:    "test_entry",
				metadata: "test_metadata",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Update(tt.args.user, tt.args.service, tt.args.entry, tt.args.metadata)
		})
	}
}

func TestSender_Get(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		user    string
		service string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    saver.New("jwt"),
			},
			args: args{
				user:    "test_user",
				service: "test_service",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Get(tt.args.user, tt.args.service)
		})
	}
}

func TestSender_Delete(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		user    string
		service string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    saver.New("jwt"),
			},
			args: args{
				user:    "test_user",
				service: "test_service",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			s.Delete(tt.args.user, tt.args.service)
		})
	}
}

func TestSender_entryBuild(t *testing.T) {
	type fields struct {
		endpoint string
		saver    *saver.Saver
	}
	type args struct {
		user     string
		service  string
		entry    string
		metadata string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "positive test",
			fields: fields{
				endpoint: endpoint,
				saver:    saver.New("jwt"),
			},
			args: args{
				user:     "test_user",
				service:  "test_service",
				entry:    "test_entry",
				metadata: "test_metadata",
			},
			want: []byte(`{"user":"test_user","service":"test_service","entry":"test_entry","metadata":"test_metadata"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				endpoint: tt.fields.endpoint,
				saver:    tt.fields.saver,
			}
			if got := s.entryBuild(tt.args.user, tt.args.service, tt.args.entry, tt.args.metadata); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sender.entryBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}
