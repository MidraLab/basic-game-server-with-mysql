package repository

import (
	"context"
	"example.com/domain"
)

//go:generate gomockhandler -config=$EXAMPLE_TESTS_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type UserRepository interface {
	AddUser(ctx context.Context, id, authToken, name string) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByUserId(ctx context.Context, id string) (domain.User, error)
	GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error)
	GetUserRanking(ctx context.Context) ([]*domain.UserRanking, error)
}
