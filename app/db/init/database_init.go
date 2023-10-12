package db_init

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type User struct {
	Id        string `bun:"type:varchar(128),primary" json:"id"`        // ユーザID
	AuthToken string `bun:"type:varchar(128),unique" json:"auth_token"` // 認証トークン
	Name      string `bun:"type:varchar(64)" json:"name"`               // ユーザ名
	HighScore int    `bun:"type:int unsigned" json:"high_score"`        // ハイスコア
}

func CreateTable(db *bun.DB) {
	_, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("createTable")
}
