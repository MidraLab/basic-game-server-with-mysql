package infrastructure

import (
	"context"
	"reflect"
	"testing"

	"example.com/domain"
	"github.com/uptrace/bun"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		Conn *bun.DB
	}
	tests := []struct {
		name string
		args args
		want *UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.Conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_AddUser(t *testing.T) {
	type fields struct {
		Conn *bun.DB
	}
	type args struct {
		ctx       context.Context
		id        string
		authToken string
		name      string
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
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			if err := u.AddUser(tt.args.ctx, tt.args.id, tt.args.authToken, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateUser(t *testing.T) {
	type fields struct {
		Conn *bun.DB
	}
	type args struct {
		ctx  context.Context
		user domain.User
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
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			if err := u.UpdateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_DeleteUser(t *testing.T) {
	type fields struct {
		Conn *bun.DB
	}
	type args struct {
		ctx context.Context
		id  string
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
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			if err := u.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetUserByUserId(t *testing.T) {
	type fields struct {
		Conn *bun.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			got, err := u.GetUserByUserId(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUserByAuthToken(t *testing.T) {
	type fields struct {
		Conn *bun.DB
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
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			got, err := u.GetUserByAuthToken(tt.args.ctx, tt.args.authToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByAuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUserRanking(t *testing.T) {
	type fields struct {
		Conn *bun.DB
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
			u := &UserRepository{
				Conn: tt.fields.Conn,
			}
			got, err := u.GetUserRanking(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}
