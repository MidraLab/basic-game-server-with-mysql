package auth

import (
	"context"
	"reflect"
	"testing"
)

func TestSetUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUserID(tt.args.ctx, tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserIDFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
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
			if got := GetUserIDFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("GetUserIDFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
