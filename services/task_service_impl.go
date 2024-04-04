package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator"
	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/model/entity"
	"github.com/raihanki/todolist_go/repositories"
	"github.com/raihanki/todolist_go/request"
	"github.com/raihanki/todolist_go/resources"
)

type TaskServiceImpl struct {
	TaskRepository repositories.TaskRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTaskService(taskRepository repositories.TaskRepository, db *sql.DB, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *TaskServiceImpl) Store(ctx context.Context, taskCreateRequest request.TaskCreateRequest) resources.TaskResource {
	err := service.Validate.Struct(taskCreateRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task := entity.Task{
		Name:        taskCreateRequest.Name,
		Description: taskCreateRequest.Description,
		IsCompleted: false,
	}

	taskCreated := service.TaskRepository.Store(ctx, tx, task)

	return resources.ToTaskResource(taskCreated)
}

func (service *TaskServiceImpl) Update(ctx context.Context, taskUpdateRequest request.TaskUpdateRequest) (resources.TaskResource, error) {
	err := service.Validate.Struct(taskUpdateRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindTaskById(ctx, tx, taskUpdateRequest.Id)
	if err != nil {
		return resources.TaskResource{}, errors.New("task not found")
	}

	task.Name = taskUpdateRequest.Name
	task.Description = taskUpdateRequest.Description
	task.IsCompleted = taskUpdateRequest.IsCompleted
	task.CompletedAt = taskUpdateRequest.CompletedAt

	updatedTask := service.TaskRepository.Update(ctx, tx, task)
	return resources.ToTaskResource(updatedTask), nil
}

func (service *TaskServiceImpl) Destroy(ctx context.Context, taskId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindTaskById(ctx, tx, taskId)
	helper.PanicIfError(err)

	service.TaskRepository.Destroy(ctx, tx, task.Id)
}

func (service *TaskServiceImpl) GetTaskById(ctx context.Context, taskId int) (resources.TaskResource, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindTaskById(ctx, tx, taskId)
	if err != nil {
		return resources.ToTaskResource(task), errors.New("task not found")
	}

	return resources.ToTaskResource(task), nil
}

func (service *TaskServiceImpl) GetTaskByUsername(ctx context.Context, username string) []resources.TaskResource {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	results := service.TaskRepository.FindTaskByUsername(ctx, tx, username)
	tasks := []resources.TaskResource{}
	for _, result := range results {
		tasks = append(tasks, resources.ToTaskResource(result))
	}

	return tasks
}
