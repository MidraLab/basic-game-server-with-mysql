package middleware

import "github.com/uptrace/bunrouter"

type MiddlewareInterface interface {
	AuthenticateMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	CorsMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
	RecoverMiddleware() func(bunrouter.HandlerFunc) bunrouter.HandlerFunc
}
