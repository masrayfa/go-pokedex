package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"ngacak-go/controller"
	"ngacak-go/exception"

	"github.com/julienschmidt/httprouter"
)

type Pokemon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func NewRouter(pokemonController controller.PokemonController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	// Pokemon routes
	router.GET("/api/pokemon", pokemonController.FindAll)
	router.GET("/api/pokemon/collections/:userId", pokemonController.FindCollections)
	router.GET("/api/pokemon/id/:pokemonId", pokemonController.FindById)
	router.GET("/api/pokemon/name/:pokemonName", pokemonController.FindByName)
	router.POST("/api/pokemon", pokemonController.Create)
	router.PUT("/api/pokemon/:pokemonId", pokemonController.Update)
	router.DELETE("/api/pokemon/:pokemonId", pokemonController.Delete)

	// User routes
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.GET("/api/user", userController.FindAll)
	router.GET("/api/user/id/:userId", userController.FindById)
	router.GET("/api/user/name/:userName", userController.FindByUsername)
	router.DELETE("/api/user/:userId", userController.Delete)

	// Views routes
	router.GET("/", (func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Hit API first
		var pokedex []Pokemon
		for i := 1; i <= 12; i++ {
			url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", i)

			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Error membuat permintaan:", err)
				return
			}
			defer res.Body.Close()

			// Memeriksa apakah permintaan berhasil (kode status 200)
			if res.StatusCode == http.StatusOK {
				fmt.Printf("Permintaan ke %s berhasil!\n", url)
				// Handle data response sesuai kebutuhan Anda
				var pokemon Pokemon
				err = json.NewDecoder(res.Body).Decode(&pokemon)
				if err != nil {
					fmt.Println("Error parsing JSON:", err)
					return
				}

				pokedex = append(pokedex, pokemon)
				fmt.Println(pokedex)
			} else {
				fmt.Printf("Permintaan ke %s gagal, status code: %d\n", url, res.StatusCode)
			}
		}

		// }
		// webResponse := web.WebResponse {
		// 	Code: 200,
		// 	Status: "OK",
		// 	Data: pokedex,
		// }
		// formattedResult, err := json.MarshalIndent(webResponse, "", "    ")
		// if err != nil {
		// 	fmt.Println("Error formatting JSON:", err)
		// 	return
		// }

		// url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/ditto")
		// respPokedex, err := http.Get(url)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// defer respPokedex.Body.Close()

		// if respPokedex.StatusCode != http.StatusOK {
		// 	http.Error(w, respPokedex.Status, respPokedex.StatusCode)
		// 	return 
		// }

		// var pokedex interface{}
		// err = json.NewDecoder(respPokedex.Body).Decode(&pokedex)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Println("ini pokem resp",pokedex)

		// LAMA
		// resp, err := http.Get("http://localhost:8080/api/pokemon")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// defer resp.Body.Close()

		// if resp.StatusCode != http.StatusOK {
		// 	http.Error(w, resp.Status, resp.StatusCode)
		// 	return 
		// }

		// var pokemonResponse interface{}
		// err = json.NewDecoder(resp.Body).Decode(&pokemonResponse)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// formattedResult, err := json.MarshalIndent(pokemonResponse, "", "    ")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Println(pokemonResponse)

		// fmt.Println(string(formattedResult))

		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, pokedex)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// http.ServeFile(w, r, "views/index.html")
	}))

	// Exception handler
	router.PanicHandler = exception.ErrorHandler

	return router
}