package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/request"
	"github.com/raihanki/todolist_go/services"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	userCreateRequest := request.UserCreateRequest{}
	err := decoder.Decode(&userCreateRequest)
	helper.PanicIfError(err)

	userResource := controller.UserService.Store(r.Context(), userCreateRequest)
	response := helper.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    userResource,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	userLoginRequest := request.UserLoginRequest{}
	err := decoder.Decode(&userLoginRequest)
	helper.PanicIfError(err)

	_, err = controller.UserService.Login(r.Context(), userLoginRequest)
	if err != nil {
		response := helper.ApiResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	} else {
		response := helper.ApiResponse{
			Code:    http.StatusOK,
			Message: "OK",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	}
}
