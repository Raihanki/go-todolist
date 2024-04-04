package exception

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/raihanki/todolist_go/helper"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		r.Header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := helper.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation Error",
			Data:    exception.Error(),
		}
		encoder := json.NewEncoder(w)
		e := encoder.Encode(response)
		helper.PanicIfError(e)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	response := helper.ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    err,
	}
	encoder := json.NewEncoder(w)
	e := encoder.Encode(response)
	helper.PanicIfError(e)
}
