package _interface

import "github.com/uptrace/bunrouter"

//go:generate gomockhandler -config=$EXAMPLE_TESTS_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type UserHandlerInterface interface {
	UserCreateHandle() bunrouter.HandlerFunc
	UserGetHandle() bunrouter.HandlerFunc
	ScoreUpdateHandle() bunrouter.HandlerFunc
	UserRankingGetHandle() bunrouter.HandlerFunc
	DestroyHandle() bunrouter.HandlerFunc
}
