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
	pokemonRepository := repository.NewPokemonRepository()
	pokemonService := service.NewPokemonService(pokemonRepository, db, validate)
	pokemonController := controller.NewPokemonController(pokemonService)

	router := app.NewRouter(pokemonController)

	server := http.Server{
		Addr:   "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}