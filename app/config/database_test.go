package config

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func TestNewDBConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    *bun.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDBConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDBConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDBConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvWithDefault(t *testing.T) {
	type args struct {
		name string
		def  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnvWithDefault(tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("getEnvWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
