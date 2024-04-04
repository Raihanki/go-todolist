package repositories

import (
	"context"
	"database/sql"

	"github.com/raihanki/todolist_go/model/entity"
)

type UserRepository interface {
	Store(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error)
}
