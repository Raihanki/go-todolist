package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/raihanki/todolist_go/config"
	"github.com/raihanki/todolist_go/controllers"
	"github.com/raihanki/todolist_go/exception"
	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/repositories"
	"github.com/raihanki/todolist_go/services"
)

func main() {
	db := config.ConnectDB()
	validate := validator.New()
	router := httprouter.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	taskRepository := repositories.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository, db, validate)
	taskController := controllers.NewTaskController(taskService)

	//user route
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)

	//task route
	router.GET("/api/tasks/all/:username", taskController.GetTaskByUsername)
	router.GET("/api/tasks/single/:taskId", taskController.GetTaskById)
	router.POST("/api/tasks", taskController.Store)
	router.PUT("/api/tasks/:taskId", taskController.Update)
	router.DELETE("/api/tasks/:taskId", taskController.Destroy)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
