package resources

import "github.com/raihanki/todolist_go/model/entity"

type UserResource struct {
	Username string `json:"username"`
}

func ToUserResource(user entity.User) UserResource {
	return UserResource{
		Username: user.Username,
	}
}
