package request

import (
	"database/sql"
)

type TaskUpdateRequest struct {
	Id          int          `json:"id"`
	Name        string       `json:"name" validate:"required,min=3,max=199"`
	Description string       `json:"description" validate:"required,min=3"`
	IsCompleted bool         `json:"is_completed"`
	CompletedAt sql.NullTime `json:"completed_at"`
}
