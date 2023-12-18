package controller

import (
	"encoding/json"
	"net/http"
	"ngacak-go/helper"
	"ngacak-go/model/web"
	"ngacak-go/service"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var templates *template.Template

type PokemonControllerImpl struct {
	PokemonService service.PokemonService 
}

func NewPokemonController(pokemonService service.PokemonService) PokemonController {
	return &PokemonControllerImpl{PokemonService: pokemonService}
}


func (controller *PokemonControllerImpl) Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	pokemonCreateRequest := web.PokemonCreateRequest{}
	helper.ReadFromRequestBody(req, &pokemonCreateRequest)

	pokemonResponse, err := controller.PokemonService.Create(req.Context(), pokemonCreateRequest)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(controller *PokemonControllerImpl) Update (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	pokemonUpdateRequest := web.PokemonUpdateRequest{}
	helper.ReadFromRequestBody(req, &web.PokemonUpdateRequest{})

	paramsId := params.ByName("pokemonId")
	id, err := strconv.Atoi(paramsId)
	helper.PanicIfError(err)

	pokemonUpdateRequest.Id = id

	pokemonResponse, err := controller.PokemonService.Update(req.Context(), pokemonUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PokemonControllerImpl) Delete (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	pokemonId := params.ByName("pokemonId")
	id, err := strconv.Atoi(pokemonId)
	helper.PanicIfError(err)

	controller.PokemonService.Delete(req.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PokemonControllerImpl) FindAll (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// template := template.Must(template.ParseFiles("views/index.html"))
	
	pokemonResponses, err := controller.PokemonService.FindAll(req.Context())
	helper.PanicIfError(err)
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	// template.Execute(writer, webResponse)
}

func (controller *PokemonControllerImpl) FindById (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	pokemonId := params.ByName("pokemonId")
	id, err := strconv.Atoi(pokemonId)
	helper.PanicIfError(err)

	pokemonResponse, err := controller.PokemonService.FindById(req.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PokemonControllerImpl) FindByName (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	pokemonName := params.ByName("pokemonName")

	pokemonResponse , err := controller.PokemonService.FindByName(req.Context(), pokemonName)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder((writer))
	errorResponse := encoder.Encode(webResponse)
	helper.PanicIfError(errorResponse)
}

func (controller *PokemonControllerImpl) FindCollections (writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	pokemonResponses, err := controller.PokemonService.FindCollections(req.Context(), id)
	helper.PanicIfError(err)
	
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pokemonResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}