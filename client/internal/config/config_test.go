package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "Positive congif",
			want: &Config{
				Address:     "127.0.0.1:8080",
				RequestType: "",
				Login:       "",
				Password:    "",
				Entry:       "",
				Meta:        "",
				Service:     "",
				JWT:         "",
				Key:         "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
