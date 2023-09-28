package controller

import (
	"net/http"
	"ngacak-go/helper"
	"ngacak-go/model/web"
	"ngacak-go/service"
	"strconv"

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
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse, err := controller.UserService.FindById(request.Context(), id)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("userName")

	response, err := controller.UserService.FindByUsername(request.Context(), username)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) FindByUsernameAndPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse, err := controller.UserService.Create(request.Context(), userCreateRequest)
	helper.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *userControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func (controller *userControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}
