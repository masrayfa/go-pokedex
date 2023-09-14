package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUsernameAndPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
