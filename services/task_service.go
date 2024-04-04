package services

import (
	"context"

	"github.com/raihanki/todolist_go/request"
	"github.com/raihanki/todolist_go/resources"
)

type TaskService interface {
	Store(ctx context.Context, taskCreateRequest request.TaskCreateRequest) resources.TaskResource
	Update(ctx context.Context, taskUpdateRequest request.TaskUpdateRequest) (resources.TaskResource, error)
	Destroy(ctx context.Context, taskId int)
	GetTaskByUsername(ctx context.Context, username string) []resources.TaskResource
	GetTaskById(ctx context.Context, taskId int) (resources.TaskResource, error)
}
