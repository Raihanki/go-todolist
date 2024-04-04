package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/raihanki/todolist_go/helper"
	"github.com/raihanki/todolist_go/model/entity"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) Store(ctx context.Context, tx *sql.Tx, task entity.Task) entity.Task {
	query := "INSERT INTO tasks (name, description) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, task.Name, task.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	task.Id = int(id)
	task.IsCompleted = false
	return task
}

func (repository *TaskRepositoryImpl) FindTaskByUsername(ctx context.Context, tx *sql.Tx, username string) []entity.Task {
	query := "SELECT id, name, description, is_completed, completed_at FROM tasks WHERE username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	helper.PanicIfError(err)
	defer rows.Close()

	tasks := []entity.Task{}
	for rows.Next() {
		task := entity.Task{}
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.IsCompleted, &task.CompletedAt)
		helper.PanicIfError(err)

		tasks = append(tasks, task)
	}

	return tasks
}

func (repository *TaskRepositoryImpl) FindTaskById(ctx context.Context, tx *sql.Tx, taskId int) (entity.Task, error) {
	query := "SELECT id, name, description, is_completed, completed_at FROM tasks WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, taskId)
	helper.PanicIfError(err)
	defer rows.Close()

	task := entity.Task{}
	var completedAt sql.NullTime
	if rows.Next() {
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.IsCompleted, &completedAt)
		task.CompletedAt = completedAt
		if err != nil {
			log.Println("Error: ", err)
			return task, err
		}

		return task, nil
	} else {
		return task, errors.New("task not found")
	}
}

func (repository *TaskRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, task entity.Task) entity.Task {
	query := "UPDATE tasks SET name = ?, description = ?, is_completed = ?, completed_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, task.Name, task.Description, task.IsCompleted, task.CompletedAt, task.Id)
	helper.PanicIfError(err)

	return task
}

func (repository *TaskRepositoryImpl) Destroy(ctx context.Context, tx *sql.Tx, taskId int) {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, taskId)
	helper.PanicIfError(err)
}
