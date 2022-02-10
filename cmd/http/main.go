package main

import (
	"test/infrastructure/DB"
	"test/interanl/core/usecase"
	"test/interanl/delivery/http"
	"test/interanl/repository"
)

func main() {

	db := DB.New()

	userRepository := repository.NewPgUserRepository(db)

	usecase := usecase.New(userRepository)

	httpServer := http.New(usecase)
	httpServer.Run()
}
