package services

import (
	"context"

	"github.com/raihanki/todolist_go/request"
	"github.com/raihanki/todolist_go/resources"
)

type UserService interface {
	Store(ctx context.Context, userCreateRequest request.UserCreateRequest) resources.UserResource
	Login(ctx context.Context, userLoginRequest request.UserLoginRequest) (bool, error)
}
