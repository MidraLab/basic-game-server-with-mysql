package _interface

import (
	"reflect"
	"testing"

	"example.com/application/service"
	"github.com/uptrace/bunrouter"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		userService *service.UserService
	}
	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_UserCreateHandle(t *testing.T) {
	type fields struct {
		userService service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				userService: tt.fields.userService,
			}
			if got := u.UserCreateHandle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHandler.UserCreateHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_UserGetHandle(t *testing.T) {
	type fields struct {
		userService service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				userService: tt.fields.userService,
			}
			if got := u.UserGetHandle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHandler.UserGetHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_ScoreUpdateHandle(t *testing.T) {
	type fields struct {
		userService service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				userService: tt.fields.userService,
			}
			if got := u.ScoreUpdateHandle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHandler.ScoreUpdateHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_UserRankingGetHandle(t *testing.T) {
	type fields struct {
		userService service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				userService: tt.fields.userService,
			}
			if got := u.UserRankingGetHandle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHandler.UserRankingGetHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_DestroyHandle(t *testing.T) {
	type fields struct {
		userService service.UserService
	}
	tests := []struct {
		name   string
		fields fields
		want   bunrouter.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				userService: tt.fields.userService,
			}
			if got := u.DestroyHandle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserHandler.DestroyHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}
