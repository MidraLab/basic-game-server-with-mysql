package _interface

import "github.com/uptrace/bunrouter"

//go:generate mockgen -source=handler_interface.go -destination=../../mocks/handler_mock.go -package=mocks
type UserHandlerInterface interface {
	UserCreateHandle() bunrouter.HandlerFunc
	UserGetHandle() bunrouter.HandlerFunc
	ScoreUpdateHandle() bunrouter.HandlerFunc
	UserRankingGetHandle() bunrouter.HandlerFunc
	DestroyHandle() bunrouter.HandlerFunc
}
