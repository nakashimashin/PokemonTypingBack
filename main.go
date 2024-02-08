package main

import (
	"back/controller"
	"back/db"
	"back/repository"
	"back/router"
	"back/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	r := router.NewRouter(userController)
	r.Run()
}
