package controller

import (
	"net/http"
	"ngacak-go/service"

	"github.com/julienschmidt/httprouter"
)

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

type userControllerImpl struct {
	UserService service.UserService 
}

func (controller *userControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) FindByUsernameAndPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
