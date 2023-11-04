package service

import (
	"context"
	"example.com/domain"
)

//go:generate gomockhandler -config=$EXAMPLE_TESTS_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type UserServiceInterface interface {
	Add(ctx context.Context, name string) (string, error)
	UpdateUser(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) (string, error)
	GetUserByUserId(ctx context.Context, id string) (domain.User, error)
	GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error)
	GetUserRanking(ctx context.Context) ([]*domain.UserRanking, error)
}
