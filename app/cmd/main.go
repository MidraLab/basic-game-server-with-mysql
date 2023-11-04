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
	handler "example.com/interface/handler"
	router "example.com/interface/router"
	"log"
	"net/http"
)

func main() {
	db, _ := config.NewDBConnection()
	//db_init.CreateTable(db)

	userRepository := infrastructure.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	middleware := middleware.NewMiddleware(userService)
	router := router.NewRouter(userHandler, middleware)
	r := router.InitRouter()

	log.Println("listening on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", r))
}
