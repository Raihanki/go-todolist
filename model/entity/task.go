package entity

import "database/sql"

type Task struct {
	Id          int
	Name        string
	Description string
	IsCompleted bool
	CompletedAt sql.NullTime
}
