package service

import (
	"context"
	"reflect"
	"testing"

	"example.com/domain"
	"example.com/domain/repository"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		UserRepository repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.UserRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Add(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.Add(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			if err := u.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByUserId(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.GetUserByUserId(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByAuthToken(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx       context.Context
		authToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.GetUserByAuthToken(tt.args.ctx, tt.args.authToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByAuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserRanking(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.UserRanking
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.GetUserRanking(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}
