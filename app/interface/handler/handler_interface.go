package _interface

import "github.com/uptrace/bunrouter"

type UserHandlerInterface interface {
	UserCreateHandle() bunrouter.HandlerFunc
	UserGetHandle() bunrouter.HandlerFunc
	ScoreUpdateHandle() bunrouter.HandlerFunc
	UserRankingGetHandle() bunrouter.HandlerFunc
	DestroyHandle() bunrouter.HandlerFunc
}
