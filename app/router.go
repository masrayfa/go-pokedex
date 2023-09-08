package app

import (
	"ngacak-go/controller"
	"ngacak-go/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(pokemonController controller.PokemonController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/pokemon", pokemonController.Create)
	router.PUT("/api/pokemon/:pokemonId", pokemonController.Update)
	router.GET("/api/pokemon", pokemonController.FindAll)
	router.GET("/api/pokemon/id/:pokemonId", pokemonController.FindById)
	router.GET("/api/pokemon/name/:pokemonName", pokemonController.FindByName)
	router.DELETE("/api/pokemon/:pokemonId", pokemonController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}