package middleware

import "github.com/uptrace/bunrouter"

//go:generate mockgen -source=auth_interface.go -destination=../../mocks/auth_mocks.go -package=mocks
type MiddlewareInterface interface {
	AuthenticateMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	CorsMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	RecoverMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
}
