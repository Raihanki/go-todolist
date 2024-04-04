package repositories

import (
	"context"
	"database/sql"

	"github.com/raihanki/todolist_go/model/entity"
)

type TaskRepository interface {
	Store(ctx context.Context, tx *sql.Tx, task entity.Task) entity.Task
	FindTaskByUsername(ctx context.Context, tx *sql.Tx, username string) []entity.Task
	FindTaskById(ctx context.Context, tx *sql.Tx, taskId int) (entity.Task, error)
	Update(ctx context.Context, tx *sql.Tx, task entity.Task) entity.Task
	Destroy(ctx context.Context, tx *sql.Tx, taskId int)
}
