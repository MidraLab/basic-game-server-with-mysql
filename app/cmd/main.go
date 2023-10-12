// @title           Game Server API
// @version         1.0
// @description     This project is a one-command, feature-complete game server starter.
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name x-token
package main

import (
	"example.com/application/middleware"
	"example.com/application/service"
	"example.com/config"
	infrastructure "example.com/infrastructure/persistence"
	_interface "example.com/interface/handler"
	"github.com/uptrace/bunrouter"
	"log"
	"net/http"
)

func main() {
	db, _ := config.NewDBConnection()
	//db_init.CreateTable(db)

	userRepository := infrastructure.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := _interface.NewUserHandler(userService)
	middleware := middleware.NewMiddleware(userService)

	r := bunrouter.New()
	r.Use(middleware.RecoverMiddleware())
	r.Use(middleware.CorsMiddleware())

	r.POST("/user/create", userHandler.UserCreateHandle())

	r.Use(middleware.AuthenticateMiddleware()).WithGroup("", func(group *bunrouter.Group) {
		group.POST("/user/get", userHandler.UserGetHandle())
		group.POST("/user/score", userHandler.ScoreUpdateHandle())
		group.POST("/user/destroy", userHandler.DestroyHandle())
		group.GET("/users/get", userHandler.UserRankingGetHandle())
	})

	log.Println("listening on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", r))
}
