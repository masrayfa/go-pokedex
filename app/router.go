package app

import (
	"ngacak-go/controller"
	"ngacak-go/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(pokemonController controller.PokemonController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	// Pokemon routes
	router.POST("/api/pokemon", pokemonController.Create)
	router.PUT("/api/pokemon/:pokemonId", pokemonController.Update)
	router.GET("/api/pokemon", pokemonController.FindAll)
	router.GET("/api/pokemon/id/:pokemonId", pokemonController.FindById)
	router.GET("/api/pokemon/name/:pokemonName", pokemonController.FindByName)
	router.DELETE("/api/pokemon/:pokemonId", pokemonController.Delete)

	// User routes
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.GET("/api/user", userController.FindAll)
	router.GET("/api/user/id/:userId", userController.FindById)
	router.GET("/api/user/name/:userName", userController.FindByUsername)
	router.DELETE("/api/user/:userId", userController.Delete)

	// Exception handler
	router.PanicHandler = exception.ErrorHandler

	return router
}