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
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repositories.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Store(ctx context.Context, userCreateRequest request.UserCreateRequest) resources.UserResource {
	err := service.Validate.Struct(userCreateRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreateRequest.Password), 12)
	helper.PanicIfError(err)

	user := entity.User{
		Username: userCreateRequest.Username,
		Password: string(hashedPassword),
	}

	result := service.UserRepository.Store(ctx, tx, user)

	return resources.ToUserResource(result)
}

func (service *UserServiceImpl) Login(ctx context.Context, userLoginRequest request.UserLoginRequest) (bool, error) {
	err := service.Validate.Struct(userLoginRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, userLoginRequest.Username)
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return false, errors.New("wrong username password")
	} else {
		return true, nil
	}
}
