package middleware

import "github.com/uptrace/bunrouter"

//go:generate gomockhandler -config=$EXAMPLE_TESTS_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
type MiddlewareInterface interface {
	AuthenticateMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	CorsMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	RecoverMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
}
