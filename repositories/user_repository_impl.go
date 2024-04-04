package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/model/entity"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Store(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
