package infrastructure

import (
	"context"
	"example.com/domain"
	"github.com/uptrace/bun"
	"log"
)

//go:generate gotests -w -all $GOFILE
type UserRepository struct {
	Conn *bun.DB
}

func NewUserRepository(Conn *bun.DB) *UserRepository {
	return &UserRepository{Conn: Conn}
}

func (u *UserRepository) AddUser(ctx context.Context, id, authToken, name string) error {
	user := &domain.User{
		Id:        id,
		AuthToken: authToken,
		Name:      name,
		HighScore: 0,
	}
	_, err := u.Conn.NewInsert().Model(user).Exec(ctx)
	return err
}

func (u *UserRepository) UpdateUser(ctx context.Context, user domain.User) error {
	uuser := &domain.User{
		Id:        user.Id,
		AuthToken: user.AuthToken,
		Name:      user.Name,
		HighScore: user.HighScore,
	}
	_, err := u.Conn.NewUpdate().Model(uuser).Where("id = ?", uuser.Id).Exec(ctx)
	log.Println(uuser.HighScore)
	return err
}

func (u *UserRepository) DeleteUser(ctx context.Context, id string) error {
	user := &domain.User{Id: id}
	_, err := u.Conn.NewDelete().Model(user).Where("id = ?", id).Exec(ctx)
	return err
}

func (u *UserRepository) GetUserByUserId(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.Conn.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByAuthToken(ctx context.Context, authToken string) (*domain.User, error) {
	user := new(domain.User)
	err := u.Conn.NewSelect().Model(user).Where("auth_token = ?", authToken).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserRanking(ctx context.Context) ([]*domain.UserRanking, error) {
	// Change var users []domain.User to var users []*domain.User
	var users []*domain.User

	log.Println("GetUserRanking")

	err := u.Conn.NewSelect().
		Column("name", "high_score").
		Model(&users).                // No change needed here as we're passing a pointer to a slice
		OrderExpr("high_score DESC"). // ハイスコアで降順にソート
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var userRankings []*domain.UserRanking
	for _, user := range users {
		ranking := &domain.UserRanking{
			Name:      user.Name,
			HighScore: user.HighScore,
		}
		userRankings = append(userRankings, ranking)
	}

	return userRankings, nil
}
