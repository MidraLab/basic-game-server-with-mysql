package middleware

import (
	"reflect"
	"testing"

	"example.com/application/service"
	"github.com/uptrace/bunrouter"
)

func TestNewMiddleware(t *testing.T) {
	type args struct {
		userService *service.UserService
	}
	tests := []struct {
		name string
		args args
		want *Middleware
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMiddleware(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleware_AuthenticateMiddleware(t *testing.T) {
	type fields struct {
		UserService *service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := m.AuthenticateMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.AuthenticateMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleware_CorsMiddleware(t *testing.T) {
	type fields struct {
		UserService *service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := m.CorsMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.CorsMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleware_RecoverMiddleware(t *testing.T) {
	type fields struct {
		UserService *service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := m.RecoverMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.RecoverMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
