package _interface

import (
	"example.com/application/middleware"
	_interface "example.com/interface/handler"
	"github.com/uptrace/bunrouter"
)

type Router struct {
	UserHandler _interface.UserHandlerInterface
	Middleware  middleware.MiddlewareInterface
}

func NewRouter(userHandler _interface.UserHandlerInterface, middleware middleware.MiddlewareInterface) *Router {
	return &Router{
		UserHandler: userHandler,
		Middleware:  middleware,
	}
}

func (i *Router) InitRouter() *bunrouter.Router {
	b := bunrouter.New()
	b.Use(i.Middleware.RecoverMiddleware())
	b.Use(i.Middleware.CorsMiddleware())

	b.POST("/user/create", i.UserHandler.UserCreateHandle())

	b.Use(i.Middleware.AuthenticateMiddleware()).WithGroup("", func(group *bunrouter.Group) {
		group.POST("/user/get", i.UserHandler.UserGetHandle())
		group.POST("/user/score", i.UserHandler.ScoreUpdateHandle())
		group.POST("/user/destroy", i.UserHandler.DestroyHandle())
		group.GET("/users/get", i.UserHandler.UserRankingGetHandle())
	})

	return b
}
