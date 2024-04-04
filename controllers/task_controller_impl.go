package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/request"
	"github.com/raihanki/todolist_go/services"
)

type TaskControllerImpl struct {
	TaskService services.TaskService
}

func NewTaskController(taskService services.TaskService) TaskController {
	return &TaskControllerImpl{
		TaskService: taskService,
	}
}

func (controller *TaskControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	taskCreateRequest := request.TaskCreateRequest{}
	err := decoder.Decode(&taskCreateRequest)
	helper.PanicIfError(err)

	task := controller.TaskService.Store(r.Context(), taskCreateRequest)
	response := helper.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Created",
		Data:    task,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}

func (controller *TaskControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	taskUpdateRequest := request.TaskUpdateRequest{}
	err := decoder.Decode(&taskUpdateRequest)
	helper.PanicIfError(err)

	task, err := controller.TaskService.Update(r.Context(), taskUpdateRequest)
	if err != nil {
		response := helper.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	} else {
		response := helper.ApiResponse{
			Code:    http.StatusOK,
			Message: "Updated",
			Data:    task,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	}
}

func (controller *TaskControllerImpl) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	intTaskId, err := strconv.Atoi(taskId)
	helper.PanicIfError(err)

	controller.TaskService.Destroy(r.Context(), intTaskId)
	response := helper.ApiResponse{
		Code:    http.StatusOK,
		Message: "Deleted",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	helper.PanicIfError(err)
}

func (controller *TaskControllerImpl) GetTaskByUsername(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	tasks := controller.TaskService.GetTaskByUsername(r.Context(), username)
	response := helper.ApiResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	helper.PanicIfError(err)
}

func (controller *TaskControllerImpl) GetTaskById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	taskId := params.ByName("taskId")
	intTaskId, err := strconv.Atoi(taskId)
	helper.PanicIfError(err)

	task, err := controller.TaskService.GetTaskById(r.Context(), intTaskId)
	if err != nil {
		response := helper.ApiResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	} else {
		response := helper.ApiResponse{
			Code:    http.StatusOK,
			Message: "OK",
			Data:    task,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(response)
		helper.PanicIfError(err)
	}
}
