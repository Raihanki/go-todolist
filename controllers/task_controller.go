package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TaskController interface {
	Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Destroy(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTaskByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTaskById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
