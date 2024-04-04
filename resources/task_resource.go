package resources

import (
	"encoding/json"
	"log"
	"time"

	"github.com/raihanki/todolist_go/model/entity"
)

type CustomJsonTime struct {
	time.Time
}

type TaskResource struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	IsCompleted bool           `json:"is_completed"`
	CompletedAt CustomJsonTime `json:"completed_at"`
}

func (t *CustomJsonTime) MarshalJSON() ([]byte, error) {
	log.Println(t.Time)
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}

func ToTaskResource(task entity.Task) TaskResource {
	return TaskResource{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CompletedAt: CustomJsonTime{task.CompletedAt.Time},
	}
}
