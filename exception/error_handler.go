package exception

import (
	"net/http"
	"ngacak-go/helper"
	"ngacak-go/model/web"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(writer, req, err) {
		return
	}

	if validationError(writer, req, err) {
		return
	}

	internalServerError(writer, req, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception , ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data: exception.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{ 
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, req *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse {
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
	
}