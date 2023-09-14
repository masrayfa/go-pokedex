package main

import (
	"net/http"
	"ngacak-go/app"
	"ngacak-go/controller"
	"ngacak-go/helper"
	"ngacak-go/repository"
	"ngacak-go/service"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	// Reposistory
	pokemonRepository := repository.NewPokemonRepository()
	userRepository := repository.NewUserRepository()

	// Service
	pokemonService := service.NewPokemonService(pokemonRepository, db, validate)
	userService := service.NewUserService(userRepository, db, validate)

	// Controller
	pokemonController := controller.NewPokemonController(pokemonService)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(pokemonController, userController)

	server := http.Server{
		Addr:   "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}